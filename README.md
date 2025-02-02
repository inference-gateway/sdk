# Inference Gateway Go SDK

An SDK written in Go for the [Inference Gateway](https://github.com/inference-gateway/inference-gateway).

- [Inference Gateway Go SDK](#inference-gateway-go-sdk)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Client](#creating-a-client)
    - [Listing Models](#listing-models)
    - [Generating Content](#generating-content)
    - [Health Check](#health-check)
    - [Streaming Content](#streaming-content)
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
    client := sdk.NewClient("http://localhost:8080")
    ctx := context.Background()

    // Health check
    if err := client.HealthCheck(ctx); err != nil {
        log.Fatalf("Health check failed: %v", err)
    }

    // List models
    models, err := client.ListModels(ctx)
    if err != nil {
        log.Fatalf("Error listing models: %v", err)
    }
    fmt.Printf("Available models: %+v\n", models)

    // Generate content using the llama2 model
    response, err := client.GenerateContent(
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
        log.Fatalf("Error generating content: %v", err)
    }

    fmt.Printf("Generated content: %s\n", response.Response.Content)
}
```

### Listing Models

To list available models, use the ListModels method:

```go
ctx := context.Background()

// List all models from all providers
models, err := client.ListModels(ctx)
if err != nil {
    log.Fatalf("Error listing models: %v", err)
}
fmt.Printf("All available models: %+v\n", models)

// List models for a specific provider
providerModels, err := client.ListProviderModels(ctx, sdk.ProviderGroq)
if err != nil {
    log.Fatalf("Error listing provider models: %v", err)
}
fmt.Printf("Available Groq models: %+v\n", providerModels)
```

### Generating Content

To generate content using a model, use the GenerateContent method:

```go
ctx := context.Background()
response, err := client.GenerateContent(
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
    }
)
if err != nil {
    log.Fatalf("Error generating content: %v", err)
}
fmt.Println("Generated content:", response.Response.Content)
```

### Health Check

To check if the API is healthy:

```go
err := client.HealthCheck()
if err != nil {
    log.Fatalf("Health check failed: %v", err)
}
```

### Streaming Content

To generate content using streaming mode, use the GenerateContentStream method:

```go
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
    switch event.Event {
    case sdk.StreamEventContentDelta:
        // Option 1: Use anonymous struct for simple cases
        var delta struct {
            Content string `json:"content"`
        }
        if err := json.Unmarshal(event.Data, &delta); err != nil {
            log.Printf("Error parsing delta: %v", err)
            continue
        }
        fmt.Print(delta.Content)

        // Option 2: Use GenerateResponseTokens for full response structure
        var tokens sdk.GenerateResponseTokens
        if err := json.Unmarshal(event.Data, &tokens); err != nil {
            log.Printf("Error parsing tokens: %v", err)
            continue
        }
        fmt.Printf("Model: %s, Role: %s, Content: %s\n",
            tokens.Model, tokens.Role, tokens.Content)

    case sdk.StreamEventMessageError:
        var errResp struct {
            Error string `json:"error"`
        }
        if err := json.Unmarshal(event.Data, &errResp); err != nil {
            log.Printf("Error parsing error: %v", err)
            continue
        }
        log.Printf("Error: %s", errResp.Error)
    }
}
```

## Supported Providers

The SDK supports the following LLM providers:

-   Ollama (sdk.ProviderOllama)
-   Groq (sdk.ProviderGroq)
-   OpenAI (sdk.ProviderOpenAI)
-   Cloudflare (sdk.ProviderCloudflare)
-   Cohere (sdk.ProviderCohere)
-   Anthropic (sdk.ProviderAnthropic)

## Documentation

1. Run: `task docs`
2. Open: `http://localhost:6060/pkg/github.com/inference-gateway/sdk`

## Contributing

Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for information about how to get involved. We welcome issues, questions, and pull requests.

## License

This SDK is distributed under the MIT License, see [LICENSE](LICENSE) for more information.
