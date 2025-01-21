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
		expectedResp  []ProviderModels
	}{
		{
			name: "successful list models",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/llms", r.URL.Path)

				w.Header().Set("Content-Type", "application/json")
				models := []ProviderModels{
					{
						Provider: ProviderOllama,
						Models: []Model{
							{
								ID:      "llama2",
								Object:  "model",
								OwnedBy: "ollama",
								Created: 1631318400,
							},
						},
					},
				}
				err := json.NewEncoder(w).Encode(models)
				assert.NoError(t, err)
			},
			expectedResp: []ProviderModels{
				{
					Provider: ProviderOllama,
					Models: []Model{
						{
							ID:      "llama2",
							Object:  "model",
							OwnedBy: "ollama",
							Created: 1631318400,
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
				json.NewEncoder(w).Encode(ErrorResponse{Error: "internal error"})
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

func TestGenerateContent(t *testing.T) {
	tests := []struct {
		name          string
		provider      Provider
		model         string
		prompt        string
		serverHandler func(w http.ResponseWriter, r *http.Request)
		expectedError string
		expectedResp  *GenerateResponse
	}{
		{
			name:     "successful generation",
			provider: ProviderOllama,
			model:    "llama2",
			prompt:   "What is Go?",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/llms/ollama/generate", r.URL.Path)

				var req GenerateRequest
				err := json.NewDecoder(r.Body).Decode(&req)
				assert.NoError(t, err)
				assert.Equal(t, "llama2", req.Model)
				assert.Equal(t, "What is Go?", req.Messages[0].Content)

				w.Header().Set("Content-Type", "application/json")
				resp := &GenerateResponse{
					Provider: ProviderOllama,
					Response: struct {
						Role    string `json:"role"`
						Model   string `json:"model"`
						Content string `json:"content"`
					}{
						Role:    "assistant",
						Model:   "llama2",
						Content: "Go is a programming language.",
					},
				}
				err = json.NewEncoder(w).Encode(resp)
				assert.NoError(t, err)
			},
			expectedResp: &GenerateResponse{
				Provider: ProviderOllama,
				Response: struct {
					Role    string `json:"role"`
					Model   string `json:"model"`
					Content string `json:"content"`
				}{
					Role:    "assistant",
					Model:   "llama2",
					Content: "Go is a programming language.",
				},
			},
		},
		{
			name:     "server error",
			provider: ProviderOllama,
			model:    "llama2",
			prompt:   "What is Go?",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{Error: "model not found"})
			},
			expectedError: "API error: model not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverHandler))
			defer server.Close()

			client := NewClient(server.URL)
			resp, err := client.GenerateContent(tt.provider, tt.model, tt.prompt)

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
