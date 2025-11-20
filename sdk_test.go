package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient(&ClientOptions{
		BaseURL: "http://localhost:8080/v1",
	})
	assert.NotNil(t, client, "NewClient should return a non-nil client")
}

func TestListModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/models", r.URL.Path, "Path should be /v1/models")
		assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")

		response := ListModelsResponse{
			Object: "list",
			Data: []Model{
				{
					Id:       "openai/gpt-4o",
					Object:   "model",
					Created:  1686935002,
					OwnedBy:  "openai",
					ServedBy: Openai,
				},
				{
					Id:       "groq/llama-3.3-70b-versatile",
					Object:   "model",
					Created:  1723651281,
					OwnedBy:  "groq",
					ServedBy: Groq,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)

	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	ctx := context.Background()
	models, err := client.ListModels(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, models)
	assert.Equal(t, "list", models.Object)
	assert.Len(t, models.Data, 2)
	assert.Equal(t, "openai/gpt-4o", models.Data[0].Id)
	assert.Equal(t, "groq/llama-3.3-70b-versatile", models.Data[1].Id)
}

func TestListProviderModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/models", r.URL.Path, "Path should be /v1/models")
		assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")
		assert.Equal(t, "openai", r.URL.Query().Get("provider"), "Provider should be specified in query")

		response := ListModelsResponse{
			Provider: providerPtr(Openai),
			Object:   "list",
			Data: []Model{
				{
					Id:       "openai/gpt-4o",
					Object:   "model",
					Created:  1686935002,
					OwnedBy:  "openai",
					ServedBy: Openai,
				},
				{
					Id:       "openai/gpt-4-turbo",
					Object:   "model",
					Created:  1687882410,
					OwnedBy:  "openai",
					ServedBy: Openai,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	ctx := context.Background()
	models, err := client.ListProviderModels(ctx, Openai)

	assert.NoError(t, err)
	assert.NotNil(t, models)
	assert.Equal(t, Openai, *models.Provider)
	assert.Equal(t, "list", models.Object)
	assert.Len(t, models.Data, 2)
	assert.Equal(t, "openai/gpt-4o", models.Data[0].Id)
	assert.Equal(t, "openai/gpt-4-turbo", models.Data[1].Id)
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

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	ctx := context.Background()
	models, err := client.ListProviderModels(ctx, Groq)

	assert.Error(t, err)
	assert.Nil(t, models)
	assert.Contains(t, err.Error(), "API error")
}

func TestListTools(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/mcp/tools", r.URL.Path, "Path should be /v1/mcp/tools")
		assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")

		response := ListToolsResponse{
			Object: "list",
			Data: []MCPTool{
				{
					Name:        "read_file",
					Description: "Read content from a file",
					Server:      "http://mcp-filesystem-server:8083/mcp",
					InputSchema: &map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"file_path": map[string]interface{}{
								"type":        "string",
								"description": "Path to the file to read",
							},
						},
						"required": []string{"file_path"},
					},
				},
				{
					Name:        "write_file",
					Description: "Write content to a file",
					Server:      "http://mcp-filesystem-server:8083/mcp",
					InputSchema: &map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"file_path": map[string]interface{}{
								"type":        "string",
								"description": "Path to the file to write",
							},
							"content": map[string]interface{}{
								"type":        "string",
								"description": "Content to write to the file",
							},
						},
						"required": []string{"file_path", "content"},
					},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	ctx := context.Background()
	tools, err := client.ListTools(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, tools)
	assert.Equal(t, "list", tools.Object)
	assert.Len(t, tools.Data, 2)

	assert.Equal(t, "read_file", tools.Data[0].Name)
	assert.Equal(t, "Read content from a file", tools.Data[0].Description)
	assert.Equal(t, "http://mcp-filesystem-server:8083/mcp", tools.Data[0].Server)
	assert.NotNil(t, tools.Data[0].InputSchema)
}

func TestListTools_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		response := map[string]interface{}{
			"error": "MCP tools endpoint is not exposed. Set EXPOSE_MCP=true to enable.",
		}
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	ctx := context.Background()
	tools, err := client.ListTools(ctx)

	assert.Error(t, err)
	assert.Nil(t, tools)
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
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

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

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

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

		chunk1 := `{"id": "chatcmpl-123","object": "chat.completion.chunk","created": 1698819810,"model": "llama2","choices": [{"delta": {"content": "Go"},"index": 0,"finish_reason": null}]}`
		chunk2 := `{"id": "chatcmpl-123","object": "chat.completion.chunk","created": 1698819810,"model": "llama2","choices": [{"delta": {"content": " is"},"index": 0,"finish_reason": null}]}`
		chunk3 := `{"id": "chatcmpl-123","object": "chat.completion.chunk","created": 1698819810,"model": "llama2","choices": [{"delta": {"content": " amazing"},"index": 0,"finish_reason": "stop"}]}`

		_, err = fmt.Fprintf(w, "data: %s\n\n", chunk1)
		require.NoError(t, err)
		flusher.Flush()

		_, err = fmt.Fprintf(w, "data: %s\n\n", chunk2)
		require.NoError(t, err)
		flusher.Flush()

		_, err = fmt.Fprintf(w, "data: %s\n\n", chunk3)
		require.NoError(t, err)
		flusher.Flush()

		_, err = fmt.Fprintf(w, "data: [DONE]\n\n")
		require.NoError(t, err)
		flusher.Flush()
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})
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
	assert.Equal(t, 4, eventCount)
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

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

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
	assert.Contains(t, err.Error(), "API stream error")

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

	baseURL := server.URL
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	ctx := context.Background()
	err := client.HealthCheck(ctx)

	assert.NoError(t, err)
}

func TestHealthCheck_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
		RetryConfig: &RetryConfig{
			Enabled: false,
		},
	})

	ctx := context.Background()
	err := client.HealthCheck(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "health check failed")
}

