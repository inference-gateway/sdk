// Package sdk provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package sdk

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for ChatCompletionChoiceFinishReason.
const (
	ContentFilter ChatCompletionChoiceFinishReason = "content_filter"
	FunctionCall  ChatCompletionChoiceFinishReason = "function_call"
	Length        ChatCompletionChoiceFinishReason = "length"
	Stop          ChatCompletionChoiceFinishReason = "stop"
	ToolCalls     ChatCompletionChoiceFinishReason = "tool_calls"
)

// Defines values for ChatCompletionToolType.
const (
	Function ChatCompletionToolType = "function"
)

// Defines values for MessageRole.
const (
	Assistant MessageRole = "assistant"
	System    MessageRole = "system"
	Tool      MessageRole = "tool"
	User      MessageRole = "user"
)

// Defines values for Providers.
const (
	Anthropic  Providers = "anthropic"
	Cloudflare Providers = "cloudflare"
	Cohere     Providers = "cohere"
	Groq       Providers = "groq"
	Ollama     Providers = "ollama"
	Openai     Providers = "openai"
)

// Defines values for SSEventEvent.
const (
	ContentDelta SSEventEvent = "content-delta"
	ContentEnd   SSEventEvent = "content-end"
	ContentStart SSEventEvent = "content-start"
	MessageEnd   SSEventEvent = "message-end"
	MessageStart SSEventEvent = "message-start"
	StreamEnd    SSEventEvent = "stream-end"
	StreamStart  SSEventEvent = "stream-start"
)

// ChatCompletionChoice defines model for ChatCompletionChoice.
type ChatCompletionChoice struct {
	// FinishReason The reason the model stopped generating tokens. This will be `stop` if the model hit a natural stop point or a provided stop sequence,
	// `length` if the maximum number of tokens specified in the request was reached,
	// `content_filter` if content was omitted due to a flag from our content filters,
	// `tool_calls` if the model called a tool.
	FinishReason ChatCompletionChoiceFinishReason `json:"finish_reason"`

	// Index The index of the choice in the list of choices.
	Index int `json:"index"`

	// Message Message structure for provider requests
	Message Message `json:"message"`
}

// ChatCompletionChoiceFinishReason The reason the model stopped generating tokens. This will be `stop` if the model hit a natural stop point or a provided stop sequence,
// `length` if the maximum number of tokens specified in the request was reached,
// `content_filter` if content was omitted due to a flag from our content filters,
// `tool_calls` if the model called a tool.
type ChatCompletionChoiceFinishReason string

// ChatCompletionMessageToolCall defines model for ChatCompletionMessageToolCall.
type ChatCompletionMessageToolCall struct {
	// Function The function that the model called.
	Function ChatCompletionMessageToolCallFunction `json:"function"`

	// Id The ID of the tool call.
	Id string `json:"id"`

	// Type The type of the tool. Currently, only `function` is supported.
	Type ChatCompletionToolType `json:"type"`
}

// ChatCompletionMessageToolCallFunction The function that the model called.
type ChatCompletionMessageToolCallFunction struct {
	// Arguments The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.
	Arguments string `json:"arguments"`

	// Name The name of the function to call.
	Name string `json:"name"`
}

// ChatCompletionToolType The type of the tool. Currently, only `function` is supported.
type ChatCompletionToolType string

// CompletionUsage Usage statistics for the completion request.
type CompletionUsage struct {
	// CompletionTokens Number of tokens in the generated completion.
	CompletionTokens int64 `json:"completion_tokens"`

	// PromptTokens Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens"`

	// TotalTokens Total number of tokens used in the request (prompt + completion).
	TotalTokens int64 `json:"total_tokens"`
}

// Error defines model for Error.
type Error struct {
	Error *string `json:"error,omitempty"`
}

// ListModelsResponse Response structure for listing models
type ListModelsResponse struct {
	Data     *[]Model `json:"data,omitempty"`
	Object   *string  `json:"object,omitempty"`
	Provider *string  `json:"provider,omitempty"`
}

// Message Message structure for provider requests
type Message struct {
	Content   string  `json:"content"`
	Reasoning *string `json:"reasoning,omitempty"`

	// Role Role of the message sender
	Role       MessageRole                      `json:"role"`
	ToolCallId *string                          `json:"tool_call_id,omitempty"`
	ToolCalls  *[]ChatCompletionMessageToolCall `json:"tool_calls,omitempty"`
}

// MessageRole Role of the message sender
type MessageRole string

// Model Common model information
type Model struct {
	Created  *int64  `json:"created,omitempty"`
	Id       *string `json:"id,omitempty"`
	Object   *string `json:"object,omitempty"`
	OwnedBy  *string `json:"owned_by,omitempty"`
	ServedBy *string `json:"served_by,omitempty"`
}

// ProviderSpecificResponse Provider-specific response format. Examples:
//
// OpenAI GET /v1/models?provider=openai response:
// ```json
//
//	{
//	  "provider": "openai",
//	  "object": "list",
//	  "data": [
//	    {
//	      "id": "gpt-4",
//	      "object": "model",
//	      "created": 1687882410,
//	      "owned_by": "openai",
//	      "served_by": "openai"
//	    }
//	  ]
//	}
//
// ```
//
// Anthropic GET /v1/models?provider=anthropic response:
// ```json
//
//	{
//	  "provider": "anthropic",
//	  "object": "list",
//	  "data": [
//	    {
//	      "id": "gpt-4",
//	      "object": "model",
//	      "created": 1687882410,
//	      "owned_by": "openai",
//	      "served_by": "openai"
//	    }
//	  ]
//	}
//
// ```
type ProviderSpecificResponse map[string]interface{}

