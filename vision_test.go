package sdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVisionMessage_TextContent(t *testing.T) {
	// Test creating a message with simple text content
	msg := Message{
		Role: User,
	}

	err := msg.Content.FromMessageContent0("What is in this image?")
	require.NoError(t, err)

	// Verify we can retrieve it
	text, err := msg.Content.AsMessageContent0()
	require.NoError(t, err)
	assert.Equal(t, "What is in this image?", text)
}

func TestVisionMessage_ImageURLContent(t *testing.T) {
	// Test creating a message with image URL content
	msg := Message{
		Role: User,
	}

	// Create content parts with text and image
	contentParts := []ContentPart{}

	// Add text part
	var textPart ContentPart
	err := textPart.FromTextContentPart(TextContentPart{
		Type: Text,
		Text: "What is in this image?",
	})
	require.NoError(t, err)
	contentParts = append(contentParts, textPart)

	// Add image part
	var imagePart ContentPart
	err = imagePart.FromImageContentPart(ImageContentPart{
		Type: ImageUrl,
		ImageUrl: ImageURL{
			Url: "https://example.com/image.jpg",
			Detail: detailPtr(Auto),
		},
	})
	require.NoError(t, err)
	contentParts = append(contentParts, imagePart)

	// Set the content
	err = msg.Content.FromMessageContent1(contentParts)
	require.NoError(t, err)

	// Verify we can retrieve it
	parts, err := msg.Content.AsMessageContent1()
	require.NoError(t, err)
	assert.Len(t, parts, 2)

	// Verify text part
	textPartRetrieved, err := parts[0].AsTextContentPart()
	require.NoError(t, err)
	assert.Equal(t, Text, textPartRetrieved.Type)
	assert.Equal(t, "What is in this image?", textPartRetrieved.Text)

	// Verify image part
	imagePartRetrieved, err := parts[1].AsImageContentPart()
	require.NoError(t, err)
	assert.Equal(t, ImageUrl, imagePartRetrieved.Type)
	assert.Equal(t, "https://example.com/image.jpg", imagePartRetrieved.ImageUrl.Url)
	assert.NotNil(t, imagePartRetrieved.ImageUrl.Detail)
	assert.Equal(t, Auto, *imagePartRetrieved.ImageUrl.Detail)
}

func TestVisionMessage_DataURLImage(t *testing.T) {
	// Test creating a message with base64 encoded image
	msg := Message{
		Role: User,
	}

	dataURL := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD..."

	contentParts := []ContentPart{}

	var imagePart ContentPart
	err := imagePart.FromImageContentPart(ImageContentPart{
		Type: ImageUrl,
		ImageUrl: ImageURL{
			Url: dataURL,
			Detail: detailPtr(High),
		},
	})
	require.NoError(t, err)
	contentParts = append(contentParts, imagePart)

	err = msg.Content.FromMessageContent1(contentParts)
	require.NoError(t, err)

	// Verify
	parts, err := msg.Content.AsMessageContent1()
	require.NoError(t, err)
	assert.Len(t, parts, 1)

	imagePartRetrieved, err := parts[0].AsImageContentPart()
	require.NoError(t, err)
	assert.Equal(t, dataURL, imagePartRetrieved.ImageUrl.Url)
	assert.Equal(t, High, *imagePartRetrieved.ImageUrl.Detail)
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

		// Verify the content is an array (multimodal)
		parts, err := requestBody.Messages[0].Content.AsMessageContent1()
		assert.NoError(t, err)
		assert.Len(t, parts, 2)

		// Verify text part
		textPart, err := parts[0].AsTextContentPart()
		assert.NoError(t, err)
		assert.Equal(t, "What is in this image?", textPart.Text)

		// Verify image part
		imagePart, err := parts[1].AsImageContentPart()
		assert.NoError(t, err)
		assert.Equal(t, "https://example.com/image.jpg", imagePart.ImageUrl.Url)

		// Send response with text content
		var responseMsg Message
		responseMsg.Role = Assistant
		err = responseMsg.Content.FromMessageContent0("The image shows a beautiful landscape with mountains and a lake.")
		assert.NoError(t, err)

		response := CreateChatCompletionResponse{
			Id:      "chat-vision-123",
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

	// Create a vision message
	var msg Message
	msg.Role = User

	contentParts := []ContentPart{}

	// Add text
	var textPart ContentPart
	err := textPart.FromTextContentPart(TextContentPart{
		Type: Text,
		Text: "What is in this image?",
	})
	require.NoError(t, err)
	contentParts = append(contentParts, textPart)

	// Add image
	var imagePart ContentPart
	err = imagePart.FromImageContentPart(ImageContentPart{
		Type: ImageUrl,
		ImageUrl: ImageURL{
			Url: "https://example.com/image.jpg",
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
	assert.Equal(t, "chat-vision-123", response.Id)
	assert.Equal(t, "gpt-4o", response.Model)
	assert.Len(t, response.Choices, 1)

	// Verify response content
	responseText, err := response.Choices[0].Message.Content.AsMessageContent0()
	assert.NoError(t, err)
	assert.Contains(t, responseText, "landscape")
}

// Helper function for creating image detail pointer
func detailPtr(d ImageURLDetail) *ImageURLDetail {
	return &d
}