func TestWithOptions(t *testing.T) {
	testCases := []struct {
		name            string
		provider        Provider
		model           string
		messages        []Message
		options         *CreateChatCompletionRequest
		isStreaming     bool
		expectedOptions func(t *testing.T, req CreateChatCompletionRequest)
		mockResponse    func(t *testing.T) interface{}
	}{
		{
			name:     "Basic content with no options",
			provider: Openai,
			model:    "openai/gpt-4o",
			messages: []Message{
				{Role: User, Content: "Hello"},
			},
			options:     nil,
			isStreaming: false,
			expectedOptions: func(t *testing.T, req CreateChatCompletionRequest) {
				assert.Equal(t, "openai/gpt-4o", req.Model)
				assert.Equal(t, 1, len(req.Messages))
				assert.NotNil(t, req.Stream)
				assert.False(t, *req.Stream)
			},
			mockResponse: func(t *testing.T) interface{} {
				return CreateChatCompletionResponse{
					Id:      "test-1",
					Object:  "chat.completion",
					Created: 1693672537,
					Model:   "openai/gpt-4o",
					Choices: []ChatCompletionChoice{
						{
							Index: 0,
							Message: Message{
								Role:    Assistant,
								Content: "Hello there!",
							},
							FinishReason: Stop,
						},
					},
				}
			},
		},
		{
			name:     "Content with reasoning format parsed",
			provider: Anthropic,
			model:    "anthropic/claude-3-opus-20240229",
			messages: []Message{
				{Role: User, Content: "What is the square root of 144?"},
			},
			options: &CreateChatCompletionRequest{
				ReasoningFormat: stringPtr("parsed"),
			},
			isStreaming: false,
			expectedOptions: func(t *testing.T, req CreateChatCompletionRequest) {
				assert.Equal(t, "anthropic/claude-3-opus-20240229", req.Model)
				assert.NotNil(t, req.ReasoningFormat)
				assert.Equal(t, "parsed", *req.ReasoningFormat)
				assert.NotNil(t, req.Stream)
				assert.False(t, *req.Stream)
			},
			mockResponse: func(t *testing.T) interface{} {
				reasoningContent := "I need to calculate the square root of 144. The square root of a number is a value that, when multiplied by itself, gives the original number. For 144, the square root is 12 because 12 × 12 = 144."
				return CreateChatCompletionResponse{
					Id:      "test-2",
					Object:  "chat.completion",
					Created: 1693672537,
					Model:   "anthropic/claude-3-opus-20240229",
					Choices: []ChatCompletionChoice{
						{
							Index: 0,
							Message: Message{
								Role:             Assistant,
								Content:          "The square root of 144 is 12.",
								ReasoningContent: &reasoningContent,
								Reasoning:        &reasoningContent,
							},
							FinishReason: Stop,
						},
					},
				}
			},
		},
		{
			name:     "Content with reasoning format raw",
			provider: Anthropic,
			model:    "anthropic/claude-3-opus-20240229",
			messages: []Message{
				{Role: User, Content: "What is the square root of 144?"},
			},
			options: &CreateChatCompletionRequest{
				ReasoningFormat: stringPtr("raw"),
			},
			isStreaming: false,
			expectedOptions: func(t *testing.T, req CreateChatCompletionRequest) {
				assert.Equal(t, "anthropic/claude-3-opus-20240229", req.Model)
				assert.NotNil(t, req.ReasoningFormat)
				assert.Equal(t, "raw", *req.ReasoningFormat)
				assert.NotNil(t, req.Stream)
				assert.False(t, *req.Stream)
			},
			mockResponse: func(t *testing.T) interface{} {
				return CreateChatCompletionResponse{
					Id:      "test-3",
					Object:  "chat.completion",
					Created: 1693672537,
					Model:   "anthropic/claude-3-opus-20240229",
					Choices: []ChatCompletionChoice{
						{
							Index: 0,
							Message: Message{
								Role:    Assistant,
								Content: "<think>\nI need to calculate the square root of 144. \n\nThe square root of a number is a value that, when multiplied by itself, gives the original number.\n\nFor 144:\n√144 = x means x² = 144\n\n12² = 144 because 12 × 12 = 144\n\nTherefore, √144 = 12\n</think>\n\nThe square root of 144 is 12.",
							},
							FinishReason: Stop,
						},
					},
				}
			},
		},
		{
			name:     "Streaming content with options",
			provider: Ollama,
			model:    "ollama/llama2",
			messages: []Message{
				{Role: User, Content: "Tell me about streaming"},
			},
			options: &CreateChatCompletionRequest{
				Stream: boolPtr(false),
			},
			isStreaming: true,
			expectedOptions: func(t *testing.T, req CreateChatCompletionRequest) {
				assert.Equal(t, "ollama/llama2", req.Model)
				assert.NotNil(t, req.Stream)
				assert.True(t, *req.Stream)
			},
			mockResponse: func(t *testing.T) interface{} {
				return nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/v1/chat/completions", r.URL.Path)
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, string(tc.provider), r.URL.Query().Get("provider"))

				var requestBody CreateChatCompletionRequest
				err := json.NewDecoder(r.Body).Decode(&requestBody)
				assert.NoError(t, err)

				tc.expectedOptions(t, requestBody)

				if tc.isStreaming {
					w.Header().Set("Content-Type", "text/event-stream")
					w.Header().Set("Cache-Control", "no-cache")
					w.Header().Set("Connection", "keep-alive")

					flusher, ok := w.(http.Flusher)
					if !ok {
						t.Fatalf("Streaming not supported")
					}

					chunk := `{"id": "chatcmpl-123","object": "chat.completion.chunk","created": 1698819810,"model": "ollama/llama2","choices": [{"delta": {"content": "Streaming content"},"index": 0,"finish_reason": null}]}`
					_, err = fmt.Fprintf(w, "data: %s\n\n", chunk)
					require.NoError(t, err)
					flusher.Flush()

					_, err = fmt.Fprintf(w, "data: [DONE]\n\n")
					require.NoError(t, err)
					flusher.Flush()
				} else {
					resp := tc.mockResponse(t)
					w.Header().Set("Content-Type", "application/json")
					err = json.NewEncoder(w).Encode(resp)
					assert.NoError(t, err)
				}
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			if tc.options != nil {
				client = client.WithOptions(tc.options)
			}

			ctx := context.Background()

			if tc.isStreaming {
				events, err := client.GenerateContentStream(ctx, tc.provider, tc.model, tc.messages)
				assert.NoError(t, err)
				assert.NotNil(t, events)

				var content string
				var eventCount int

				for event := range events {
					eventCount++

					if event.Event != nil && *event.Event == ContentDelta {
						var streamResponse CreateChatCompletionStreamResponse
						err := json.Unmarshal(*event.Data, &streamResponse)
						if err != nil {
							continue
						}

						for _, choice := range streamResponse.Choices {
							content += choice.Delta.Content
						}
					}
				}

				assert.Equal(t, "Streaming content", content)
				assert.Equal(t, 2, eventCount)
			} else {
				result, err := client.GenerateContent(ctx, tc.provider, tc.model, tc.messages)
				assert.NoError(t, err)
				assert.NotNil(t, result)

				if tc.name == "Content with reasoning format" {
					assert.NotNil(t, result.Choices[0].Message.ReasoningContent)
					assert.NotNil(t, result.Choices[0].Message.Reasoning)
					assert.Contains(t, *result.Choices[0].Message.Reasoning, "square root of 144")
				}
			}
		})
	}
}

