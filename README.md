<div align="center">

# 🚀 Inference Gateway Go SDK

### A powerful and easy-to-use Go SDK for the Inference Gateway

[![Go Reference](https://pkg.go.dev/badge/github.com/inference-gateway/sdk.svg)](https://pkg.go.dev/github.com/inference-gateway/sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/inference-gateway/sdk)](https://goreportcard.com/report/github.com/inference-gateway/sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/inference-gateway/sdk.svg)](https://github.com/inference-gateway/sdk/releases)
[![Go Version](https://img.shields.io/github/go-mod/go-version/inference-gateway/sdk)](https://golang.org/)

Connect to multiple LLM providers through a unified interface • Stream responses • Function calling • MCP tools support • Middleware control

[Installation](#installation) • [Quick Start](#usage) • [Examples](#examples) • [Documentation](#documentation)

</div>

---

- [🚀 Inference Gateway Go SDK](#-inference-gateway-go-sdk)
    - [A powerful and easy-to-use Go SDK for the Inference Gateway](#a-powerful-and-easy-to-use-go-sdk-for-the-inference-gateway)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Client](#creating-a-client)
    - [Using Custom Headers](#using-custom-headers)
    - [Middleware Options](#middleware-options)
    - [Listing Models](#listing-models)
    - [Listing MCP Tools](#listing-mcp-tools)
    - [Generating Content](#generating-content)
    - [Using ReasoningFormat](#using-reasoningformat)
    - [Streaming Content](#streaming-content)
    - [Tool-Use](#tool-use)
    - [Health Check](#health-check)
  - [Examples](#examples)
  - [Supported Providers](#supported-providers)
  - [Documentation](#documentation)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

To install the SDK, use `go get`:

```sh
go get github.com/inference-gateway/sdk
```

## Usage

### Creating a Client

To create a client, use the `NewClient` function:

```go
package main

import (
    "fmt"
    "log"

    sdk "github.com/inference-gateway/sdk"
)

func main() {
    client := sdk.NewClient(&sdk.ClientOptions{
        BaseURL: "http://localhost:8080/v1",
    })
}
```

### Using Custom Headers

The SDK supports custom HTTP headers that can be included with all requests. You can set headers in three ways:

1. **Initial headers via ClientOptions:**

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
    Headers: map[string]string{
        "X-Custom-Header": "my-value",
        "User-Agent":      "my-app/1.0",
    },
})
```

2. **Multiple headers using WithHeaders:**

```go
client = client.WithHeaders(map[string]string{
    "X-Request-ID": "abc123",
    "X-Source":     "sdk",
})
```

3. **Single header using WithHeader:**

```go
client = client.WithHeader("Authorization", "Bearer token123")
```

Headers can be combined and will override previous values if the same header name is used:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
    Headers: map[string]string{
        "X-App-Name": "my-app",
    },
})

// Add more headers
client = client.WithHeaders(map[string]string{
    "X-Request-ID": "req-123",
    "X-Version":    "1.0",
}).WithHeader("Authorization", "Bearer token")

// All subsequent requests will include all these headers
response, err := client.GenerateContent(ctx, provider, model, messages)
```

### Middleware Options

The Inference Gateway supports various middleware layers (MCP tools, A2A agents) that can be bypassed for direct provider access. The SDK provides `WithMiddlewareOptions` to control middleware behavior:

```go
package main

import (
    "context"
    "fmt"
    "log"

    sdk "github.com/inference-gateway/sdk"
)

func main() {
    client := sdk.NewClient(&sdk.ClientOptions{
        BaseURL: "http://localhost:8080/v1",
        APIKey:  "your-api-key",
    })

    ctx := context.Background()
    messages := []sdk.Message{
        {Role: sdk.User, Content: "Hello, world!"},
    }

    // 1. Skip MCP middleware only
    response1, err := client.WithMiddlewareOptions(&sdk.MiddlewareOptions{
        SkipMCP: true,
    }).GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)

    // 2. Skip A2A (Agent-to-Agent) middleware only
    response2, err := client.WithMiddlewareOptions(&sdk.MiddlewareOptions{
        SkipA2A: true,
    }).GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)

    // 3. Direct provider access (bypasses all middleware)
    response3, err := client.WithMiddlewareOptions(&sdk.MiddlewareOptions{
        DirectProvider: true,
    }).GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)

    // 4. Skip both MCP and A2A middleware
    response4, err := client.WithMiddlewareOptions(&sdk.MiddlewareOptions{
        SkipMCP: true,
        SkipA2A: true,
    }).GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
}
```

**Middleware Options:**

-   **`SkipMCP`** - Bypasses MCP (Model Context Protocol) middleware processing
-   **`SkipA2A`** - Bypasses A2A (Agent-to-Agent) middleware processing
-   **`DirectProvider`** - Routes directly to the provider without any middleware

**Method Chaining:**

Middleware options can be chained with other configuration methods:

```go
response, err := client.
    WithHeader("X-Custom-Header", "value").
    WithMiddlewareOptions(&sdk.MiddlewareOptions{
        SkipMCP: true,
        SkipA2A: true,
    }).
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

**Alternative Header Approach:**

You can also control middleware using custom headers directly:

```go
response, err := client.
    WithHeader("X-MCP-Bypass", "true").
    WithHeader("X-A2A-Bypass", "true").
    WithHeader("X-Direct-Provider", "true").
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

> **Note:** Middleware options apply to all subsequent API calls until overridden with a new `WithMiddlewareOptions` call. The gateway must support the corresponding headers for this functionality to work properly.

### Listing Models

To list available models, use the ListModels method:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
})

ctx := context.Background()

// List all models from all providers
resp, err := client.ListModels(ctx)
if err != nil {
    log.Fatalf("Error listing models: %v", err)
}
fmt.Printf("All available models: %+v\n", resp.Data)

// List models for a specific provider
groqResp, err := client.ListProviderModels(ctx, sdk.Groq)
if err != nil {
    log.Fatalf("Error listing provider models: %v", err)
}
fmt.Printf("Provider: %s\n", *groqResp.Provider)
fmt.Printf("Available Groq models: %+v\n", groqResp.Data)
```

### Listing MCP Tools

To list available MCP (Model Context Protocol) tools, use the `ListTools` method. This functionality is only available when `EXPOSE_MCP=true` is set on the Inference Gateway server:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
    APIKey:  "your-api-key", // Required for MCP tools access
})

ctx := context.Background()
tools, err := client.ListTools(ctx)
if err != nil {
    log.Fatalf("Error listing tools: %v", err)
}

fmt.Printf("Found %d MCP tools:\n", len(tools.Data))
for _, tool := range tools.Data {
    fmt.Printf("- %s: %s (Server: %s)\n", tool.Name, tool.Description, tool.Server)
    if tool.InputSchema != nil {
        fmt.Printf("  Input Schema: %+v\n", *tool.InputSchema)
    }
}
```

> **Note:** The MCP tools endpoint requires authentication and is only accessible when the server has `EXPOSE_MCP=true` configured. If the endpoint is not exposed, you'll receive a 403 error with the message "MCP tools endpoint is not exposed. Set EXPOSE_MCP=true to enable."

### Generating Content

To generate content using a model, use the GenerateContent method:

> **Note:** Some models support reasoning capabilities. You can use the `ReasoningFormat` parameter to control how reasoning is provided in the response. The model's reasoning will be available in the `Reasoning` or `ReasoningContent` fields of the response message.

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
})

ctx := context.Background()
response, err := client.GenerateContent(
    ctx,
    sdk.Ollama,
    "ollama/llama2",
    []sdk.Message{
        {
            Role:    sdk.System,
            Content: "You are a helpful assistant.",
        },
        {
            Role:    sdk.User,
            Content: "What is Go?",
        },
    },
)
if err != nil {
    log.Printf("Error generating content: %v", err)
    return
}

var chatCompletion CreateChatCompletionResponse
if err := json.Unmarshal(response.RawResponse, &chatCompletion); err != nil {
    log.Printf("Error unmarshaling response: %v", err)
    return
}

fmt.Printf("Generated content: %s\n", chatCompletion.Choices[0].Message.Content)

// If reasoning was requested and the model supports it
if chatCompletion.Choices[0].Message.Reasoning != nil {
    fmt.Printf("Reasoning: %s\n", *chatCompletion.Choices[0].Message.Reasoning)
}
```

### Using ReasoningFormat

You can enable reasoning capabilities by setting the ReasoningFormat parameter in your request:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
})

ctx := context.Background()

// Set up your messages
messages := []sdk.Message{
    {
        Role:    sdk.System,
        Content: "You are a helpful assistant. Please include your reasoning for complex questions.",
    },
    {
        Role:    sdk.User,
        Content: "What is the square root of 144 and why?",
    },
}

// Create a request with reasoning format
reasoningFormat := "parsed"  // Use "raw" or "parsed" - default to "parsed" if not specified
options := &sdk.CreateChatCompletionRequest{
    ReasoningFormat: &reasoningFormat,
}

// Set options and make the request
response, err := client.WithOptions(options).GenerateContent(
    ctx,
    sdk.Anthropic,
    "anthropic/claude-3-opus-20240229",
    messages,
)

if err != nil {
    log.Fatalf("Error generating content: %v", err)
}

fmt.Printf("Content: %s\n", response.Choices[0].Message.Content)
if response.Choices[0].Message.Reasoning != nil {
    fmt.Printf("Reasoning: %s\n", *response.Choices[0].Message.Reasoning)
}
```

### Streaming Content

To generate content using streaming mode, use the GenerateContentStream method:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
})
ctx := context.Background()
events, err := client.GenerateContentStream(
    ctx,
    sdk.Ollama,
    "ollama/llama2",
    []sdk.Message{
        {
            Role:    sdk.System,
            Content: "You are a helpful assistant.",
        },
        {
            Role:    sdk.User,
            Content: "What is Go?",
        },
    },
)
if err != nil {
    log.Fatalf("Error generating content stream: %v", err)
}

// Read events from the stream / channel
for event := range events {
    if event.Event != nil {
        continue
    }

    switch *event.Event {
    case sdk.ContentDelta:
        if event.Data != nil {
            // Parse the streaming response
            var streamResponse sdk.CreateChatCompletionStreamResponse
            if err := json.Unmarshal(*event.Data, &streamResponse); err != nil {
                log.Printf("Error parsing stream response: %v", err)
                continue
            }

            // Process each choice in the response
            for _, choice := range streamResponse.Choices {
                // Handle reasoning content (both reasoning and reasoning_content fields)
                if choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "" {
                    fmt.Printf("💭 Reasoning: %s\n", *choice.Delta.Reasoning)
                }
                if choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "" {
                    fmt.Printf("💭 Reasoning: %s\n", *choice.Delta.ReasoningContent)
                }
                
                if choice.Delta.Content != "" {
                    // Just print the content as it comes in
                    fmt.Print(choice.Delta.Content)
                }
            }
        }

    case sdk.StreamEnd:
        // Stream has ended
        fmt.Println("\nStream ended")

    case sdk.MessageError:
        // Handle error events
        if event.Data != nil {
            var errResp struct {
                Error string `json:"error"`
            }
            if err := json.Unmarshal(*event.Data, &errResp); err != nil {
                log.Printf("Error parsing error: %v", err)
                continue
            }
            log.Printf("Error: %s", errResp.Error)
        }
    }
}
```

### Tool-Use

To use tools with the SDK, you can define a tool and provide it to the client:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
})

// Create tools array with our function
tools := []sdk.ChatCompletionTool{
    {
        Type:     sdk.Function,
        Function: sdk.FunctionObject{
            Name:        "get_current_weather",
            Description: stringPtr("Get the current weather in a given location"),
            Parameters: &sdk.FunctionParameters{
                "type": "object",
                "properties": map[string]interface{}{
                    "location": map[string]interface{}{
                        "type":        "string",
                        "enum":        []string{"san francisco", "new york", "london", "tokyo", "sydney"},
                        "description": "The city and state, e.g. San Francisco, CA",
                    },
                    "unit": map[string]interface{}{
                        "type":        "string",
                        "enum":        []string{"celsius", "fahrenheit"},
                        "description": "The temperature unit to use",
                    },
                },
                "required": []string{"location"},
            },
        }
    },
    {
        Type:     sdk.Function,
        Function: sdk.FunctionObject{
            Name:        "get_current_time",
            Description: stringPtr("Get the current time in a given location"),
            Parameters: &sdk.FunctionParameters{
                "type": "object",
                "properties": map[string]interface{}{
                    "location": map[string]interface{}{
                        "type":        "string",
                        "enum":        []string{"san francisco", "new york", "london", "tokyo", "sydney"},
                        "description": "The city and state, e.g. San Francisco, CA",
                    },
                },
                "required": []string{"location"},
            },
        }
    }
}

// Provide the tool to the client
client.WithTools(&tools).GenerateContent(ctx, provider, modelName, messages)
```

### Health Check

To check if the API is healthy:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
})

ctx := context.Background()
err := client.HealthCheck(ctx)
if err != nil {
    log.Fatalf("Health check failed: %v", err)
}
```

## Examples

For more detailed examples and use cases, check out the [examples directory](./examples/). The examples include:

-   **[Generation Example](./examples/generation/)** - Basic content generation examples
-   **[MCP List Tools Example](./examples/mcp-list-tools/)** - How to list available MCP tools
-   **[Middleware Bypass Example](./examples/middleware-bypass/)** - How to bypass middleware layers for direct provider access
-   **[Models Example](./examples/models/)** - How to list and work with different models
-   **[Stream Example](./examples/stream/)** - Streaming content generation
-   **[Stream Tools Example](./examples/stream-tools/)** - Advanced streaming with tool usage
-   **[Tools Example](./examples/tools/)** - Function calling and tool usage

Each example includes its own README with specific instructions and explanations.

## Supported Providers

The SDK supports the following LLM providers:

-   Ollama (sdk.Ollama)
-   Groq (sdk.Groq)
-   OpenAI (sdk.Openai)
-   DeepSeek (sdk.Deepseek)
-   Cloudflare (sdk.Cloudflare)
-   Cohere (sdk.Cohere)
-   Anthropic (sdk.Anthropic)
-   Google (sdk.Google)
-   Mistral AI (sdk.Mistral)

## Documentation

1. Run: `task docs`
2. Open: `http://localhost:6060/pkg/github.com/inference-gateway/sdk`

## Contributing

Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for information about how to get involved. We welcome issues, questions, and pull requests.

## License

This SDK is distributed under the MIT License, see [LICENSE](LICENSE) for more information.
