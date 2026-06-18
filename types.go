package sdk

import (
	"fmt"
	"time"
)

// ClientOptions represents the options that can be passed to the client.
type ClientOptions struct {
	// APIKey is the API key to use for the client.
	APIKey string
	// BaseURL is the base URL to use for the client.
	BaseURL string
	// Timeout is the timeout to use for the client.
	Timeout time.Duration
	// Tools is the tools to use for the client.
	Tools *[]ChatCompletionTool
	// Headers is a map of custom headers to include with all requests.
	Headers map[string]string
	// RetryConfig is the retry configuration for HTTP requests.
	RetryConfig *RetryConfig
}

// RetryConfig represents the retry configuration for HTTP requests
type RetryConfig struct {
	// Enabled controls whether retry logic is enabled
	Enabled bool
	// MaxAttempts is the maximum number of retry attempts (including initial request)
	MaxAttempts int
	// InitialBackoffSec is the initial backoff delay in seconds
	InitialBackoffSec int
	// MaxBackoffSec is the maximum backoff delay in seconds
	MaxBackoffSec int
	// BackoffMultiplier is the multiplier for exponential backoff
	BackoffMultiplier int
	// RetryableStatusCodes is a custom list of HTTP status codes that should trigger a retry.
	// If nil or empty, uses default status codes (408, 429, 500, 502, 503, 504)
	RetryableStatusCodes []int
	// OnRetry is called before each retry attempt with attempt number, error, and delay.
	// The attempt number starts from 1 for the first retry (after initial request fails)
	OnRetry func(attempt int, err error, delay time.Duration)
}

// MiddlewareOptions represents options for controlling middleware behavior
type MiddlewareOptions struct {
	// SkipMCP bypasses MCP middleware processing
	SkipMCP bool
	// DirectProvider routes directly to provider without middleware
	DirectProvider bool
}

// Helper functions for working with vision messages

// NewMessageContent creates MessageContent from either a string or []ContentPart.
// This provides backward compatibility similar to OpenAI's SDK pattern.
//
// Usage:
//
//	content := NewMessageContent("Hello world")              // string
//	content := NewMessageContent([]ContentPart{...})         // multimodal
func NewMessageContent[T string | []ContentPart](value T) MessageContent {
	var content MessageContent
	var err error

	switch v := any(value).(type) {
	case string:
		err = content.FromMessageContent0(v)
	case []ContentPart:
		err = content.FromMessageContent1(v)
	}

	if err != nil {
		panic(fmt.Sprintf("failed to create message content: %v", err))
	}
	return content
}

// NewTextMessage creates a message with text content (backward compatible).
func NewTextMessage(role MessageRole, text string) (Message, error) {
	var msg Message
	msg.Role = role
	err := msg.Content.FromMessageContent0(text)
	return msg, err
}

// NewImageMessage creates a message with multimodal content (text + images).
func NewImageMessage(role MessageRole, parts []ContentPart) (Message, error) {
	var msg Message
	msg.Role = role
	err := msg.Content.FromMessageContent1(parts)
	return msg, err
}

// NewTextContentPart creates a text content part for multimodal messages.
func NewTextContentPart(text string) (ContentPart, error) {
	var part ContentPart
	err := part.FromTextContentPart(TextContentPart{
		Type: TextContentPartTypeText,
		Text: text,
	})
	return part, err
}

// NewImageContentPart creates an image content part for multimodal messages.
func NewImageContentPart(imageURL string, detail *ImageURLDetail) (ContentPart, error) {
	var part ContentPart
	err := part.FromImageContentPart(ImageContentPart{
		Type: ImageContentPartTypeImageURL,
		ImageURL: ImageURL{
			URL:    imageURL,
			Detail: detail,
		},
	})
	return part, err
}