func TestWithHeaders(t *testing.T) {
	tests := []struct {
		name            string
		initialHeaders  map[string]string
		withHeaders     map[string]string
		singleHeaders   map[string]string
		expectedHeaders map[string]string
	}{
		{
			name: "Set headers from ClientOptions",
			initialHeaders: map[string]string{
				"X-Initial-Header": "initial-value",
				"User-Agent":       "test-app/1.0",
			},
			expectedHeaders: map[string]string{
				"X-Initial-Header": "initial-value",
				"User-Agent":       "test-app/1.0",
			},
		},
		{
			name: "WithHeaders sets multiple headers",
			withHeaders: map[string]string{
				"X-Custom-Header-1": "value1",
				"X-Custom-Header-2": "value2",
			},
			expectedHeaders: map[string]string{
				"X-Custom-Header-1": "value1",
				"X-Custom-Header-2": "value2",
			},
		},
		{
			name: "WithHeader sets single header",
			singleHeaders: map[string]string{
				"X-Single-Header": "single-value",
			},
			expectedHeaders: map[string]string{
				"X-Single-Header": "single-value",
			},
		},
		{
			name: "Mixed headers from options and methods",
			initialHeaders: map[string]string{
				"X-Initial": "initial",
			},
			withHeaders: map[string]string{
				"X-Multi-1": "multi1",
				"X-Multi-2": "multi2",
			},
			singleHeaders: map[string]string{
				"X-Single": "single",
			},
			expectedHeaders: map[string]string{
				"X-Initial": "initial",
				"X-Multi-1": "multi1",
				"X-Multi-2": "multi2",
				"X-Single":  "single",
			},
		},
		{
			name: "Headers override",
			initialHeaders: map[string]string{
				"X-Override": "initial",
			},
			withHeaders: map[string]string{
				"X-Override": "withHeaders",
			},
			singleHeaders: map[string]string{
				"X-Override": "withHeader",
			},
			expectedHeaders: map[string]string{
				"X-Override": "withHeader",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for key, expectedValue := range tt.expectedHeaders {
					actualValue := r.Header.Get(key)
					assert.Equal(t, expectedValue, actualValue, "Header %s should have value %s", key, expectedValue)
				}

				response := ListModelsResponse{Object: "list", Data: []Model{}}
				w.Header().Set("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(response)
				assert.NoError(t, err)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
				Headers: tt.initialHeaders,
			})

			if tt.withHeaders != nil {
				client = client.WithHeaders(tt.withHeaders)
			}

			for key, value := range tt.singleHeaders {
				client = client.WithHeader(key, value)
			}

			ctx := context.Background()
			_, err := client.ListModels(ctx)
			assert.NoError(t, err)
		})
	}
}

func TestHeadersInAllRequests(t *testing.T) {
	customHeaders := map[string]string{
		"X-API-Version": "v1.0",
		"X-Client-ID":   "test-client",
	}

	testCases := []struct {
		name     string
		endpoint string
		makeCall func(client Client) error
	}{
		{
			name:     "ListModels",
			endpoint: "/v1/models",
			makeCall: func(client Client) error {
				_, err := client.ListModels(context.Background())
				return err
			},
		},
		{
			name:     "ListProviderModels",
			endpoint: "/v1/models",
			makeCall: func(client Client) error {
				_, err := client.ListProviderModels(context.Background(), Openai)
				return err
			},
		},
		{
			name:     "ListTools",
			endpoint: "/v1/mcp/tools",
			makeCall: func(client Client) error {
				_, err := client.ListTools(context.Background())
				return err
			},
		},
		{
			name:     "HealthCheck",
			endpoint: "/v1/health",
			makeCall: func(client Client) error {
				return client.HealthCheck(context.Background())
			},
		},
		{
			name:     "GenerateContent",
			endpoint: "/v1/chat/completions",
			makeCall: func(client Client) error {
				_, err := client.GenerateContent(context.Background(), Openai, "gpt-4o", []Message{{Role: User, Content: "test"}})
				return err
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for key, expectedValue := range customHeaders {
					actualValue := r.Header.Get(key)
					assert.Equal(t, expectedValue, actualValue, "Header %s should be present in %s request", key, tc.name)
				}

				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/v1/models":
					response := ListModelsResponse{Object: "list", Data: []Model{}}
					err := json.NewEncoder(w).Encode(response)
					assert.NoError(t, err)
				case "/v1/mcp/tools":
					response := ListToolsResponse{Data: []MCPTool{}}
					err := json.NewEncoder(w).Encode(response)
					assert.NoError(t, err)
				case "/v1/health":
					w.WriteHeader(http.StatusOK)
				case "/v1/chat/completions":
					response := CreateChatCompletionResponse{
						Id:      "test",
						Object:  "chat.completion",
						Created: 123456789,
						Model:   "gpt-4o",
						Choices: []ChatCompletionChoice{{
							Index:        0,
							FinishReason: Stop,
							Message:      Message{Role: Assistant, Content: "test response"},
						}},
					}
					err := json.NewEncoder(w).Encode(response)
					assert.NoError(t, err)
				}
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
				Headers: customHeaders,
			})

			err := tc.makeCall(client)
			assert.NoError(t, err)
		})
	}
}

func TestWithMiddlewareOptions(t *testing.T) {
	tests := []struct {
		name            string
		middlewareOpts  *MiddlewareOptions
		expectedHeaders map[string]string
	}{
		{
			name:           "nil middleware options",
			middlewareOpts: nil,
			expectedHeaders: map[string]string{
				"X-MCP-Bypass":      "",
				"X-A2A-Bypass":      "",
				"X-Direct-Provider": "",
			},
		},
		{
			name: "skip MCP only",
			middlewareOpts: &MiddlewareOptions{
				SkipMCP: true,
			},
			expectedHeaders: map[string]string{
				"X-MCP-Bypass":      "true",
				"X-A2A-Bypass":      "",
				"X-Direct-Provider": "",
			},
		},
		{
			name: "skip A2A only",
			middlewareOpts: &MiddlewareOptions{
				SkipA2A: true,
			},
			expectedHeaders: map[string]string{
				"X-MCP-Bypass":      "",
				"X-A2A-Bypass":      "true",
				"X-Direct-Provider": "",
			},
		},
		{
			name: "direct provider only",
			middlewareOpts: &MiddlewareOptions{
				DirectProvider: true,
			},
			expectedHeaders: map[string]string{
				"X-MCP-Bypass":      "",
				"X-A2A-Bypass":      "",
				"X-Direct-Provider": "true",
			},
		},
		{
			name: "all middleware options enabled",
			middlewareOpts: &MiddlewareOptions{
				SkipMCP:        true,
				SkipA2A:        true,
				DirectProvider: true,
			},
			expectedHeaders: map[string]string{
				"X-MCP-Bypass":      "true",
				"X-A2A-Bypass":      "true",
				"X-Direct-Provider": "true",
			},
		},
		{
			name: "mixed middleware options",
			middlewareOpts: &MiddlewareOptions{
				SkipMCP:        true,
				SkipA2A:        false,
				DirectProvider: true,
			},
			expectedHeaders: map[string]string{
				"X-MCP-Bypass":      "true",
				"X-A2A-Bypass":      "",
				"X-Direct-Provider": "true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for header, expectedValue := range tt.expectedHeaders {
					actualValue := r.Header.Get(header)
					if expectedValue == "" {
						assert.Empty(t, actualValue, "Header %s should be empty or not present", header)
					} else {
						assert.Equal(t, expectedValue, actualValue, "Header %s should have value %s", header, expectedValue)
					}
				}

				response := ListModelsResponse{Object: "list", Data: []Model{}}
				w.Header().Set("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(response)
				assert.NoError(t, err)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			client = client.WithMiddlewareOptions(tt.middlewareOpts)

			ctx := context.Background()
			_, err := client.ListModels(ctx)
			assert.NoError(t, err)
		})
	}
}