// Providers defines model for Providers.
type Providers string

// SSEvent defines model for SSEvent.
type SSEvent struct {
	Data *struct {
		Choices *[]ChatCompletionChoice `json:"choices,omitempty"`
		Created *int64                  `json:"created,omitempty"`
		Id      *string                 `json:"id,omitempty"`
		Model   *string                 `json:"model,omitempty"`
		Object  *string                 `json:"object,omitempty"`

		// Usage Usage statistics for the completion request.
		Usage *CompletionUsage `json:"usage,omitempty"`
	} `json:"data,omitempty"`
	Event *SSEventEvent `json:"event,omitempty"`
	Id    *string       `json:"id,omitempty"`
	Retry *int          `json:"retry,omitempty"`
}

// SSEventEvent defines model for SSEvent.Event.
type SSEventEvent string

// BadRequest defines model for BadRequest.
type BadRequest = Error

// InternalError defines model for InternalError.
type InternalError = Error

// ProviderResponse Provider-specific response format. Examples:
//
// OpenAI GET /v1/models?provider=openai response:
// ```json
//
//	{
//	  "provider": "openai",
//	  "object": "list",
//	  "data": [
//	    {
//	      "id": "gpt-4",
//	      "object": "model",
//	      "created": 1687882410,
//	      "owned_by": "openai",
//	      "served_by": "openai"
//	    }
//	  ]
//	}
//
// ```
//
// Anthropic GET /v1/models?provider=anthropic response:
// ```json
//
//	{
//	  "provider": "anthropic",
//	  "object": "list",
//	  "data": [
//	    {
//	      "id": "gpt-4",
//	      "object": "model",
//	      "created": 1687882410,
//	      "owned_by": "openai",
//	      "served_by": "openai"
//	    }
//	  ]
//	}
//
// ```
type ProviderResponse = ProviderSpecificResponse

// Unauthorized defines model for Unauthorized.
type Unauthorized = Error

// ProviderRequest defines model for ProviderRequest.
type ProviderRequest struct {
	Messages *[]struct {
		Content *string `json:"content,omitempty"`
		Role    *string `json:"role,omitempty"`
	} `json:"messages,omitempty"`
	Model       *string  `json:"model,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`
}

// ProxyPatchJSONBody defines parameters for ProxyPatch.
type ProxyPatchJSONBody struct {
	Messages *[]struct {
		Content *string `json:"content,omitempty"`
		Role    *string `json:"role,omitempty"`
	} `json:"messages,omitempty"`
	Model       *string  `json:"model,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`
}

// ProxyPostJSONBody defines parameters for ProxyPost.
type ProxyPostJSONBody struct {
	Messages *[]struct {
		Content *string `json:"content,omitempty"`
		Role    *string `json:"role,omitempty"`
	} `json:"messages,omitempty"`
	Model       *string  `json:"model,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`
}

// ProxyPutJSONBody defines parameters for ProxyPut.
type ProxyPutJSONBody struct {
	Messages *[]struct {
		Content *string `json:"content,omitempty"`
		Role    *string `json:"role,omitempty"`
	} `json:"messages,omitempty"`
	Model       *string  `json:"model,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`
}

// PostV1ChatCompletionsJSONBody defines parameters for PostV1ChatCompletions.
type PostV1ChatCompletionsJSONBody struct {
	MaxTokens *int      `json:"max_tokens,omitempty"`
	Messages  []Message `json:"messages"`

	// Model Model ID to use
	Model       string                    `json:"model"`
	Stream      *bool                     `json:"stream,omitempty"`
	Temperature *float32                  `json:"temperature,omitempty"`
	Tools       *[]map[string]interface{} `json:"tools,omitempty"`
}

// PostV1ChatCompletionsParams defines parameters for PostV1ChatCompletions.
type PostV1ChatCompletionsParams struct {
	// Provider Specific provider to use (default determined by model)
	Provider *Providers `form:"provider,omitempty" json:"provider,omitempty"`
}

// ListModelsParams defines parameters for ListModels.
type ListModelsParams struct {
	// Provider Specific provider to query (optional)
	Provider *Providers `form:"provider,omitempty" json:"provider,omitempty"`
}

// ProxyPatchJSONRequestBody defines body for ProxyPatch for application/json ContentType.
type ProxyPatchJSONRequestBody ProxyPatchJSONBody

// ProxyPostJSONRequestBody defines body for ProxyPost for application/json ContentType.
type ProxyPostJSONRequestBody ProxyPostJSONBody

// ProxyPutJSONRequestBody defines body for ProxyPut for application/json ContentType.
type ProxyPutJSONRequestBody ProxyPutJSONBody

// PostV1ChatCompletionsJSONRequestBody defines body for PostV1ChatCompletions for application/json ContentType.
type PostV1ChatCompletionsJSONRequestBody PostV1ChatCompletionsJSONBody
