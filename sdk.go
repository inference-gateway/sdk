package sdk

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Client represents the SDK client interface
type Client interface {
	WithAuthToken(token string) *clientImpl
	ListModels(ctx context.Context) (*ListModelsResponse, error)
	ListProviderModels(ctx context.Context, provider Provider) (*ListModelsResponse, error)
	GenerateContent(ctx context.Context, provider Provider, model string, messages []Message) (*CreateChatCompletionResponse, error)
	GenerateContentStream(ctx context.Context, provider Provider, model string, messages []Message) (<-chan SSEvent, error)
	WithTools(tools *[]ChatCompletionTool) *clientImpl
	HealthCheck(ctx context.Context) error
}

// clientImpl represents the concrete implementation of the SDK client
type clientImpl struct {
	baseURL string        // Base URL of the Inference Gateway API
	http    *resty.Client // HTTP client for making requests
	token   string        // Authentication token
	tools   *[]ChatCompletionTool
}

// NewClient creates a new SDK client with the specified options.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//		APIKey: "your-api-key",
//		Timeout: 30 * time.Second,
//		Tools: nil,
//	})
func NewClient(options *ClientOptions) Client {
	client := resty.New()

	// Set timeout if provided
	if options.Timeout > 0 {
		client.SetTimeout(options.Timeout)
	}

	// Set auth token if provided
	if options.APIKey != "" {
		client.SetAuthToken(options.APIKey)
	}

	return &clientImpl{
		baseURL: options.BaseURL,
		http:    client,
		token:   options.APIKey,
		tools:   options.Tools,
	}
}

// WithAuthToken sets the authentication token for the client.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	client = client.WithAuthToken("your-auth-token")
//	resp, err := client.ListModels(ctx)
func (c *clientImpl) WithAuthToken(token string) *clientImpl {
	c.token = token
	c.http.SetAuthToken(token)
	return c
}

// WithTools sets the tools for the client.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	tools := []sdk.ChatCompletionTool{
//		{
//			Name: "Weather",
//			Functions: []sdk.FunctionObject{
//				{
//					Name: "get_current_weather",
//					Description: stringPtr("Get the current weather in a given location"),
//					Parameters: &sdk.FunctionParameters{
//						Type: stringPtr("object"),
//						Properties: &map[string]interface{}{
//							"location": map[string]interface{}{
//								"type":        "string",
//								"description": "The city and state, e.g. San Francisco, CA",
//							},
//							"unit": map[string]interface{}{
//								"type":        "string",
//								"enum":        []string{"celsius", "fahrenheit"},
//								"description": "The temperature unit to use",
//							},
//						},
//						Required: &[]string{"location"},
//					},
//				},
//			},
//		},
//	}
//	resp, err = client.WithTools(tools).GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
func (c *clientImpl) WithTools(tools *[]ChatCompletionTool) *clientImpl {
	c.tools = tools
	return c
}

// ListModels returns all available language models from all providers.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	ctx := context.Background()
//	models, err := client.ListModels(ctx)
//	if err != nil {
//	    log.Fatalf("Error listing models: %v", err)
//	}
//	fmt.Printf("Available models: %+v\n", models)
func (c *clientImpl) ListModels(ctx context.Context) (*ListModelsResponse, error) {
	resp, err := c.http.R().
		SetContext(ctx).
		SetResult(&ListModelsResponse{}).
		Get(fmt.Sprintf("%s/models", c.baseURL))

	if err != nil {
		return &ListModelsResponse{}, err
	}

	if resp.IsError() {
		return &ListModelsResponse{}, fmt.Errorf("failed to list models, status code: %d", resp.StatusCode())
	}

	result, ok := resp.Result().(*ListModelsResponse)
	if !ok || result == nil {
		return &ListModelsResponse{}, fmt.Errorf("failed to parse response")
	}

	return result, nil
}

// ListProviderModels returns all available language models for a specific provider.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	ctx := context.Background()
//	resp, err := client.ListProviderModels(ctx, sdk.Ollama)
//	if err != nil {
//	    log.Fatalf("Error listing models: %v", resp)
//	}
//	fmt.Printf("Provider: %s", resp.Provider)
//	fmt.Printf("Available models: %+v\n", resp.Data)
func (c *clientImpl) ListProviderModels(ctx context.Context, provider Provider) (*ListModelsResponse, error) {
	resp, err := c.http.R().
		SetContext(ctx).
		SetResult(&ListModelsResponse{}).
		Get(fmt.Sprintf("%s/models?provider=%s", c.baseURL, provider))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errorResp Error
		if err := json.Unmarshal(resp.Body(), &errorResp); err == nil && errorResp.Error != nil {
			return nil, fmt.Errorf("API error: %s", *errorResp.Error)
		}
		return nil, fmt.Errorf("failed to list provider models, status code: %d", resp.StatusCode())
	}

	result, ok := resp.Result().(*ListModelsResponse)
	if !ok || result == nil {
		return nil, fmt.Errorf("failed to parse response")
	}

	return result, nil
}

