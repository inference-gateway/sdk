package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://example.com", nil)
	assert.NotNil(t, client, "NewClient should return a non-nil client")
}

func TestListModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path, "Path should be /models")
		assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")

		response := ListModelsResponse{
			Object: stringPtr("list"),
			Data: &[]Model{
				{
					Id:      stringPtr("gpt-4o"),
					Object:  stringPtr("model"),
					Created: int64Ptr(1686935002),
					OwnedBy: stringPtr("openai"),
				},
				{
					Id:      stringPtr("llama-3.3-70b-versatile"),
					Object:  stringPtr("model"),
					Created: int64Ptr(1723651281),
					OwnedBy: stringPtr("groq"),
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)

	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	ctx := context.Background()
	models, err := client.ListModels(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, models)
	assert.Equal(t, "list", *models.Object)
	assert.Len(t, *models.Data, 2)
	assert.Equal(t, "gpt-4o", *(*models.Data)[0].Id)
	assert.Equal(t, "llama-3.3-70b-versatile", *(*models.Data)[1].Id)
}

func TestListProviderModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path, "Path should be /models")
		assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")
		assert.Equal(t, "openai", r.URL.Query().Get("provider"), "Provider should be specified in query")

		response := ListModelsResponse{
			Provider: providerPtr(Openai),
			Object:   stringPtr("list"),
			Data: &[]Model{
				{
					Id:      stringPtr("gpt-4o"),
					Object:  stringPtr("model"),
					Created: int64Ptr(1686935002),
					OwnedBy: stringPtr("openai"),
				},
				{
					Id:      stringPtr("gpt-4-turbo"),
					Object:  stringPtr("model"),
					Created: int64Ptr(1687882410),
					OwnedBy: stringPtr("openai"),
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	ctx := context.Background()
	models, err := client.ListProviderModels(ctx, Openai)

	assert.NoError(t, err)
	assert.NotNil(t, models)
	assert.Equal(t, Openai, *models.Provider)
	assert.Equal(t, "list", *models.Object)
	assert.Len(t, *models.Data, 2)
	assert.Equal(t, "gpt-4o", *(*models.Data)[0].Id)
	assert.Equal(t, "gpt-4-turbo", *(*models.Data)[1].Id)
}

func TestListProviderModels_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(Error{
			Error: stringPtr("Invalid API key"),
		})
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	ctx := context.Background()
	models, err := client.ListProviderModels(ctx, Groq)

	assert.Error(t, err)
	assert.Nil(t, models)
	assert.Contains(t, err.Error(), "API error")
}

func TestGenerateContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/chat/completions", r.URL.Path, "Path should be /v1/chat/completions")
		assert.Equal(t, http.MethodPost, r.Method, "Method should be POST")
		assert.Equal(t, "openai", r.URL.Query().Get("provider"), "Provider should be specified in query")

		var requestBody CreateChatCompletionRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		assert.NoError(t, err, "Should be able to decode request body")
		assert.Equal(t, "gpt-4o", requestBody.Model, "Model should match")
		assert.Len(t, requestBody.Messages, 2, "Should have 2 messages")
		assert.Equal(t, System, requestBody.Messages[0].Role, "First message should have system role")
		assert.Equal(t, User, requestBody.Messages[1].Role, "Second message should have user role")

		response := CreateChatCompletionResponse{
			Id:      "chat-12345",
			Object:  "chat.completion",
			Created: 1693672537,
			Model:   "gpt-4o",
			Choices: []ChatCompletionChoice{
				{
					Index: 0,
					Message: Message{
						Role:    Assistant,
						Content: "Go is a programming language designed by Google engineers in 2007. It's known for its simplicity, efficiency, and strong support for concurrency.",
					},
					FinishReason: Stop,
				},
			},
			Usage: &CompletionUsage{
				PromptTokens:     42,
				CompletionTokens: 25,
				TotalTokens:      67,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		assert.NoError(t, err, "Should be able to encode response")
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(baseURL, nil)

	ctx := context.Background()
	response, err := client.GenerateContent(
		ctx,
		Openai,
		"gpt-4o",
		[]Message{
			{
				Role:    System,
				Content: "You are a helpful assistant.",
			},
			{
				Role:    User,
				Content: "What is Go?",
			},
		},
	)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "chat-12345", response.Id)
	assert.Equal(t, "gpt-4o", response.Model)
	assert.Len(t, response.Choices, 1)
	assert.Equal(t, Assistant, response.Choices[0].Message.Role)
	assert.Contains(t, response.Choices[0].Message.Content, "Go is a programming language")
	assert.Equal(t, Stop, response.Choices[0].FinishReason)
	assert.NotNil(t, response.Usage)
	assert.Equal(t, int64(67), response.Usage.TotalTokens)
}

func TestGenerateContent_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(Error{
			Error: stringPtr("Invalid model specified"),
		})
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	ctx := context.Background()
	response, err := client.GenerateContent(
		ctx,
		Groq,
		"invalid-model",
		[]Message{
			{
				Role:    User,
				Content: "What is Go?",
			},
		},
	)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "API error")
	assert.Contains(t, err.Error(), "Invalid model")
}

func TestGenerateContentStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/chat/completions", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "ollama", r.URL.Query().Get("provider"))

		var requestBody CreateChatCompletionRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		assert.NoError(t, err)
		assert.Equal(t, "llama2", requestBody.Model)
		assert.True(t, *requestBody.Stream)

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatalf("Streaming not supported")
		}

		// Send stream responses as single lines
		chunk1 := `{"id": "chatcmpl-123","object": "chat.completion.chunk","created": 1698819810,"model": "llama2","choices": [{"delta": {"content": "Go"},"index": 0,"finish_reason": null}]}`
		chunk2 := `{"id": "chatcmpl-123","object": "chat.completion.chunk","created": 1698819810,"model": "llama2","choices": [{"delta": {"content": " is"},"index": 0,"finish_reason": null}]}`
		chunk3 := `{"id": "chatcmpl-123","object": "chat.completion.chunk","created": 1698819810,"model": "llama2","choices": [{"delta": {"content": " amazing"},"index": 0,"finish_reason": "stop"}]}`

		fmt.Fprintf(w, "data: %s\n\n", chunk1)
		flusher.Flush()

		fmt.Fprintf(w, "data: %s\n\n", chunk2)
		flusher.Flush()

		fmt.Fprintf(w, "data: %s\n\n", chunk3)
		flusher.Flush()

		fmt.Fprintf(w, "data: [DONE]\n\n")
		flusher.Flush()
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(baseURL, nil)

	ctx := context.Background()
	eventCh, err := client.GenerateContentStream(
		ctx,
		Ollama,
		"llama2",
		[]Message{
			{
				Role:    System,
				Content: "You are a helpful assistant.",
			},
			{
				Role:    User,
				Content: "What is Go?",
			},
		},
	)

	assert.NoError(t, err)
	assert.NotNil(t, eventCh)

	var content string
	var eventCount int
	var streamEndReceived bool

	for event := range eventCh {
		eventCount++

		if event.Event != nil && *event.Event == StreamEnd {
			streamEndReceived = true
			continue
		}

		if event.Event != nil && *event.Event == ContentDelta && event.Data != nil {
			var streamResponse CreateChatCompletionStreamResponse
			err := json.Unmarshal(*event.Data, &streamResponse)
			if err != nil {
				continue
			}

			for _, choice := range streamResponse.Choices {
				if choice.Delta.Content != "" {
					content += choice.Delta.Content
				}
			}
		}
	}

	assert.Equal(t, "Go is amazing", content)
	assert.Equal(t, 4, eventCount) // 3 content chunks + DONE event
	assert.True(t, streamEndReceived)
}

func TestGenerateContentStream_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(Error{
			Error: stringPtr("Invalid model for streaming"),
		})
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	ctx := context.Background()
	eventCh, err := client.GenerateContentStream(
		ctx,
		Groq,
		"invalid-model",
		[]Message{
			{
				Role:    User,
				Content: "What is Go?",
			},
		},
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "stream request failed")

	_, open := <-eventCh
	assert.False(t, open, "Channel should be closed on error")
}

func TestHealthCheck(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/health", r.URL.Path, "Path should be /health")
		assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	ctx := context.Background()
	err := client.HealthCheck(ctx)

	assert.NoError(t, err)
}

func TestHealthCheck_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	ctx := context.Background()
	err := client.HealthCheck(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "health check failed")
}

func stringPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}

func providerPtr(p Provider) *Provider {
	return &p
}