func TestMiddlewareOptionsInAllRequests(t *testing.T) {
	middlewareOpts := &MiddlewareOptions{
		SkipMCP:        true,
		SkipA2A:        true,
		DirectProvider: true,
	}

	expectedHeaders := map[string]string{
		"X-MCP-Bypass":      "true",
		"X-A2A-Bypass":      "true",
		"X-Direct-Provider": "true",
	}

	testCases := []struct {
		name     string
		endpoint string
		makeCall func(client Client) error
	}{
		{
			name:     "ListModels",
			endpoint: "/v1/models",
			makeCall: func(client Client) error {
				_, err := client.ListModels(context.Background())
				return err
			},
		},
		{
			name:     "ListProviderModels",
			endpoint: "/v1/models",
			makeCall: func(client Client) error {
				_, err := client.ListProviderModels(context.Background(), Openai)
				return err
			},
		},
		{
			name:     "ListTools",
			endpoint: "/v1/mcp/tools",
			makeCall: func(client Client) error {
				_, err := client.ListTools(context.Background())
				return err
			},
		},
		{
			name:     "GenerateContent",
			endpoint: "/v1/chat/completions",
			makeCall: func(client Client) error {
				_, err := client.GenerateContent(context.Background(), Openai, "gpt-4o", []Message{
					{Role: User, Content: "test"},
				})
				return err
			},
		},
		{
			name:     "HealthCheck",
			endpoint: "/v1/health",
			makeCall: func(client Client) error {
				return client.HealthCheck(context.Background())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for header, expectedValue := range expectedHeaders {
					actualValue := r.Header.Get(header)
					assert.Equal(t, expectedValue, actualValue, "Header %s should have value %s", header, expectedValue)
				}

				switch r.URL.Path {
				case "/v1/models":
					response := ListModelsResponse{Object: "list", Data: []Model{}}
					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(response)
					assert.NoError(t, err)
				case "/v1/mcp/tools":
					response := ListToolsResponse{Data: []MCPTool{}}
					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(response)
					assert.NoError(t, err)
				case "/v1/chat/completions":
					response := CreateChatCompletionResponse{
						Id:      "test-id",
						Object:  "chat.completion",
						Created: 1234567890,
						Model:   "gpt-4o",
						Choices: []ChatCompletionChoice{
							{
								Index: 0,
								Message: Message{
									Role:    Assistant,
									Content: "Test response",
								},
								FinishReason: "stop",
							},
						},
					}
					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(response)
					assert.NoError(t, err)
				case "/v1/health":
					w.WriteHeader(http.StatusOK)
				}
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			client = client.WithMiddlewareOptions(middlewareOpts)

			err := tc.makeCall(client)
			assert.NoError(t, err)
		})
	}
}

func TestMiddlewareOptionsChaining(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "true", r.Header.Get("X-MCP-Bypass"))
		assert.Equal(t, "true", r.Header.Get("X-A2A-Bypass"))
		assert.Equal(t, "", r.Header.Get("X-Direct-Provider"))

		assert.Equal(t, "custom-value", r.Header.Get("X-Custom-Header"))

		response := ListModelsResponse{Object: "list", Data: []Model{}}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
	})

	client = client.
		WithHeader("X-Custom-Header", "custom-value").
		WithMiddlewareOptions(&MiddlewareOptions{
			SkipMCP:        true,
			DirectProvider: true,
		}).
		WithMiddlewareOptions(&MiddlewareOptions{
			SkipMCP: true,
			SkipA2A: true,
		})

	ctx := context.Background()
	_, err := client.ListModels(ctx)
	assert.NoError(t, err)
}

func stringPtr(s string) *string {
	return &s
}

func providerPtr(p Provider) *Provider {
	return &p
}

func TestRetryLogic(t *testing.T) {
	tests := []struct {
		name          string
		retryConfig   *RetryConfig
		statusCodes   []int
		networkErrors []bool
		expectedCalls int
		expectedError bool
	}{
		{
			name: "no retry on success",
			retryConfig: &RetryConfig{
				Enabled:           true,
				MaxAttempts:       3,
				InitialBackoffSec: 1,
				MaxBackoffSec:     10,
				BackoffMultiplier: 2,
			},
			statusCodes:   []int{200},
			networkErrors: []bool{false},
			expectedCalls: 1,
			expectedError: false,
		},
		{
			name: "retry on 500 error",
			retryConfig: &RetryConfig{
				Enabled:           true,
				MaxAttempts:       3,
				InitialBackoffSec: 1,
				MaxBackoffSec:     10,
				BackoffMultiplier: 2,
			},
			statusCodes:   []int{500, 500, 200},
			networkErrors: []bool{false, false, false},
			expectedCalls: 3,
			expectedError: false,
		},
		{
			name: "retry on 503 error",
			retryConfig: &RetryConfig{
				Enabled:           true,
				MaxAttempts:       3,
				InitialBackoffSec: 1,
				MaxBackoffSec:     10,
				BackoffMultiplier: 2,
			},
			statusCodes:   []int{503, 503, 200},
			networkErrors: []bool{false, false, false},
			expectedCalls: 3,
			expectedError: false,
		},
		{
			name: "retry on 429 error",
			retryConfig: &RetryConfig{
				Enabled:           true,
				MaxAttempts:       3,
				InitialBackoffSec: 1,
				MaxBackoffSec:     10,
				BackoffMultiplier: 2,
			},
			statusCodes:   []int{429, 429, 200},
			networkErrors: []bool{false, false, false},
			expectedCalls: 3,
			expectedError: false,
		},
		{
			name: "max attempts exhausted",
			retryConfig: &RetryConfig{
				Enabled:           true,
				MaxAttempts:       3,
				InitialBackoffSec: 1,
				MaxBackoffSec:     10,
				BackoffMultiplier: 2,
			},
			statusCodes:   []int{500, 500, 500},
			networkErrors: []bool{false, false, false},
			expectedCalls: 3,
			expectedError: true,
		},
		{
			name: "no retry on 400 error",
			retryConfig: &RetryConfig{
				Enabled:           true,
				MaxAttempts:       3,
				InitialBackoffSec: 1,
				MaxBackoffSec:     10,
				BackoffMultiplier: 2,
			},
			statusCodes:   []int{400},
			networkErrors: []bool{false},
			expectedCalls: 1,
			expectedError: true,
		},
		{
			name: "retry disabled",
			retryConfig: &RetryConfig{
				Enabled: false,
			},
			statusCodes:   []int{500},
			networkErrors: []bool{false},
			expectedCalls: 1,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if callCount < len(tt.statusCodes) {
					w.WriteHeader(tt.statusCodes[callCount])
					if tt.statusCodes[callCount] == 200 {
						response := ListModelsResponse{Object: "list", Data: []Model{}}
						err := json.NewEncoder(w).Encode(response)
						assert.NoError(t, err)
					} else {
						err := json.NewEncoder(w).Encode(Error{Error: stringPtr("Server error")})
						assert.NoError(t, err)
					}
				}
				callCount++
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL:     baseURL,
				RetryConfig: tt.retryConfig,
			})

			ctx := context.Background()
			_, err := client.ListModels(ctx)

			assert.Equal(t, tt.expectedCalls, callCount)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRetryConfigDefaults(t *testing.T) {
	client := NewClient(&ClientOptions{
		BaseURL: "http://localhost:8080/v1",
	})

	clientImpl, ok := client.(*clientImpl)
	require.True(t, ok)
	require.NotNil(t, clientImpl.retryConfig)

	assert.True(t, clientImpl.retryConfig.Enabled)
	assert.Equal(t, 3, clientImpl.retryConfig.MaxAttempts)
	assert.Equal(t, 2, clientImpl.retryConfig.InitialBackoffSec)
	assert.Equal(t, 30, clientImpl.retryConfig.MaxBackoffSec)
	assert.Equal(t, 2, clientImpl.retryConfig.BackoffMultiplier)
}

func TestCustomRetryConfig(t *testing.T) {
	customConfig := &RetryConfig{
		Enabled:           false,
		MaxAttempts:       5,
		InitialBackoffSec: 1,
		MaxBackoffSec:     60,
		BackoffMultiplier: 3,
	}

	client := NewClient(&ClientOptions{
		BaseURL:     "http://localhost:8080/v1",
		RetryConfig: customConfig,
	})

	clientImpl, ok := client.(*clientImpl)
	require.True(t, ok)
	require.NotNil(t, clientImpl.retryConfig)

	assert.False(t, clientImpl.retryConfig.Enabled)
	assert.Equal(t, 5, clientImpl.retryConfig.MaxAttempts)
	assert.Equal(t, 1, clientImpl.retryConfig.InitialBackoffSec)
	assert.Equal(t, 60, clientImpl.retryConfig.MaxBackoffSec)
	assert.Equal(t, 3, clientImpl.retryConfig.BackoffMultiplier)
}

func TestCalculateBackoff(t *testing.T) {
	config := &RetryConfig{
		InitialBackoffSec: 2,
		MaxBackoffSec:     30,
		BackoffMultiplier: 2,
	}

	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{0, 0},
		{1, 2 * time.Second},
		{2, 4 * time.Second},
		{3, 8 * time.Second},
		{4, 16 * time.Second},
		{5, 30 * time.Second},
		{6, 30 * time.Second},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("attempt_%d", tt.attempt), func(t *testing.T) {
			result := calculateBackoff(tt.attempt, config)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseRetryAfter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		ok       bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
			ok:       false,
		},
		{
			name:     "seconds as integer",
			input:    "120",
			expected: 120 * time.Second,
			ok:       true,
		},
		{
			name:     "seconds as decimal",
			input:    "1.5",
			expected: 1500 * time.Millisecond,
			ok:       true,
		},
		{
			name:     "zero seconds",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			name:     "invalid format",
			input:    "invalid",
			expected: 0,
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration, ok := parseRetryAfter(tt.input)
			assert.Equal(t, tt.ok, ok, "parseRetryAfter(%q) ok = %v, want %v", tt.input, ok, tt.ok)
			if ok {
				assert.Equal(t, tt.expected, duration, "parseRetryAfter(%q) = %v, want %v", tt.input, duration, tt.expected)
			}
		})
	}

	t.Run("http-date future", func(t *testing.T) {
		futureTime := time.Now().Add(30 * time.Second)
		httpDate := futureTime.UTC().Format(http.TimeFormat)
		duration, ok := parseRetryAfter(httpDate)
		assert.True(t, ok, "parseRetryAfter with future HTTP-date should return true")
		assert.InDelta(t, 30*time.Second, duration, float64(1*time.Second), "Duration should be approximately 30 seconds")
	})

	t.Run("http-date past", func(t *testing.T) {
		pastTime := time.Now().Add(-30 * time.Second)
		httpDate := pastTime.UTC().Format(http.TimeFormat)
		duration, ok := parseRetryAfter(httpDate)
		assert.False(t, ok, "parseRetryAfter with past HTTP-date should return false")
		assert.Equal(t, time.Duration(0), duration, "Duration for past date should be 0")
	})
}

