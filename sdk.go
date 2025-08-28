package sdk

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"strings"
	"syscall"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

// Client represents the SDK client interface
type Client interface {
	WithAuthToken(token string) *clientImpl
	WithTools(tools *[]ChatCompletionTool) *clientImpl
	WithOptions(options *CreateChatCompletionRequest) *clientImpl
	WithHeaders(headers map[string]string) *clientImpl
	WithHeader(name, value string) *clientImpl
	WithMiddlewareOptions(options *MiddlewareOptions) *clientImpl
	ListModels(ctx context.Context) (*ListModelsResponse, error)
	ListProviderModels(ctx context.Context, provider Provider) (*ListModelsResponse, error)
	ListTools(ctx context.Context) (*ListToolsResponse, error)
	ListAgents(ctx context.Context) (*ListAgentsResponse, error)
	GetAgent(ctx context.Context, id string) (*A2AAgentCard, error)
	GenerateContent(ctx context.Context, provider Provider, model string, messages []Message) (*CreateChatCompletionResponse, error)
	GenerateContentStream(ctx context.Context, provider Provider, model string, messages []Message) (<-chan SSEvent, error)
	HealthCheck(ctx context.Context) error
}

// isRetryableError determines if an error should trigger a retry
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	if netErr, ok := err.(net.Error); ok {
		if netErr.Timeout() {
			return true
		}
	}

	if opErr, ok := err.(*net.OpError); ok {
		if opErr.Op == "dial" || opErr.Op == "read" {
			return true
		}
		if sysErr, ok := opErr.Err.(*syscall.Errno); ok {
			return *sysErr == syscall.ECONNREFUSED || *sysErr == syscall.ECONNRESET
		}
	}

	if _, ok := err.(*net.DNSError); ok {
		return true
	}

	if err == context.DeadlineExceeded {
		return true
	}

	errStr := err.Error()
	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "EOF") ||
		strings.Contains(errStr, "timeout awaiting response headers")
}

// isRetryableStatusCode determines if an HTTP status code should trigger a retry
func isRetryableStatusCode(statusCode int, config *RetryConfig) bool {
	if len(config.RetryableStatusCodes) > 0 {
		for _, code := range config.RetryableStatusCodes {
			if statusCode == code {
				return true
			}
		}
		return false
	}

	switch statusCode {
	case
		http.StatusRequestTimeout,      // 408
		http.StatusTooManyRequests,     // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	default:
		return false
	}
}

// calculateBackoff calculates the backoff delay for exponential backoff
func calculateBackoff(attempt int, config *RetryConfig) time.Duration {
	if attempt == 0 {
		return 0
	}

	backoff := float64(config.InitialBackoffSec) * math.Pow(float64(config.BackoffMultiplier), float64(attempt-1))

	if backoff > float64(config.MaxBackoffSec) {
		backoff = float64(config.MaxBackoffSec)
	}

	return time.Duration(backoff) * time.Second
}

// getDefaultRetryConfig returns the default retry configuration
func getDefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		Enabled:           true,
		MaxAttempts:       3,
		InitialBackoffSec: 2,
		MaxBackoffSec:     30,
		BackoffMultiplier: 2,
	}
}

