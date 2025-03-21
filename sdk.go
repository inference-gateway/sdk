package sdk

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

// Client represents the SDK client interface
type Client interface {
	ListModels(ctx context.Context) ([]ListModelsResponse, error)
	ListProviderModels(ctx context.Context, provider Provider) ([]Model, error)
	GenerateContent(ctx context.Context, provider Provider, model string, messages []Message) (*GenerateResponse, error)
	GenerateContentStream(ctx context.Context, provider Provider, model string, messages []Message) (<-chan SSEvent, error)
	HealthCheck(ctx context.Context) error
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
//	ctx := context.Background()
//	models, err := client.ListModels(ctx)
//	if err != nil {
//	    log.Fatalf("Error listing models: %v", err)
//	}
//	fmt.Printf("Available models: %+v\n", models)
func (c *clientImpl) ListModels(ctx context.Context) ([]ListModelsResponse, error) {
	// TODO - implement it properly
}

// ListProviderModels returns all available language models for a specific provider.
//
// Example:
//
//	ctx := context.Background()
//	models, err := client.ListProviderModels(sdk.ProviderOllama)
//	if err != nil {
//	    log.Fatalf("Error listing models: %v", err)
//	}
//	fmt.Printf("Available models: %+v\n", models)
func (c *clientImpl) ListProviderModels(ctx context.Context, provider Provider) ([]Model, error) {
	// TODO - implement it properly
}

// GenerateContent generates content using the specified provider and model.
//
// Example:
//
//	ctx := context.Background()
//	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
//	defer cancel()
//
//	response, err := client.GenerateContent(
//	    ctx,
//	    sdk.ProviderOllama,
//	    "llama2",
//	    []sdk.Message{
//	        {
//	            Role:    sdk.MessageRoleSystem,
//	            Content: "You are a helpful assistant.",
//	        },
//	        {
//	            Role:    sdk.MessageRoleUser,
//	            Content: "What is Go?",
//	        },
//	    },
//	)
//	if err != nil {
//	    log.Fatalf("Error generating content: %v", err)
//	}
//	fmt.Printf("Generated content: %s\n", response.Response.Content)
func (c *clientImpl) GenerateContent(ctx context.Context, provider Provider, model string, messages []Message) (*GenerateResponse, error) {
	// TODO - implement it properly

	// return &result, nil
}

// GenerateContentStream generates content using streaming mode and returns a channel of events.
//
// Example:
//
//	ctx := context.Background()
//	events, err := client.GenerateContentStream(
//		ctx,
//		sdk.ProviderOllama,
//		"llama2",
//		[]sdk.Message{
//			{Role: sdk.MessageRoleSystem, Content: "You are a helpful assistant."},
//			{Role: sdk.MessageRoleUser, Content: "What is Go?"},
//		},
//	)
//	if err != nil {
//		log.Fatalf("Error: %v", err)
//	}
//
//	for event := range events {
//		switch event.Event {
//		case sdk.StreamEventContentDelta:
//			var delta struct {
//				Content string `json:"content"`
//			}
//			if err := json.Unmarshal(event.Data, &delta); err != nil {
//				log.Printf("Error parsing delta: %v", err)
//				continue
//			}
//			fmt.Print(delta.Content)
//		case sdk.StreamEventMessageError:
//			var errResp struct {
//				Error string `json:"error"`
//			}
//			if err := json.Unmarshal(event.Data, &errResp); err != nil {
//				log.Printf("Error parsing error: %v", err)
//				continue
//			}
//			log.Printf("Error: %s", errResp.Error)
//		}
//	}
func (c *clientImpl) GenerateContentStream(ctx context.Context, provider Provider, model string, messages []Message) (<-chan SSEvent, error) {
	// TODO - implement it properly - send the stream as-is
	ssevent := make(chan SSEvent, 100)

	return ssevent, nil
}

// HealthCheck performs a health check request to verify API availability.
//
// Example:
//
//	ctx := context.Background()
//	err := client.HealthCheck(ctx)
//	if err != nil {
//	    log.Fatalf("Health check failed: %v", err)
//	}
func (c *clientImpl) HealthCheck(ctx context.Context) error {
	resp, err := c.http.R().
		SetContext(ctx).
		Get(fmt.Sprintf("%s/health", c.baseURL))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("health check failed with status: %d", resp.StatusCode())
	}

	return nil
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

// ParseSSEvents parses a Server-Sent Event from a byte slice
func parseSSEvents(line []byte) (*SSEvent, error) {
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
			event.Data = &value
		case "event":
			eventVal := SSEventEvent(string(value))
			event.Event = &eventVal
		}
	}

	return event, nil
}
