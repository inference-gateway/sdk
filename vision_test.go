package sdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func TestVisionMessage_TextContent(t *testing.T) {
	msg := Message{
		Role: User,
	}

	err := msg.Content.FromMessageContent0("What is in this image?")
	require.NoError(t, err)

	text, err := msg.Content.AsMessageContent0()
	require.NoError(t, err)
	assert.Equal(t, "What is in this image?", text)
}

func TestVisionMessage_ImageURLContent(t *testing.T) {
	msg := Message{
		Role: User,
	}

	contentParts := []ContentPart{}

	var textPart ContentPart
	err := textPart.FromTextContentPart(TextContentPart{
		Type: Text,
		Text: "What is in this image?",
	})
	require.NoError(t, err)
	contentParts = append(contentParts, textPart)

	var imagePart ContentPart
	err = imagePart.FromImageContentPart(ImageContentPart{
		Type: ImageContentPartTypeImageURL,
		ImageURL: ImageURL{
			URL:    "https://example.com/image.jpg",
			Detail: new(Auto),
		},
	})
	require.NoError(t, err)
	contentParts = append(contentParts, imagePart)

	err = msg.Content.FromMessageContent1(contentParts)
	require.NoError(t, err)

	parts, err := msg.Content.AsMessageContent1()
	require.NoError(t, err)
	assert.Len(t, parts, 2)

	textPartRetrieved, err := parts[0].AsTextContentPart()
	require.NoError(t, err)
	assert.Equal(t, Text, textPartRetrieved.Type)
	assert.Equal(t, "What is in this image?", textPartRetrieved.Text)

	imagePartRetrieved, err := parts[1].AsImageContentPart()
	require.NoError(t, err)
	assert.Equal(t, ImageContentPartTypeImageURL, imagePartRetrieved.Type)
	assert.Equal(t, "https://example.com/image.jpg", imagePartRetrieved.ImageURL.URL)
	assert.NotNil(t, imagePartRetrieved.ImageURL.Detail)
	assert.Equal(t, Auto, *imagePartRetrieved.ImageURL.Detail)
}

func TestVisionMessage_DataURLImage(t *testing.T) {
	msg := Message{
		Role: User,
	}

	dataURL := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD..."

	contentParts := []ContentPart{}

	var imagePart ContentPart
	err := imagePart.FromImageContentPart(ImageContentPart{
		Type: ImageContentPartTypeImageURL,
		ImageURL: ImageURL{
			URL:    dataURL,
			Detail: new(High),
		},
	})
	require.NoError(t, err)
	contentParts = append(contentParts, imagePart)

	err = msg.Content.FromMessageContent1(contentParts)
	require.NoError(t, err)

	parts, err := msg.Content.AsMessageContent1()
	require.NoError(t, err)
	assert.Len(t, parts, 1)

	imagePartRetrieved, err := parts[0].AsImageContentPart()
	require.NoError(t, err)
	assert.Equal(t, dataURL, imagePartRetrieved.ImageURL.URL)
	assert.Equal(t, High, *imagePartRetrieved.ImageURL.Detail)
}

func TestGenerateContent_WithVision(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/chat/completions", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)

		var requestBody CreateChatCompletionRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		assert.NoError(t, err)
		assert.Equal(t, "gpt-4o", requestBody.Model)
		assert.Len(t, requestBody.Messages, 1)
		assert.Equal(t, User, requestBody.Messages[0].Role)

		parts, err := requestBody.Messages[0].Content.AsMessageContent1()
		assert.NoError(t, err)
		assert.Len(t, parts, 2)

		textPart, err := parts[0].AsTextContentPart()
		assert.NoError(t, err)
		assert.Equal(t, "What is in this image?", textPart.Text)

		imagePart, err := parts[1].AsImageContentPart()
		assert.NoError(t, err)
		assert.Equal(t, "https://example.com/image.jpg", imagePart.ImageURL.URL)

		var responseMsg Message
		responseMsg.Role = Assistant
		err = responseMsg.Content.FromMessageContent0("The image shows a beautiful landscape with mountains and a lake.")
		assert.NoError(t, err)

		response := CreateChatCompletionResponse{
			ID:      "chat-vision-123",
			Object:  "chat.completion",
			Created: 1693672537,
			Model:   "gpt-4o",
			Choices: []ChatCompletionChoice{
				{
					Index:        0,
					Message:      responseMsg,
					FinishReason: Stop,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	var msg Message
	msg.Role = User

	contentParts := []ContentPart{}

	var textPart ContentPart
	err := textPart.FromTextContentPart(TextContentPart{
		Type: Text,
		Text: "What is in this image?",
	})
	require.NoError(t, err)
	contentParts = append(contentParts, textPart)

	var imagePart ContentPart
	err = imagePart.FromImageContentPart(ImageContentPart{
		Type: ImageContentPartTypeImageURL,
		ImageURL: ImageURL{
			URL: "https://example.com/image.jpg",
		},
	})
	require.NoError(t, err)
	contentParts = append(contentParts, imagePart)

	err = msg.Content.FromMessageContent1(contentParts)
	require.NoError(t, err)

	ctx := context.Background()
	response, err := client.GenerateContent(
		ctx,
		Openai,
		"gpt-4o",
		[]Message{msg},
	)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "chat-vision-123", response.ID)
	assert.Equal(t, "gpt-4o", response.Model)
	assert.Len(t, response.Choices, 1)

	responseText, err := response.Choices[0].Message.Content.AsMessageContent0()
	assert.NoError(t, err)
	assert.Contains(t, responseText, "landscape")
}