// clientImpl represents the concrete implementation of the SDK client
type clientImpl struct {
	baseURL     string        // Base URL of the Inference Gateway API
	http        *resty.Client // HTTP client for making requests
	token       string        // Authentication token
	tools       *[]ChatCompletionTool
	options     *CreateChatCompletionRequest // Custom request options
	retryConfig *RetryConfig                 // Retry configuration
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
//		Headers: map[string]string{
//			"X-Custom-Header": "custom-value",
//			"User-Agent": "my-app/1.0",
//		},
//	})
func NewClient(options *ClientOptions) Client {
	client := resty.New()

	if options.Timeout > 0 {
		client.SetTimeout(options.Timeout)
	}

	if options.APIKey != "" {
		client.SetAuthToken(options.APIKey)
	}

	if len(options.Headers) > 0 {
		client.SetHeaders(options.Headers)
	}

	retryConfig := options.RetryConfig
	if retryConfig == nil {
		retryConfig = getDefaultRetryConfig()
	}

	return &clientImpl{
		baseURL:     options.BaseURL,
		http:        client,
		token:       options.APIKey,
		tools:       options.Tools,
		options:     nil,
		retryConfig: retryConfig,
	}
}

// parseRetryAfter parses the Retry-After header and returns the delay duration
// The header can be either a delay in seconds or an HTTP-date
func parseRetryAfter(retryAfter string) (time.Duration, bool) {
	if retryAfter == "" {
		return 0, false
	}

	if seconds, err := time.ParseDuration(retryAfter + "s"); err == nil {
		return seconds, true
	}

	if retryTime, err := http.ParseTime(retryAfter); err == nil {
		delay := time.Until(retryTime)
		if delay > 0 {
			return delay, true
		}
	}

	return 0, false
}

// executeWithRetry executes an HTTP request with retry logic
func (c *clientImpl) executeWithRetry(ctx context.Context, request func() (*resty.Response, error)) (*resty.Response, error) {
	if !c.retryConfig.Enabled {
		return request()
	}

	var lastErr error
	var resp *resty.Response

	for attempt := 0; attempt < c.retryConfig.MaxAttempts; attempt++ {
		if attempt > 0 {
			var delay time.Duration

			if resp != nil && resp.StatusCode() == 429 {
				if retryAfterDelay, ok := parseRetryAfter(resp.Header().Get("Retry-After")); ok {
					delay = retryAfterDelay
				} else {
					delay = calculateBackoff(attempt, c.retryConfig)
				}
			} else {
				delay = calculateBackoff(attempt, c.retryConfig)
			}

			if c.retryConfig.OnRetry != nil {
				c.retryConfig.OnRetry(attempt, lastErr, delay)
			}

			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(delay):
			}
		}

		resp, lastErr = request()

		if lastErr == nil {
			if !resp.IsError() || !isRetryableStatusCode(resp.StatusCode(), c.retryConfig) {
				return resp, nil
			}
			lastErr = fmt.Errorf("HTTP %d", resp.StatusCode())
		}

		if !isRetryableError(lastErr) && (resp == nil || !isRetryableStatusCode(resp.StatusCode(), c.retryConfig)) {
			break
		}

		if ctx.Err() != nil {
			return resp, lastErr
		}
	}

	return resp, lastErr
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

// WithOptions sets custom request options for subsequent API calls.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//
//	// Set reasoning format for subsequent requests
//	reasoningFormat := "parsed"
//	options := &sdk.CreateChatCompletionRequest{
//		ReasoningFormat: &reasoningFormat,
//	}
//
//	// Use the options in a request
//	response, err := client.WithOptions(options).GenerateContent(
//		ctx,
//		sdk.Anthropic,
//		"anthropic/claude-3-opus-20240229",
//		messages,
//	)
//
// Notes:
//   - For GenerateContent calls, Stream will always be set to false regardless of options
//   - For GenerateContentStream calls, Stream will always be set to true regardless of options
//   - Model and Messages provided in the actual method calls will override options
//   - Options will persist for all future calls until cleared with WithOptions(nil)
func (c *clientImpl) WithOptions(options *CreateChatCompletionRequest) *clientImpl {
	c.options = options
	return c
}

// WithHeaders sets custom headers for the client.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	headers := map[string]string{
//		"X-Custom-Header": "value",
//	}
//	client = client.WithHeaders(headers)
//	resp, err := client.ListModels(ctx)
func (c *clientImpl) WithHeaders(headers map[string]string) *clientImpl {
	for name, value := range headers {
		c.http.Header.Set(name, value)
	}
	return c
}

// WithHeader sets a single custom header for the client.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	client = client.WithHeader("X-Custom-Header", "value")
//	resp, err := client.ListModels(ctx)
func (c *clientImpl) WithHeader(name, value string) *clientImpl {
	c.http.Header.Set(name, value)
	return c
}

// WithMiddlewareOptions sets middleware control options for subsequent API calls.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//	})
//	middlewareOpts := &sdk.MiddlewareOptions{
//		SkipMCP: true,
//		SkipA2A: true,
//	}
//	resp, err := client.WithMiddlewareOptions(middlewareOpts).GenerateContent(ctx, provider, model, messages)
//
// Note: This functionality requires the Inference Gateway to support the corresponding headers:
//   - X-MCP-Bypass: Skip MCP middleware processing
//   - X-A2A-Bypass: Skip A2A middleware processing
//   - X-Direct-Provider: Route directly to provider
func (c *clientImpl) WithMiddlewareOptions(options *MiddlewareOptions) *clientImpl {
	if options == nil {
		return c
	}

	if options.SkipMCP {
		c.http.Header.Set("X-MCP-Bypass", "true")
	} else {
		c.http.Header.Del("X-MCP-Bypass")
	}

	if options.SkipA2A {
		c.http.Header.Set("X-A2A-Bypass", "true")
	} else {
		c.http.Header.Del("X-A2A-Bypass")
	}

	if options.DirectProvider {
		c.http.Header.Set("X-Direct-Provider", "true")
	} else {
		c.http.Header.Del("X-Direct-Provider")
	}

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
	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			SetResult(&ListModelsResponse{}).
			Get(fmt.Sprintf("%s/models", c.baseURL))
	})

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
	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			SetResult(&ListModelsResponse{}).
			Get(fmt.Sprintf("%s/models?provider=%s", c.baseURL, provider))
	})

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errorResp Error
		if err := json.Unmarshal(resp.Body(), &errorResp); err == nil && errorResp.Error != nil {
			return nil, fmt.Errorf("API error: %s (status code: %d)", *errorResp.Error, resp.StatusCode())
		}

		errMsg := fmt.Sprintf("failed to list provider models, status code: %d", resp.StatusCode())

		if len(resp.Body()) > 0 {
			errMsg = fmt.Sprintf("%s, response body: %s", errMsg, string(resp.Body()))
		}

		return nil, fmt.Errorf("%s", errMsg)
	}

	result, ok := resp.Result().(*ListModelsResponse)
	if !ok || result == nil {
		return nil, fmt.Errorf("failed to parse response")
	}

	return result, nil
}

