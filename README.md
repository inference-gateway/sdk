# Inference Gateway Go SDK

An SDK written in Go for the [Inference Gateway](https://github.com/inference-gateway/inference-gateway).

- [Inference Gateway Go SDK](#inference-gateway-go-sdk)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Client](#creating-a-client)
    - [Listing Models](#listing-models)
    - [Generating Content](#generating-content)
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
resp, err := client.ListProviderModels(ctx, sdk.Groq)
fmt.PrintF("Provider %s", resp.Provider)
if err != nil {
    log.Fatalf("Error listing provider models: %v", err)
}
fmt.Printf("Available Groq models: %+v\n", resp.Data)
```

### Generating Content

To generate content using a model, use the GenerateContent method:

```go
client := sdk.NewClient(&sdk.ClientOptions{
    BaseURL: "http://localhost:8080/v1",
})

ctx := context.Background()
response, err := client.GenerateContent(
    ctx,
    sdk.Ollama,
    "llama2",
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
    sdk.ProviderOllama,
    "llama2",
    []sdk.Message{
        {
            Role:    sdk.MessageRoleSystem,
            Content: "You are a helpful assistant.",
        },
        {
            Role:    sdk.MessageRoleUser,
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
                Type: stringPtr("object"),
                Properties: &map[string]any{
                    "location": map[string]any{
                        "type":        "string",
                        "enum":        []string{"san francisco", "new york", "london", "tokyo", "sydney"},
                        "description": "The city and state, e.g. San Francisco, CA",
                    },
                    "unit": map[string]any{
                        "type":        "string",
                        "enum":        []string{"celsius", "fahrenheit"},
                        "description": "The temperature unit to use",
                    },
                },
                Required: &[]string{"location"},
            },
        }
    },
    {
        Type:     sdk.Function,
        Function: sdk.FunctionObject{
            Name:        "get_current_time",
            Description: stringPtr("Get the current time in a given location"),
            Parameters: &sdk.FunctionParameters{
                Type: stringPtr("object"),
                Properties: &map[string]any{
                    "location": map[string]any{
                        "type":        "string",
                        "enum":        []string{"san francisco", "new york", "london", "tokyo", "sydney"},
                        "description": "The city and state, e.g. San Francisco, CA",
                    },
                },
                Required: &[]string{"location"},
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