// GenerateContent generates content using the specified provider and model.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	ctx := context.Background()
//	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
//	defer cancel()
//
//	response, err := client.GenerateContent(
//	    ctx,
//	    sdk.Ollama,
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
func (c *clientImpl) GenerateContent(ctx context.Context, provider Provider, model string, messages []Message) (*CreateChatCompletionResponse, error) {
	requestBody := CreateChatCompletionRequest{
		Model:    model,
		Messages: messages,
		Tools:    c.tools,
	}

	queryParams := make(map[string]string)
	if provider != "" {
		queryParams["provider"] = string(provider)
	}

	resp, err := c.http.R().
		SetContext(ctx).
		SetQueryParams(queryParams).
		SetBody(requestBody).
		SetResult(&CreateChatCompletionResponse{}).
		Post(fmt.Sprintf("%s/chat/completions", c.baseURL))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errorResp Error
		if err := json.Unmarshal(resp.Body(), &errorResp); err == nil && errorResp.Error != nil {
			return nil, fmt.Errorf("API error: %s", *errorResp.Error)
		}
		return nil, fmt.Errorf("failed to generate content, status code: %d", resp.StatusCode())
	}

	result, ok := resp.Result().(*CreateChatCompletionResponse)
	if !ok || result == nil {
		return nil, fmt.Errorf("failed to parse response")
	}

	return result, nil
}

// GenerateContentStream generates content using streaming mode and returns a channel of events.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	ctx := context.Background()
//	events, err := client.GenerateContentStream(
//		ctx,
//		sdk.Ollama,
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
//			var streamResponse CreateChatCompletionStreamResponse
//			if err := json.Unmarshal(*event.Data, &streamResponse); err != nil {
//				log.Printf("Error parsing stream response: %v", err)
//				continue
//			}
//
//			for _, choice := range streamResponse.Choices {
//				if choice.Delta.Content != "" {
//					log.Printf("Content: %s", choice.Delta.Content)
//				}
//			}
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
	eventChan := make(chan SSEvent, 100)

	requestBody := CreateChatCompletionRequest{
		Model:    model,
		Messages: messages,
		Stream:   boolPtr(true),
		Tools:    c.tools,
	}

	queryParams := make(map[string]string)
	if provider != "" {
		queryParams["provider"] = string(provider)
	}

	req := c.http.R().
		SetContext(ctx).
		SetQueryParams(queryParams).
		SetBody(requestBody).
		SetDoNotParseResponse(true)

	resp, err := req.Post(fmt.Sprintf("%s/chat/completions", c.baseURL))
	if err != nil {
		close(eventChan)
		return eventChan, err
	}

	if resp.IsError() {
		close(eventChan)
		return eventChan, fmt.Errorf("stream request failed with status: %d", resp.StatusCode())
	}

	rawBody := resp.RawBody()
	if rawBody == nil {
		close(eventChan)
		return eventChan, fmt.Errorf("empty response body")
	}

	go func() {
		defer close(eventChan)
		defer rawBody.Close()

		reader := bufio.NewReader(rawBody)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					errorData := []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
					eventChan <- SSEvent{
						Event: nil, // TODO - need to add error event type to enum in OpenAPI spec, but it's not very important because all providers following OpenAI and the event section is not provided in the SSEvents, i.e it's an optional and will be treated as a "nice-to-have"
						Data:  &errorData,
					}
				}
				return
			}

			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if !strings.HasPrefix(line, "data: ") {
				continue
			}

			data := strings.TrimPrefix(line, "data: ")

			if data == "[DONE]" {
				streamEnd := StreamEnd
				eventChan <- SSEvent{
					Event: &streamEnd,
				}
				return
			}

			contentDelta := ContentDelta
			dataBytes := []byte(data)
			eventChan <- SSEvent{
				Event: &contentDelta,
				Data:  &dataBytes,
			}
		}
	}()

	return eventChan, nil
}

func boolPtr(b bool) *bool {
	return &b
}

// HealthCheck performs a health check request to verify API availability.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
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