// ListTools returns all available MCP tools.
// Only accessible when EXPOSE_MCP is enabled on the server.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//		APIKey: "your-api-key",
//	})
//	ctx := context.Background()
//	tools, err := client.ListTools(ctx)
//	if err != nil {
//	    log.Fatalf("Error listing tools: %v", err)
//	}
//	fmt.Printf("Available tools: %+v\n", tools.Data)
func (c *clientImpl) ListTools(ctx context.Context) (*ListToolsResponse, error) {
	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			SetResult(&ListToolsResponse{}).
			Get(fmt.Sprintf("%s/mcp/tools", c.baseURL))
	})

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errorResp Error
		if err := json.Unmarshal(resp.Body(), &errorResp); err == nil && errorResp.Error != nil {
			return nil, fmt.Errorf("API error: %s (status code: %d)", *errorResp.Error, resp.StatusCode())
		}

		errMsg := fmt.Sprintf("failed to list MCP tools, status code: %d", resp.StatusCode())

		if len(resp.Body()) > 0 {
			errMsg = fmt.Sprintf("%s, response body: %s", errMsg, string(resp.Body()))
		}

		return nil, fmt.Errorf("%s", errMsg)
	}

	result, ok := resp.Result().(*ListToolsResponse)
	if !ok || result == nil {
		return nil, fmt.Errorf("failed to parse response")
	}

	return result, nil
}

// ListAgents returns all available A2A agents.
// Only accessible when EXPOSE_A2A is enabled on the server.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//		APIKey: "your-api-key",
//	})
//	ctx := context.Background()
//	agents, err := client.ListAgents(ctx)
//	if err != nil {
//	    log.Fatalf("Error listing agents: %v", err)
//	}
//	fmt.Printf("Available agents: %+v\n", agents.Data)
func (c *clientImpl) ListAgents(ctx context.Context) (*ListAgentsResponse, error) {
	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			SetResult(&ListAgentsResponse{}).
			Get(fmt.Sprintf("%s/a2a/agents", c.baseURL))
	})

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errorResp Error
		if err := json.Unmarshal(resp.Body(), &errorResp); err == nil && errorResp.Error != nil {
			return nil, fmt.Errorf("API error: %s (status code: %d)", *errorResp.Error, resp.StatusCode())
		}

		errMsg := fmt.Sprintf("failed to list A2A agents, status code: %d", resp.StatusCode())

		if len(resp.Body()) > 0 {
			errMsg = fmt.Sprintf("%s, response body: %s", errMsg, string(resp.Body()))
		}

		return nil, fmt.Errorf("%s", errMsg)
	}

	result, ok := resp.Result().(*ListAgentsResponse)
	if !ok || result == nil {
		return nil, fmt.Errorf("failed to parse response")
	}

	return result, nil
}

