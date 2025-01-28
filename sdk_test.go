package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
			resp, err := client.ListModels()

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
		expectedResp  ListModelsResponse // Changed from *ListModelsResponse
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
			resp, err := client.ListProviderModels(tt.provider)

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
					Role:    RoleSystem,
					Content: "You are a helpful assistant.",
				},
				{
					Role:    RoleUser,
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
				assert.Equal(t, RoleSystem, req.Messages[0].Role)
				assert.Equal(t, "You are a helpful assistant.", req.Messages[0].Content)
				assert.Equal(t, RoleUser, req.Messages[1].Role)
				assert.Equal(t, "What is Go?", req.Messages[1].Content)

				w.Header().Set("Content-Type", "application/json")
				resp := &GenerateResponse{
					Provider: ProviderOllama,
					Response: GenerateResponseTokens{
						Role:    RoleAssistant,
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
					Role:    RoleAssistant,
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
					Role:    RoleSystem,
					Content: "You are a helpful assistant.",
				},
				{
					Role:    RoleUser,
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
			resp, err := client.GenerateContent(tt.provider, tt.model, tt.messages)

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
			err := client.HealthCheck()

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
