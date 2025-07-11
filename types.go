package sdk

import "time"

// ChatCompletionStreamResponseDelta represents a chat completion delta generated by streamed model responses.
type ChatCompletionStreamResponseDelta struct {
	Content          string                               `json:"content,omitempty"`
	ToolCalls        []ChatCompletionMessageToolCallChunk `json:"tool_calls,omitempty"`
	Role             string                               `json:"role,omitempty"`
	Reasoning        *string                              `json:"reasoning,omitempty"`
	ReasoningContent *string                              `json:"reasoning_content,omitempty"`
	Refusal          string                               `json:"refusal,omitempty"`
}

// ChatCompletionMessageToolCallChunk represents a chunk of a tool call in a stream response.
type ChatCompletionMessageToolCallChunk struct {
	Index    int    `json:"index"`
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Function struct {
		Name      string `json:"name,omitempty"`
		Arguments string `json:"arguments,omitempty"`
	} `json:"function,omitempty"`
}

// ChatCompletionTokenLogprob represents token log probability information.
type ChatCompletionTokenLogprob struct {
	Token   string  `json:"token"`
	Logprob float64 `json:"logprob"`
	Bytes   []int   `json:"bytes"`
}

// ChatCompletionStreamChoice represents a choice in a streaming chat completion response.
type ChatCompletionStreamChoice struct {
	Delta        ChatCompletionStreamResponseDelta `json:"delta"`
	Index        int                               `json:"index"`
	FinishReason string                            `json:"finish_reason"`
}

// CreateChatCompletionStreamResponse represents a streamed chunk of a chat completion response.
type CreateChatCompletionStreamResponse struct {
	ID                string                       `json:"id"`
	Choices           []ChatCompletionStreamChoice `json:"choices"`
	Created           int                          `json:"created"`
	Model             string                       `json:"model"`
	SystemFingerprint string                       `json:"system_fingerprint,omitempty"`
	Object            string                       `json:"object"`
	Usage             *CompletionUsage             `json:"usage,omitempty"`
}

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
}