func TestIsRetryableStatusCode(t *testing.T) {
	defaultConfig := getDefaultRetryConfig()

	tests := []struct {
		statusCode int
		expected   bool
	}{
		{200, false},
		{400, false},
		{401, false},
		{403, false},
		{404, false},
		{408, true},
		{429, true},
		{500, true},
		{502, true},
		{503, true},
		{504, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("status_%d", tt.statusCode), func(t *testing.T) {
			result := isRetryableStatusCode(tt.statusCode, defaultConfig)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRetryWithContext(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(Error{Error: stringPtr("Server error")})
		assert.NoError(t, err)
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
		RetryConfig: &RetryConfig{
			Enabled:           true,
			MaxAttempts:       5,
			InitialBackoffSec: 2,
			MaxBackoffSec:     10,
			BackoffMultiplier: 2,
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := client.ListModels(ctx)

	assert.Error(t, err)
	assert.GreaterOrEqual(t, callCount, 1)
	assert.LessOrEqual(t, callCount, 5)
}

func TestCustomRetryableStatusCodes(t *testing.T) {
	tests := []struct {
		name              string
		customStatusCodes []int
		statusCode        int
		expected          bool
	}{
		{
			name:              "custom codes include 418",
			customStatusCodes: []int{418, 422},
			statusCode:        418,
			expected:          true,
		},
		{
			name:              "custom codes exclude 500",
			customStatusCodes: []int{418, 422},
			statusCode:        500,
			expected:          false,
		},
		{
			name:              "custom codes include 422",
			customStatusCodes: []int{418, 422},
			statusCode:        422,
			expected:          true,
		},
		{
			name:              "custom codes exclude 200",
			customStatusCodes: []int{418, 422},
			statusCode:        200,
			expected:          false,
		},
		{
			name:              "empty custom codes use defaults",
			customStatusCodes: []int{},
			statusCode:        500,
			expected:          true,
		},
		{
			name:              "nil custom codes use defaults",
			customStatusCodes: nil,
			statusCode:        500,
			expected:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &RetryConfig{
				RetryableStatusCodes: tt.customStatusCodes,
			}
			result := isRetryableStatusCode(tt.statusCode, config)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRetryCallback(t *testing.T) {
	var callbackCalls []struct {
		attempt int
		err     error
		delay   time.Duration
	}

	retryConfig := &RetryConfig{
		Enabled:           true,
		MaxAttempts:       3,
		InitialBackoffSec: 1,
		MaxBackoffSec:     10,
		BackoffMultiplier: 2,
		OnRetry: func(attempt int, err error, delay time.Duration) {
			callbackCalls = append(callbackCalls, struct {
				attempt int
				err     error
				delay   time.Duration
			}{attempt, err, delay})
		},
	}

	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if callCount < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(Error{Error: stringPtr("Server error")})
			assert.NoError(t, err)
		} else {
			w.WriteHeader(http.StatusOK)
			response := ListModelsResponse{Object: "list", Data: []Model{}}
			err := json.NewEncoder(w).Encode(response)
			assert.NoError(t, err)
		}
		callCount++
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL:     baseURL,
		RetryConfig: retryConfig,
	})

	ctx := context.Background()
	_, err := client.ListModels(ctx)

	assert.NoError(t, err)
	assert.Equal(t, 3, callCount)
	assert.Len(t, callbackCalls, 2)

	assert.Equal(t, 1, callbackCalls[0].attempt)
	assert.Contains(t, callbackCalls[0].err.Error(), "HTTP 500")
	assert.Equal(t, 1*time.Second, callbackCalls[0].delay)

	assert.Equal(t, 2, callbackCalls[1].attempt)
	assert.Contains(t, callbackCalls[1].err.Error(), "HTTP 500")
	assert.Equal(t, 2*time.Second, callbackCalls[1].delay)
}

func TestRetryWithCustomStatusCodesAndCallback(t *testing.T) {
	var callbackCalls []struct {
		attempt int
		err     error
		delay   time.Duration
	}

	retryConfig := &RetryConfig{
		Enabled:              true,
		MaxAttempts:          3,
		InitialBackoffSec:    1,
		MaxBackoffSec:        10,
		BackoffMultiplier:    2,
		RetryableStatusCodes: []int{418, 503},
		OnRetry: func(attempt int, err error, delay time.Duration) {
			callbackCalls = append(callbackCalls, struct {
				attempt int
				err     error
				delay   time.Duration
			}{attempt, err, delay})
		},
	}

	tests := []struct {
		name           string
		statusCodes    []int
		expectRetries  bool
		expectedCalls  int
		callbackCounts int
	}{
		{
			name:           "retry on custom 418 status code",
			statusCodes:    []int{418, 418, 200},
			expectRetries:  true,
			expectedCalls:  3,
			callbackCounts: 2,
		},
		{
			name:           "retry on custom 503 status code",
			statusCodes:    []int{503, 200},
			expectRetries:  true,
			expectedCalls:  2,
			callbackCounts: 1,
		},
		{
			name:           "no retry on non-custom 500 status code",
			statusCodes:    []int{500},
			expectRetries:  false,
			expectedCalls:  1,
			callbackCounts: 0,
		},
		{
			name:           "no retry on 400 status code",
			statusCodes:    []int{400},
			expectRetries:  false,
			expectedCalls:  1,
			callbackCounts: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callbackCalls = nil
			callCount := 0

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if callCount < len(tt.statusCodes) {
					w.WriteHeader(tt.statusCodes[callCount])
					if tt.statusCodes[callCount] == 200 {
						response := ListModelsResponse{Object: "list", Data: []Model{}}
						err := json.NewEncoder(w).Encode(response)
						assert.NoError(t, err)
					} else {
						err := json.NewEncoder(w).Encode(Error{Error: stringPtr("Server error")})
						assert.NoError(t, err)
					}
				}
				callCount++
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL:     baseURL,
				RetryConfig: retryConfig,
			})

			ctx := context.Background()
			_, err := client.ListModels(ctx)

			assert.Equal(t, tt.expectedCalls, callCount)
			assert.Len(t, callbackCalls, tt.callbackCounts)

			if tt.expectRetries && tt.statusCodes[len(tt.statusCodes)-1] == 200 {
				assert.NoError(t, err)
			} else if !tt.expectRetries || tt.statusCodes[len(tt.statusCodes)-1] != 200 {
				assert.Error(t, err)
			}
		})
	}
}

func TestRetryAfterHeader(t *testing.T) {
	tests := []struct {
		name             string
		retryAfterValues []string
		expectedDelays   []time.Duration
		tolerance        time.Duration
	}{
		{
			name:             "retry with seconds in header",
			retryAfterValues: []string{"2", "1", ""},
			expectedDelays:   []time.Duration{2 * time.Second, 1 * time.Second},
			tolerance:        100 * time.Millisecond,
		},
		{
			name:             "retry with decimal seconds",
			retryAfterValues: []string{"0.5", ""},
			expectedDelays:   []time.Duration{500 * time.Millisecond},
			tolerance:        100 * time.Millisecond,
		},
		{
			name:             "no retry-after header falls back to exponential backoff",
			retryAfterValues: []string{"", "", ""},
			expectedDelays:   []time.Duration{2 * time.Second, 4 * time.Second},
			tolerance:        100 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			retryDelays := []time.Duration{}
			lastRequestTime := time.Now()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if callCount > 0 {
					retryDelays = append(retryDelays, time.Since(lastRequestTime))
				}
				lastRequestTime = time.Now()

				if callCount < len(tt.retryAfterValues)-1 {
					if tt.retryAfterValues[callCount] != "" {
						w.Header().Set("Retry-After", tt.retryAfterValues[callCount])
					}
					w.WriteHeader(http.StatusTooManyRequests)
					err := json.NewEncoder(w).Encode(Error{Error: stringPtr("Rate limited")})
					assert.NoError(t, err)
				} else {
					w.WriteHeader(http.StatusOK)
					response := ListModelsResponse{Object: "list", Data: []Model{}}
					err := json.NewEncoder(w).Encode(response)
					assert.NoError(t, err)
				}
				callCount++
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
				RetryConfig: &RetryConfig{
					Enabled:           true,
					MaxAttempts:       len(tt.retryAfterValues),
					InitialBackoffSec: 2,
					MaxBackoffSec:     30,
					BackoffMultiplier: 2,
				},
			})

			ctx := context.Background()
			_, err := client.ListModels(ctx)

			assert.NoError(t, err)
			assert.Equal(t, len(tt.retryAfterValues), callCount)
			assert.Len(t, retryDelays, len(tt.expectedDelays))

			for i, expectedDelay := range tt.expectedDelays {
				assert.InDelta(t, expectedDelay, retryDelays[i], float64(tt.tolerance),
					"Retry delay %d should be approximately %v, got %v", i+1, expectedDelay, retryDelays[i])
			}
		})
	}
}

