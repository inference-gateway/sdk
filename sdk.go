package sdk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Provider represents supported LLM providers
type Provider string

const (
	ProviderOllama     Provider = "ollama"     // Ollama LLM provider
	ProviderGroq       Provider = "groq"       // Groq LLM provider
	ProviderOpenAI     Provider = "openai"     // OpenAI LLM provider
	ProviderGoogle     Provider = "google"     // Google LLM provider
	ProviderCloudflare Provider = "cloudflare" // Cloudflare LLM provider
	ProviderCohere     Provider = "cohere"     // Cohere LLM provider
)

// Model represents an LLM model
type Model struct {
	ID      string `json:"id"`       // Unique identifier for the model
	Object  string `json:"object"`   // Type of object (always "model")
	OwnedBy string `json:"owned_by"` // Organization that owns the model
	Created int64  `json:"created"`  // Unix timestamp of when the model was created
}

// ProviderModels represents models available for a provider
type ProviderModels struct {
	Provider Provider `json:"provider"` // The LLM provider (e.g., "ollama")
	Models   []Model  `json:"models"`   // List of available models for the provider
}

// Role represents supported message roles
type Role string

const (
	RoleSystem    Role = "system"    // System message role
	RoleUser      Role = "user"      // User message role
	RoleAssistant Role = "assistant" // Assistant message role
)

// Message represents a chat message
type Message struct {
	Role    Role   `json:"role"`    // Role of the message sender
	Content string `json:"content"` // Content of the message
}

// GenerateRequest represents the request for content generation
type GenerateRequest struct {
	Model    string    `json:"model"`    // Name of the model to use
	Messages []Message `json:"messages"` // List of messages in the conversation
}

// GenerateResponseTokens represents the response tokens from content generation
type GenerateResponseTokens struct {
	Role    Role   `json:"role"`    // Role of the response (usually "assistant")
	Model   string `json:"model"`   // Model that generated the response
	Content string `json:"content"` // Generated content
}

// GenerateResponse represents the response from content generation
type GenerateResponse struct {
	Provider Provider               `json:"provider"` // Provider that generated the response
	Response GenerateResponseTokens `json:"response"` // The generated response
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error string `json:"error"` // Error message from the API
}

// Client represents the SDK client interface
type Client interface {
	ListModels() ([]ProviderModels, error)
	GenerateContent(provider Provider, model string, messages []Message) (*GenerateResponse, error)
	HealthCheck() error
}

// clientImpl represents the concrete implementation of the SDK client
type clientImpl struct {
	baseURL string        // Base URL of the Inference Gateway API
	http    *resty.Client // HTTP client for making requests
}

// NewClient creates a new SDK client with the specified base URL.
//
// Example:
//
//	client := sdk.NewClient("http://localhost:8080")
func NewClient(baseURL string) Client {
	return &clientImpl{
		baseURL: baseURL,
		http:    resty.New(),
	}
}

// ListModels returns all available language models from all providers.
//
// Example:
//
//	models, err := client.ListModels()
//	if err != nil {
//	    log.Fatalf("Error listing models: %v", err)
//	}
//	fmt.Printf("Available models: %+v\n", models)
func (c *clientImpl) ListModels() ([]ProviderModels, error) {
	var models []ProviderModels
	resp, err := c.http.R().
		SetResult(&models).
		Get(fmt.Sprintf("%s/llms", c.baseURL))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errResp ErrorResponse
		if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode())
		}
		return nil, fmt.Errorf("API error: %s", errResp.Error)
	}

	return models, nil
}

// GenerateContent generates content using the specified provider and model.
//
// Example:
//
//	response, err := client.GenerateContent(
//	    sdk.ProviderOllama,
//	    "llama2",
//	    []sdk.Message{
//	        {
//	            Role:    sdk.RoleSystem,
//	            Content: "You are a helpful assistant.",
//	        },
//	        {
//	            Role:    sdk.RoleUser,
//	            Content: "What is Go?",
//	        },
//	    },
//	)
//	if err != nil {
//	    log.Fatalf("Error generating content: %v", err)
//	}
//	fmt.Printf("Generated content: %s\n", response.Response.Content)
func (c *clientImpl) GenerateContent(provider Provider, model string, messages []Message) (*GenerateResponse, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("at least one message is required")
	}

	req := GenerateRequest{
		Model:    model,
		Messages: messages,
	}

	var result GenerateResponse
	resp, err := c.http.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s/llms/%s/generate", c.baseURL, provider))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errResp ErrorResponse
		if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode())
		}
		return nil, fmt.Errorf("API error: %s", errResp.Error)
	}

	return &result, nil
}

// HealthCheck performs a health check request to verify API availability.
//
// Example:
//
//	err := client.HealthCheck()
//	if err != nil {
//	    log.Fatalf("Health check failed: %v", err)
//	}
func (c *clientImpl) HealthCheck() error {
	resp, err := c.http.R().
		Get(fmt.Sprintf("%s/health", c.baseURL))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("health check failed with status: %d", resp.StatusCode())
	}

	return nil
}
