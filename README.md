# Inference Gateway Go SDK

An SDK written in Go for the [Inference Gateway](https://github.com/inference-gateway/inference-gateway).

- [Inference Gateway Go SDK](#inference-gateway-go-sdk)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Client](#creating-a-client)
    - [Using Custom Headers](#using-custom-headers)
    - [Listing Models](#listing-models)
    - [Listing MCP Tools](#listing-mcp-tools)
    - [Generating Content](#generating-content)
    - [Using ReasoningFormat](#using-reasoningformat)
    - [Streaming Content](#streaming-content)
    - [Tool-Use](#tool-use)
    - [Health Check](#health-check)
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

## Supported Providers

The SDK supports the following LLM providers:

-   Ollama (sdk.Ollama)
-   Groq (sdk.Groq)
-   OpenAI (sdk.Openai)
-   DeepSeek (sdk.Deepseek)
-   Cloudflare (sdk.Cloudflare)
-   Cohere (sdk.Cohere)
-   Anthropic (sdk.Anthropic)

## Documentation

1. Run: `task docs`
2. Open: `http://localhost:6060/pkg/github.com/inference-gateway/sdk`

## Contributing

Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for information about how to get involved. We welcome issues, questions, and pull requests.

## License

This SDK is distributed under the MIT License, see [LICENSE](LICENSE) for more information.
