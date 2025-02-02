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
)

func TestListModels(t *testing.T) {
	tests := []struct {
		name          string
		serverHandler func(w http.ResponseWriter, r *http.Request)
		expectedError string
		expectedResp  []ListModelsResponse
	}{
		{
			name: "successful list models",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/llms", r.URL.Path)

				w.Header().Set("Content-Type", "application/json")
				models := []ListModelsResponse{
					{
						Provider: ProviderOllama,
						Models: []Model{
							{
								Name: "llama2",
							},
						},
					},
				}
				err := json.NewEncoder(w).Encode(models)
				assert.NoError(t, err)
			},
			expectedResp: []ListModelsResponse{
				{
					Provider: ProviderOllama,
					Models: []Model{
						{
							Name: "llama2",
						},
					},
				},
			},
		},
		{
			name: "server error",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode(ErrorResponse{Error: "internal error"})
				assert.NoError(t, err)
			},
			expectedError: "API error: internal error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverHandler))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()
			resp, err := client.ListModels(ctx)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestListProviderModels(t *testing.T) {
	tests := []struct {
		name          string
		provider      Provider
		serverHandler func(w http.ResponseWriter, r *http.Request)
		expectedError string
		expectedResp  ListModelsResponse
	}{
		{
			name:     "successful list provider models",
			provider: ProviderGroq,
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/llms/groq", r.URL.Path)

				w.Header().Set("Content-Type", "application/json")
				response := ListModelsResponse{
					Provider: ProviderGroq,
					Models: []Model{
						{Name: "llama-3.3-70b-versatile"},
						{Name: "llama-3.2-3b-preview"},
						{Name: "llama-3.2-1b-preview"},
						{Name: "llama-3.3-70b-specdec"},
						{Name: "llama3-8b-8192"},
					},
				}
				err := json.NewEncoder(w).Encode(response)
				assert.NoError(t, err)
			},
			expectedResp: ListModelsResponse{
				Provider: ProviderGroq,
				Models: []Model{
					{Name: "llama-3.3-70b-versatile"},
					{Name: "llama-3.2-3b-preview"},
					{Name: "llama-3.2-1b-preview"},
					{Name: "llama-3.3-70b-specdec"},
					{Name: "llama3-8b-8192"},
				},
			},
		},
		{
			name:     "server error",
			provider: ProviderGroq,
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode(ErrorResponse{Error: "internal error"})
				assert.NoError(t, err)
			},
			expectedError: "API error: internal error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverHandler))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()
			resp, err := client.ListProviderModels(ctx, tt.provider)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp.Models, resp)
			}
		})
	}
}

