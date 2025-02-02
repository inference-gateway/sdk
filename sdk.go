package sdk

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

// Provider represents supported LLM providers
type Provider string

const (
	ProviderOllama     Provider = "ollama"     // Ollama LLM provider
	ProviderGroq       Provider = "groq"       // Groq LLM provider
	ProviderOpenAI     Provider = "openai"     // OpenAI LLM provider
	ProviderCloudflare Provider = "cloudflare" // Cloudflare LLM provider
	ProviderCohere     Provider = "cohere"     // Cohere LLM provider
	ProviderAnthropic  Provider = "anthropic"  // Anthropic LLM provider
)

// Model represents an LLM model
type Model struct {
	Name string `json:"name"` // Unique identifier for the model
}

type ListModelsResponse struct {
	Provider Provider `json:"provider"` // The LLM provider (e.g., "Ollama")
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
	Stream   bool      `json:"stream"`   // Enable streaming mode
	SSEvents bool      `json:"ssevents"` // Enable SSE events
}

// SSEvent represents a Server-Sent Event from the content generation stream
type SSEvent struct {
	Event StreamEvent     `json:"event,omitempty"` // The type of the SSE event
	Data  json.RawMessage `json:"data,omitempty"`  // The content payload of the event
}

// ResponseError represents an error response from the API
type ResponseError struct {
	Error string `json:"error"`
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
	ListModels() ([]ListModelsResponse, error)
	ListProviderModels(provider Provider) ([]Model, error)
	GenerateContent(provider Provider, model string, messages []Message) (*GenerateResponse, error)
	GenerateContentStream(ctx context.Context, provider Provider, model string, messages []Message) (<-chan SSEvent, error)
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
func (c *clientImpl) ListModels() ([]ListModelsResponse, error) {
	var models []ListModelsResponse
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

// ListProviderModels returns all available language models for a specific provider.
//
// Example:
//
//	models, err := client.ListProviderModels(sdk.ProviderOllama)
//	if err != nil {
//	    log.Fatalf("Error listing models: %v", err)
//	}
//	fmt.Printf("Available models: %+v\n", models)
func (c *clientImpl) ListProviderModels(provider Provider) ([]Model, error) {
	var response ListModelsResponse
	resp, err := c.http.R().
		SetResult(&response).
		Get(fmt.Sprintf("%s/llms/%s", c.baseURL, provider))

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

	return response.Models, nil
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

type StreamEvent string

const (
	// StreamEventMessageError represents an error message
	StreamEventMessageError StreamEvent = "error"
	// StreamEventMessageStart represents the start of a new message
	StreamEventMessageStart StreamEvent = "message-start"
	// StreamEventStreamStart represents the start of the stream
	StreamEventStreamStart StreamEvent = "stream-start"
	// StreamEventContentStart represents the start of the content
	StreamEventContentStart StreamEvent = "content-start"
	// StreamEventContentDelta represents a content delta
	StreamEventContentDelta StreamEvent = "content-delta"
	// StreamEventContentEnd represents the end of the content
	StreamEventContentEnd StreamEvent = "content-end"
	// StreamEventMessageEnd represents the end of a message
	StreamEventMessageEnd StreamEvent = "message-end"
	// StreamEventStreamEnd represents the end of the stream
	StreamEventStreamEnd StreamEvent = "stream-end"
)

// ParseSSEvents parses a Server-Sent Event from a byte slice
func ParseSSEvents(line []byte) (*SSEvent, error) {
	if len(bytes.TrimSpace(line)) == 0 {
		return nil, fmt.Errorf("empty line")
	}

	lines := bytes.Split(line, []byte("\n"))
	event := &SSEvent{}
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) != 2 {
			continue
		}

		field := string(bytes.TrimSpace(parts[0]))
		value := bytes.TrimSpace(parts[1])

		switch field {
		case "data":
			event.Data = value
		case "event":
			event.Event = StreamEvent(string(value))
		}
	}

	return event, nil
}

// GenerateContentStream generates content using streaming mode and returns a channel of events
func (c *clientImpl) GenerateContentStream(ctx context.Context, provider Provider, model string, messages []Message) (<-chan SSEvent, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("at least one message is required")
	}

	req := GenerateRequest{
		Model:    model,
		Messages: messages,
		Stream:   true,
		SSEvents: true,
	}

	resp, err := c.http.R().
		SetDoNotParseResponse(true).
		SetBody(req).
		Post(fmt.Sprintf("%s/llms/%s/generate", c.baseURL, provider))

	if err != nil {
		return nil, err
	}

	eventChan := make(chan SSEvent)
	reader := bufio.NewReader(resp.RawBody())

	go func() {
		defer close(eventChan)
		defer resp.RawBody().Close()

		for {
			select {
			case <-ctx.Done():
				eventChan <- SSEvent{
					Event: "error",
					Data:  []byte("context canceled"),
				}
				return
			default:
				chunk, err := readSSEventsChunk(reader)
				if err != nil {
					if err != io.EOF {
						eventChan <- SSEvent{
							Event: "error",
							Data:  []byte("error reading stream chunk"),
						}
					}
					return
				}

				event, err := ParseSSEvents(chunk)
				if err != nil {
					eventChan <- SSEvent{
						Event: "error",
						Data:  []byte("error parsing stream event"),
					}
					return
				}

				eventChan <- *event

				if event.Event == StreamEventStreamEnd {
					return
				}
			}
		}
	}()

	return eventChan, nil
}

// readSSEventsChunk reads a chunk of Server-Sent Events from a buffered reader
func readSSEventsChunk(reader *bufio.Reader) ([]byte, error) {
	var buffer []byte

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				if len(buffer) > 0 {
					return buffer, nil
				}
				return nil, err
			}
			return nil, err
		}

		buffer = append(buffer, line...)

		if len(buffer) > 2 {
			if bytes.HasSuffix(buffer, []byte("\n\n")) {
				return buffer, nil
			}
		}
	}
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
