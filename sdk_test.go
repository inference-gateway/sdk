package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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

		// Send stream responses as single lines
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

func stringPtr(s string) *string {
	return &s
}

func providerPtr(p Provider) *Provider {
	return &p
}