func TestGenerateContent(t *testing.T) {
	tests := []struct {
		name          string
		provider      Provider
		model         string
		messages      []Message
		serverHandler func(w http.ResponseWriter, r *http.Request)
		expectedError string
		expectedResp  *GenerateResponse
	}{
		{
			name:     "successful generation",
			provider: ProviderOllama,
			model:    "llama2",
			messages: []Message{
				{
					Role:    MessageRoleSystem,
					Content: "You are a helpful assistant.",
				},
				{
					Role:    MessageRoleUser,
					Content: "What is Go?",
				},
			},
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/llms/ollama/generate", r.URL.Path)

				var req GenerateRequest
				err := json.NewDecoder(r.Body).Decode(&req)
				assert.NoError(t, err)
				assert.Equal(t, "llama2", req.Model)

				assert.Equal(t, 2, len(req.Messages))
				assert.Equal(t, MessageRoleSystem, req.Messages[0].Role)
				assert.Equal(t, "You are a helpful assistant.", req.Messages[0].Content)
				assert.Equal(t, MessageRoleUser, req.Messages[1].Role)
				assert.Equal(t, "What is Go?", req.Messages[1].Content)

				w.Header().Set("Content-Type", "application/json")
				resp := &GenerateResponse{
					Provider: ProviderOllama,
					Response: GenerateResponseTokens{
						Role:    MessageRoleAssistant,
						Model:   "llama2",
						Content: "Go is a programming language.",
					},
				}
				err = json.NewEncoder(w).Encode(resp)
				assert.NoError(t, err)
			},
			expectedResp: &GenerateResponse{
				Provider: ProviderOllama,
				Response: GenerateResponseTokens{
					Role:    MessageRoleAssistant,
					Model:   "llama2",
					Content: "Go is a programming language.",
				},
			},
		},
		{
			name:     "server error",
			provider: ProviderOllama,
			model:    "llama2",
			messages: []Message{
				{
					Role:    MessageRoleSystem,
					Content: "You are a helpful assistant.",
				},
				{
					Role:    MessageRoleUser,
					Content: "What is Go?",
				},
			},
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode(ErrorResponse{Error: "model not found"})
				assert.NoError(t, err)
			},
			expectedError: "API error: model not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverHandler))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()
			resp, err := client.GenerateContent(ctx, tt.provider, tt.model, tt.messages)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestHealthCheck(t *testing.T) {
	tests := []struct {
		name          string
		serverHandler func(w http.ResponseWriter, r *http.Request)
		expectedError string
	}{
		{
			name: "healthy server",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/health", r.URL.Path)
				w.WriteHeader(http.StatusOK)
			},
		},
		{
			name: "unhealthy server",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusServiceUnavailable)
			},
			expectedError: "health check failed with status: 503",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverHandler))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()
			err := client.HealthCheck(ctx)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGenerateContentStream(t *testing.T) {
	tests := []struct {
		name          string
		provider      Provider
		model         string
		messages      []Message
		serverHandler func(w http.ResponseWriter, r *http.Request)
		expectedError string
		validate      func(*testing.T, <-chan SSEvent)
	}{
		{
			name:     "successful stream",
			provider: ProviderOllama,
			model:    "llama2",
			messages: []Message{
				{Role: MessageRoleSystem, Content: "You are helpful."},
				{Role: MessageRoleUser, Content: "Hi"},
			},
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/llms/ollama/generate", r.URL.Path)

				var req GenerateRequest
				err := json.NewDecoder(r.Body).Decode(&req)
				assert.NoError(t, err)
				assert.True(t, req.Stream)
				assert.True(t, req.SSEvents)

				w.Header().Set("Content-Type", "text/event-stream")
				w.WriteHeader(http.StatusOK)

				events := []SSEvent{
					{Event: StreamEventMessageStart, Data: json.RawMessage(`{"role":"assistant"}`)},
					{Event: StreamEventStreamStart, Data: json.RawMessage(`{}`)},
					{Event: StreamEventContentStart, Data: json.RawMessage(`{}`)},
					{Event: StreamEventContentDelta, Data: json.RawMessage(`{"content":"Hello"}`)},
					{Event: StreamEventContentDelta, Data: json.RawMessage(`{"content":" there"}`)},
					{Event: StreamEventContentDelta, Data: json.RawMessage(`{"content":"!"}`)},
					{Event: StreamEventContentEnd, Data: json.RawMessage(`{}`)},
					{Event: StreamEventMessageEnd, Data: json.RawMessage(`{}`)},
					{Event: StreamEventStreamEnd, Data: json.RawMessage(`{}`)},
				}

				for _, evt := range events {
					data, err := evt.Data.MarshalJSON()
					assert.NoError(t, err)
					_, err = fmt.Fprintf(w, "event: %s\ndata: %s\n\n", evt.Event, data)
					assert.NoError(t, err)
					w.(http.Flusher).Flush()
				}
			},
			validate: func(t *testing.T, events <-chan SSEvent) {
				expected := []SSEvent{
					{Event: StreamEventMessageStart, Data: []byte(`{"role":"assistant"}`)},
					{Event: StreamEventStreamStart, Data: []byte(`{}`)},
					{Event: StreamEventContentStart, Data: []byte(`{}`)},
					{Event: StreamEventContentDelta, Data: []byte(`{"content":"Hello"}`)},
					{Event: StreamEventContentDelta, Data: []byte(`{"content":" there"}`)},
					{Event: StreamEventContentDelta, Data: []byte(`{"content":"!"}`)},
					{Event: StreamEventContentEnd, Data: []byte(`{}`)},
					{Event: StreamEventMessageEnd, Data: []byte(`{}`)},
					{Event: StreamEventStreamEnd, Data: []byte(`{}`)},
				}

				for _, expectedEvent := range expected {
					event := <-events
					assert.Equal(t, expectedEvent.Event, event.Event)
					assert.JSONEq(t, string(expectedEvent.Data), string(event.Data))
				}

				_, more := <-events
				assert.False(t, more, "channel should be closed")
			},
		},
		{
			name:     "server error",
			provider: ProviderOllama,
			model:    "llama2",
			messages: []Message{
				{Role: MessageRoleUser, Content: "Hi"},
			},
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/event-stream")
				w.WriteHeader(http.StatusOK)

				events := []SSEvent{
					{Event: StreamEventMessageStart, Data: json.RawMessage(`{"role":"assistant"}`)},
					{Event: StreamEventMessageError, Data: json.RawMessage(`{"error":"error event captured"}`)},
					{Event: StreamEventStreamEnd, Data: json.RawMessage(`{}`)},
				}

				for _, evt := range events {
					data, err := evt.Data.MarshalJSON()
					assert.NoError(t, err)
					_, err = fmt.Fprintf(w, "event: %s\ndata: %s\n\n", evt.Event, data)
					assert.NoError(t, err)
					w.(http.Flusher).Flush()
				}
			},
			validate: func(t *testing.T, events <-chan SSEvent) {
				expected := []SSEvent{
					{Event: StreamEventMessageStart, Data: json.RawMessage(`{"role":"assistant"}`)},
					{Event: StreamEventMessageError, Data: json.RawMessage(`{"error":"error event captured"}`)},
					{Event: StreamEventStreamEnd, Data: json.RawMessage(`{}`)},
				}

				for _, expectedEvent := range expected {
					event := <-events
					t.Logf("Received event: %+v with data: %s", event.Event, string(event.Data))
					assert.Equal(t, expectedEvent.Event, event.Event)
					assert.JSONEq(t, string(expectedEvent.Data), string(event.Data))
				}

				_, more := <-events
				assert.False(t, more, "channel should be closed")
			},
		},
		{
			name:     "context canceled",
			provider: ProviderOllama,
			model:    "llama2",
			messages: []Message{
				{Role: MessageRoleUser, Content: "Hi"},
			},
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/event-stream")
				w.WriteHeader(http.StatusOK)

				event := SSEvent{
					Event: StreamEventMessageStart,
					Data:  json.RawMessage(`{"role":"assistant"}`),
				}

				_, err := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event.Event, event.Data)
				assert.NoError(t, err)
				w.(http.Flusher).Flush()

				time.Sleep(100 * time.Millisecond)
			},
			validate: func(t *testing.T, events <-chan SSEvent) {
				ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
				defer cancel()

				expected := SSEvent{
					Event: StreamEventMessageStart,
					Data:  json.RawMessage(`{"role":"assistant"}`),
				}

				select {
				case event := <-events:
					assert.Equal(t, expected.Event, event.Event)
					assert.JSONEq(t, string(expected.Data), string(event.Data))
				case <-ctx.Done():
					t.Fatal("timeout waiting for first event")
				}

				<-ctx.Done()
				_, more := <-events
				assert.False(t, more, "channel should be closed after context cancellation")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverHandler))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()
			events, err := client.GenerateContentStream(ctx, tt.provider, tt.model, tt.messages)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, events)
			} else {
				assert.NoError(t, err)
				tt.validate(t, events)
			}
		})
	}
}