func TestRetryAfterHeaderWithHTTPDate(t *testing.T) {
	callCount := 0
	var actualDelay time.Duration
	lastRequestTime := time.Now()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if callCount > 0 {
			actualDelay = time.Since(lastRequestTime)
		}
		lastRequestTime = time.Now()

		if callCount == 0 {
			futureTime := time.Now().Add(3 * time.Second)
			w.Header().Set("Retry-After", futureTime.UTC().Format(http.TimeFormat))
			w.WriteHeader(http.StatusTooManyRequests)
			err := json.NewEncoder(w).Encode(Error{Error: stringPtr("Rate limited")})
			assert.NoError(t, err)
		} else {
			w.WriteHeader(http.StatusOK)
			response := ListModelsResponse{Object: "list", Data: []Model{}}
			err := json.NewEncoder(w).Encode(response)
			assert.NoError(t, err)
		}
		callCount++
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL: baseURL,
		RetryConfig: &RetryConfig{
			Enabled:           true,
			MaxAttempts:       3,
			InitialBackoffSec: 1,
			MaxBackoffSec:     10,
			BackoffMultiplier: 2,
		},
	})

	ctx := context.Background()
	_, err := client.ListModels(ctx)

	assert.NoError(t, err)
	assert.Equal(t, 2, callCount)
	assert.InDelta(t, 3*time.Second, actualDelay, float64(1*time.Second),
		"Retry delay should be approximately 3 seconds based on HTTP-date header")
}

func TestRetryConfigWithNilCallback(t *testing.T) {
	retryConfig := &RetryConfig{
		Enabled:           true,
		MaxAttempts:       2,
		InitialBackoffSec: 1,
		MaxBackoffSec:     10,
		BackoffMultiplier: 2,
		OnRetry:           nil,
	}

	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if callCount == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(Error{Error: stringPtr("Server error")})
			assert.NoError(t, err)
		} else {
			w.WriteHeader(http.StatusOK)
			response := ListModelsResponse{Object: "list", Data: []Model{}}
			err := json.NewEncoder(w).Encode(response)
			assert.NoError(t, err)
		}
		callCount++
	}))
	defer server.Close()

	baseURL := server.URL + "/v1"
	client := NewClient(&ClientOptions{
		BaseURL:     baseURL,
		RetryConfig: retryConfig,
	})

	ctx := context.Background()
	_, err := client.ListModels(ctx)

	assert.NoError(t, err)
	assert.Equal(t, 2, callCount)
}