// GetAgent returns a specific A2A agent by its unique identifier.
// Only accessible when EXPOSE_A2A is enabled on the server.
//
// Example:
//
//	client := sdk.NewClient(&sdk.ClientOptions{
//		BaseURL: "http://localhost:8080/v1",
//		APIKey: "your-api-key",
//	})
//	ctx := context.Background()
//	agent, err := client.GetAgent(ctx, "agent-id-123")
//	if err != nil {
//	    log.Fatalf("Error getting agent: %v", err)
//	}
//	fmt.Printf("Agent details: %+v\n", agent)
func (c *clientImpl) GetAgent(ctx context.Context, id string) (*A2AAgentCard, error) {
	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			SetResult(&A2AAgentCard{}).
			Get(fmt.Sprintf("%s/a2a/agents/%s", c.baseURL, id))
	})

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errorResp Error
		if err := json.Unmarshal(resp.Body(), &errorResp); err == nil && errorResp.Error != nil {
			return nil, fmt.Errorf("API error: %s (status code: %d)", *errorResp.Error, resp.StatusCode())
		}

		errMsg := fmt.Sprintf("failed to get A2A agent, status code: %d", resp.StatusCode())

		if len(resp.Body()) > 0 {
			errMsg = fmt.Sprintf("%s, response body: %s", errMsg, string(resp.Body()))
		}

		return nil, fmt.Errorf("%s", errMsg)
	}

	result, ok := resp.Result().(*A2AAgentCard)
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
	request := CreateChatCompletionRequest{
		Model:    model,
		Messages: messages,
		Tools:    c.tools,
		Stream:   boolPtr(false),
	}

	if c.options != nil {
		options := *c.options

		if options.Model == "" {
			options.Model = request.Model
		}
		if len(options.Messages) == 0 {
			options.Messages = request.Messages
		}
		if options.Tools == nil && c.tools != nil {
			options.Tools = c.tools
		}

		options.Stream = boolPtr(false)

		request = options
	}

	queryParams := make(map[string]string)
	if provider != "" {
		queryParams["provider"] = string(provider)
	}

	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			SetQueryParams(queryParams).
			SetBody(request).
			SetResult(&CreateChatCompletionResponse{}).
			Post(fmt.Sprintf("%s/chat/completions", c.baseURL))
	})

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errorResp Error
		if err := json.Unmarshal(resp.Body(), &errorResp); err == nil && errorResp.Error != nil {
			return nil, fmt.Errorf("API error: %s (status code: %d)", *errorResp.Error, resp.StatusCode())
		}

		errMsg := fmt.Sprintf("failed to generate content, status code: %d", resp.StatusCode())

		if len(resp.Body()) > 0 {
			errMsg = fmt.Sprintf("%s, response body: %s", errMsg, string(resp.Body()))
		}

		return nil, fmt.Errorf("%s", errMsg)
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

	request := CreateChatCompletionRequest{
		Model:    model,
		Messages: messages,
		Stream:   boolPtr(true),
		Tools:    c.tools,
	}

	if c.options != nil {
		options := *c.options

		if options.Model == "" {
			options.Model = request.Model
		}
		if len(options.Messages) == 0 {
			options.Messages = request.Messages
		}
		if options.Tools == nil && c.tools != nil {
			options.Tools = c.tools
		}

		options.Stream = boolPtr(true)

		request = options
	}

	queryParams := make(map[string]string)
	if provider != "" {
		queryParams["provider"] = string(provider)
	}

	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			SetQueryParams(queryParams).
			SetBody(request).
			SetDoNotParseResponse(true).
			Post(fmt.Sprintf("%s/chat/completions", c.baseURL))
	})
	if err != nil {
		close(eventChan)
		return eventChan, err
	}

	if resp.IsError() {
		close(eventChan)

		body, _ := io.ReadAll(resp.RawBody())
		var errorResp Error
		if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Error != nil {
			return eventChan, fmt.Errorf("API stream error: %s (status code: %d)", *errorResp.Error, resp.StatusCode())
		}

		errMsg := fmt.Sprintf("stream request failed with status: %d", resp.StatusCode())

		if len(body) > 0 {
			errMsg = fmt.Sprintf("%s, response body: %s", errMsg, string(body))
		}

		return eventChan, fmt.Errorf("%s", errMsg)
	}

	rawBody := resp.RawBody()
	if rawBody == nil {
		close(eventChan)
		return eventChan, fmt.Errorf("empty response body")
	}

	go func() {
		defer close(eventChan)

		defer func() {
			err := rawBody.Close()
			require.NoError(nil, err, "failed to close response body")
		}()

		reader := bufio.NewReader(rawBody)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					errorData := []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
					eventChan <- SSEvent{
						Event: nil,
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
	resp, err := c.executeWithRetry(ctx, func() (*resty.Response, error) {
		return c.http.R().
			SetContext(ctx).
			Get(fmt.Sprintf("%s/health", c.baseURL))
	})

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("health check failed with status: %d", resp.StatusCode())
	}

	return nil
}
