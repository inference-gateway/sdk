package sdk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Provider represents supported LLM providers
type Provider string

const (
	ProviderOllama     Provider = "ollama"
	ProviderGroq       Provider = "groq"
	ProviderOpenAI     Provider = "openai"
	ProviderGoogle     Provider = "google"
	ProviderCloudflare Provider = "cloudflare"
	ProviderCohere     Provider = "cohere"
)

// Model represents an LLM model
type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
	Created int64  `json:"created"`
}

// ProviderModels represents models available for a provider
type ProviderModels struct {
	Provider Provider `json:"provider"`
	Models   []Model  `json:"models"`
}

// Role represents supported message roles
type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

// Message represents a chat message
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// GenerateRequest represents the request for content generation
type GenerateRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// GenerateResponseTokens represents the response tokens from content generation
type GenerateResponseTokens struct {
	Role    Role   `json:"role"`
	Model   string `json:"model"`
	Content string `json:"content"`
}

// GenerateResponse represents the response from content generation
type GenerateResponse struct {
	Provider Provider               `json:"provider"`
	Response GenerateResponseTokens `json:"response"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error string `json:"error"`
}

// Client represents the SDK client
type Client struct {
	baseURL string
	http    *resty.Client
}

// NewClient creates a new SDK client
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		http:    resty.New(),
	}
}

// ListModels returns all available language models
func (c *Client) ListModels() ([]ProviderModels, error) {
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

// GenerateContent generates content using the specified provider and model
func (c *Client) GenerateContent(provider Provider, model string, messages []Message) (*GenerateResponse, error) {
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

// HealthCheck performs a health check request
func (c *Client) HealthCheck() error {
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