func TestListAgents(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   ListAgentsResponse
		expectedAgents int
	}{
		{
			name: "successful agents listing",
			mockResponse: ListAgentsResponse{
				Object: "list",
				Data: []A2AAgentCard{
					{
						Id:                 "agent1",
						Name:               "Test Agent 1",
						Description:        "A test A2A agent for development",
						Version:            "1.0.0",
						Url:                "https://agent1.example.com",
						Capabilities:       map[string]interface{}{"chat": true, "reasoning": true},
						DefaultInputModes:  []string{"text"},
						DefaultOutputModes: []string{"text", "json"},
						Skills:             []map[string]interface{}{{"name": "chat", "type": "conversation"}},
					},
					{
						Id:                 "agent2",
						Name:               "Test Agent 2",
						Description:        "Another test A2A agent",
						Version:            "2.1.0",
						Url:                "https://agent2.example.com",
						Capabilities:       map[string]interface{}{"analysis": true, "reporting": true},
						DefaultInputModes:  []string{"text", "image"},
						DefaultOutputModes: []string{"text", "pdf"},
						Skills:             []map[string]interface{}{{"name": "analysis", "type": "data_processing"}},
					},
				},
			},
			expectedAgents: 2,
		},
		{
			name: "empty agents list",
			mockResponse: ListAgentsResponse{
				Object: "list",
				Data:   []A2AAgentCard{},
			},
			expectedAgents: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/v1/a2a/agents", r.URL.Path, "Path should be /v1/a2a/agents")
				assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")

				w.Header().Set("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(tt.mockResponse)
				assert.NoError(t, err)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			ctx := context.Background()
			agents, err := client.ListAgents(ctx)

			assert.NoError(t, err)
			assert.NotNil(t, agents)
			assert.Equal(t, "list", agents.Object)
			assert.Len(t, agents.Data, tt.expectedAgents)

			if tt.expectedAgents > 0 {
				assert.Equal(t, "agent1", agents.Data[0].Id)
				assert.Equal(t, "Test Agent 1", agents.Data[0].Name)
				assert.Equal(t, "A test A2A agent for development", agents.Data[0].Description)
				assert.Equal(t, "1.0.0", agents.Data[0].Version)
				assert.Equal(t, "https://agent1.example.com", agents.Data[0].Url)
				assert.NotNil(t, agents.Data[0].Capabilities)
				assert.Contains(t, agents.Data[0].DefaultInputModes, "text")
				assert.Contains(t, agents.Data[0].DefaultOutputModes, "text")
			}
		})
	}
}

func TestListAgents_APIError(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   map[string]interface{}
		expectedError  string
	}{
		{
			name:       "A2A not exposed",
			statusCode: http.StatusForbidden,
			responseBody: map[string]interface{}{
				"error": "A2A agents endpoint is not exposed. Set EXPOSE_A2A=true to enable.",
			},
			expectedError: "API error",
		},
		{
			name:       "unauthorized",
			statusCode: http.StatusUnauthorized,
			responseBody: map[string]interface{}{
				"error": "Unauthorized access",
			},
			expectedError: "API error",
		},
		{
			name:       "internal server error",
			statusCode: http.StatusInternalServerError,
			responseBody: map[string]interface{}{
				"error": "Internal server error",
			},
			expectedError: "HTTP 500",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				err := json.NewEncoder(w).Encode(tt.responseBody)
				assert.NoError(t, err)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			ctx := context.Background()
			agents, err := client.ListAgents(ctx)

			assert.Error(t, err)
			assert.Nil(t, agents)
			assert.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestGetAgent(t *testing.T) {
	tests := []struct {
		name         string
		agentID      string
		mockResponse A2AAgentCard
	}{
		{
			name:    "successful agent retrieval",
			agentID: "test-agent-id",
			mockResponse: A2AAgentCard{
				Id:          "test-agent-id",
				Name:        "Detailed Test Agent",
				Description: "A comprehensive test agent with full details",
				Version:     "3.2.1",
				Url:         "https://detailed-agent.example.com",
				IconUrl:     stringPtr("https://detailed-agent.example.com/icon.png"),
				DocumentationUrl: stringPtr("https://detailed-agent.example.com/docs"),
				Capabilities: map[string]interface{}{
					"chat":      true,
					"reasoning": true,
					"analysis":  true,
					"vision":    true,
				},
				DefaultInputModes:  []string{"text", "image", "audio"},
				DefaultOutputModes: []string{"text", "json", "image"},
				Skills: []map[string]interface{}{
					{"name": "conversation", "type": "chat", "enabled": true},
					{"name": "document_analysis", "type": "analysis", "enabled": true},
					{"name": "image_processing", "type": "vision", "enabled": true},
				},
				Provider: &map[string]interface{}{
					"name":    "Test Provider",
					"version": "1.0",
					"url":     "https://provider.example.com",
				},
				Security: &[]map[string]interface{}{
					{"type": "bearer", "scheme": "JWT"},
				},
				SecuritySchemes: &map[string]interface{}{
					"bearerAuth": map[string]interface{}{
						"type":         "http",
						"scheme":       "bearer",
						"bearerFormat": "JWT",
					},
				},
				SupportsAuthenticatedExtendedCard: boolPtr(true),
			},
		},
		{
			name:    "minimal agent data",
			agentID: "minimal-agent",
			mockResponse: A2AAgentCard{
				Id:                 "minimal-agent",
				Name:               "Minimal Agent",
				Description:        "Basic agent with minimal configuration",
				Version:            "1.0.0",
				Url:                "https://minimal.example.com",
				Capabilities:       map[string]interface{}{"basic": true},
				DefaultInputModes:  []string{"text"},
				DefaultOutputModes: []string{"text"},
				Skills:             []map[string]interface{}{{"name": "basic", "type": "simple"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := fmt.Sprintf("/v1/a2a/agents/%s", tt.agentID)
				assert.Equal(t, expectedPath, r.URL.Path, "Path should match agent ID")
				assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")

				w.Header().Set("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(tt.mockResponse)
				assert.NoError(t, err)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			ctx := context.Background()
			agent, err := client.GetAgent(ctx, tt.agentID)

			assert.NoError(t, err)
			assert.NotNil(t, agent)
			assert.Equal(t, tt.mockResponse.Id, agent.Id)
			assert.Equal(t, tt.mockResponse.Name, agent.Name)
			assert.Equal(t, tt.mockResponse.Description, agent.Description)
			assert.Equal(t, tt.mockResponse.Version, agent.Version)
			assert.Equal(t, tt.mockResponse.Url, agent.Url)
			assert.Equal(t, tt.mockResponse.Capabilities, agent.Capabilities)
			assert.Equal(t, tt.mockResponse.DefaultInputModes, agent.DefaultInputModes)
			assert.Equal(t, tt.mockResponse.DefaultOutputModes, agent.DefaultOutputModes)

			if tt.name == "successful agent retrieval" {
				assert.Equal(t, tt.mockResponse.IconUrl, agent.IconUrl)
				assert.Equal(t, tt.mockResponse.DocumentationUrl, agent.DocumentationUrl)
				assert.Equal(t, tt.mockResponse.Provider, agent.Provider)
				assert.Equal(t, tt.mockResponse.Security, agent.Security)
				assert.Equal(t, tt.mockResponse.SecuritySchemes, agent.SecuritySchemes)
				assert.Equal(t, tt.mockResponse.SupportsAuthenticatedExtendedCard, agent.SupportsAuthenticatedExtendedCard)
			}
		})
	}
}

func TestGetAgent_APIError(t *testing.T) {
	tests := []struct {
		name           string
		agentID        string
		statusCode     int
		responseBody   map[string]interface{}
		expectedError  string
	}{
		{
			name:       "agent not found",
			agentID:    "nonexistent-agent",
			statusCode: http.StatusNotFound,
			responseBody: map[string]interface{}{
				"error": "Agent not found",
			},
			expectedError: "API error",
		},
		{
			name:       "A2A not exposed",
			agentID:    "test-agent",
			statusCode: http.StatusForbidden,
			responseBody: map[string]interface{}{
				"error": "A2A agents endpoint is not exposed. Set EXPOSE_A2A=true to enable.",
			},
			expectedError: "API error",
		},
		{
			name:       "unauthorized",
			agentID:    "test-agent",
			statusCode: http.StatusUnauthorized,
			responseBody: map[string]interface{}{
				"error": "Unauthorized access",
			},
			expectedError: "API error",
		},
		{
			name:       "internal server error",
			agentID:    "test-agent",
			statusCode: http.StatusInternalServerError,
			responseBody: map[string]interface{}{
				"error": "Internal server error",
			},
			expectedError: "HTTP 500",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				err := json.NewEncoder(w).Encode(tt.responseBody)
				assert.NoError(t, err)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			ctx := context.Background()
			agent, err := client.GetAgent(ctx, tt.agentID)

			assert.Error(t, err)
			assert.Nil(t, agent)
			assert.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestGetAgent_InvalidID(t *testing.T) {
	tests := []struct {
		name     string
		agentID  string
		expected string
	}{
		{
			name:     "empty agent ID",
			agentID:  "",
			expected: "a2a/agents/",
		},
		{
			name:     "agent ID with special characters",
			agentID:  "agent@123test",
			expected: "a2a/agents/agent@123test",
		},
		{
			name:     "agent ID with spaces",
			agentID:  "agent with spaces",
			expected: "a2a/agents/agent with spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Contains(t, r.URL.Path, tt.expected, "Path should contain expected agent ID")
				assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")

				agent := A2AAgentCard{
					Id:                 tt.agentID,
					Name:               "Test Agent",
					Description:        "Test agent description",
					Version:            "1.0.0",
					Url:                "https://test.example.com",
					Capabilities:       map[string]interface{}{"test": true},
					DefaultInputModes:  []string{"text"},
					DefaultOutputModes: []string{"text"},
					Skills:             []map[string]interface{}{{"name": "test", "type": "basic"}},
				}

				w.Header().Set("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(agent)
				assert.NoError(t, err)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
			})

			ctx := context.Background()
			agent, err := client.GetAgent(ctx, tt.agentID)

			assert.NoError(t, err)
			assert.NotNil(t, agent)
			assert.Equal(t, tt.agentID, agent.Id)
		})
	}
}

func TestA2AWithTimeout(t *testing.T) {
	tests := []struct {
		name     string
		function string
	}{
		{
			name:     "ListAgents with timeout",
			function: "ListAgents",
		},
		{
			name:     "GetAgent with timeout",
			function: "GetAgent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(200 * time.Millisecond)
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			baseURL := server.URL + "/v1"
			client := NewClient(&ClientOptions{
				BaseURL: baseURL,
				Timeout: 100 * time.Millisecond,
			})

			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
			defer cancel()

			var err error
			if tt.function == "ListAgents" {
				_, err = client.ListAgents(ctx)
			} else {
				_, err = client.GetAgent(ctx, "test-agent")
			}

			assert.Error(t, err)
			assert.Contains(t, err.Error(), "context deadline exceeded")
		})
	}
}

func TestOllamaCloudProvider(t *testing.T) {
	t.Run("ListProviderModels for Ollama Cloud", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/models", r.URL.Path, "Path should be /v1/models")
			assert.Equal(t, http.MethodGet, r.Method, "Method should be GET")
			assert.Equal(t, "ollama_cloud", r.URL.Query().Get("provider"), "Provider should be ollama_cloud")

			response := ListModelsResponse{
				Provider: providerPtr(OllamaCloud),
				Object:   "list",
				Data: []Model{
					{
						Id:       "ollama_cloud/gpt-oss:20b",
						Object:   "model",
						Created:  1730419200,
						OwnedBy:  "ollama_cloud",
						ServedBy: OllamaCloud,
					},
					{
						Id:       "ollama_cloud/llama3.3:70b",
						Object:   "model",
						Created:  1730419200,
						OwnedBy:  "ollama_cloud",
						ServedBy: OllamaCloud,
					},
				},
			}

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(response)
			assert.NoError(t, err)
		}))
		defer server.Close()

		baseURL := server.URL + "/v1"
		client := NewClient(&ClientOptions{
			BaseURL: baseURL,
		})

		ctx := context.Background()
		models, err := client.ListProviderModels(ctx, OllamaCloud)

		assert.NoError(t, err)
		assert.NotNil(t, models)
		assert.Equal(t, OllamaCloud, *models.Provider)
		assert.Len(t, models.Data, 2)
		assert.Equal(t, "ollama_cloud/gpt-oss:20b", models.Data[0].Id)
		assert.Equal(t, "ollama_cloud/llama3.3:70b", models.Data[1].Id)
	})

	t.Run("GenerateContent with Ollama Cloud", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/chat/completions", r.URL.Path, "Path should be /v1/chat/completions")
			assert.Equal(t, http.MethodPost, r.Method, "Method should be POST")
			assert.Equal(t, "ollama_cloud", r.URL.Query().Get("provider"), "Provider should be ollama_cloud")

			var req CreateChatCompletionRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			assert.NoError(t, err)
			assert.Equal(t, "ollama_cloud/gpt-oss:20b", req.Model)

			response := CreateChatCompletionResponse{
				Id:      "chatcmpl-test-ollama-cloud",
				Object:  "chat.completion",
				Created: 1730419200,
				Model:   "ollama_cloud/gpt-oss:20b",
				Choices: []ChatCompletionChoice{
					{
						Index: 0,
						Message: Message{
							Role:    Assistant,
							Content: "Hello! I'm an AI assistant powered by Ollama Cloud. How can I help you today?",
						},
						FinishReason: Stop,
					},
				},
				Usage: &CompletionUsage{
					PromptTokens:     10,
					CompletionTokens: 20,
					TotalTokens:      30,
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

		ctx := context.Background()
		messages := []Message{
			{Role: User, Content: "Hello"},
		}

		response, err := client.GenerateContent(ctx, OllamaCloud, "ollama_cloud/gpt-oss:20b", messages)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "chatcmpl-test-ollama-cloud", response.Id)
		assert.Equal(t, "ollama_cloud/gpt-oss:20b", response.Model)
		assert.Len(t, response.Choices, 1)
		assert.Equal(t, "Hello! I'm an AI assistant powered by Ollama Cloud. How can I help you today?", response.Choices[0].Message.Content)
	})
}

