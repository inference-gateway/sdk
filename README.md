# Inference Gateway Go SDK

An SDK written in Go for the [Inference Gateway](https://github.com/inference-gateway/inference-gateway).

- [Inference Gateway Go SDK](#inference-gateway-go-sdk)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Client](#creating-a-client)
    - [Listing Models](#listing-models)
    - [Generating Content](#generating-content)
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
    client := sdk.NewClient("http://localhost:8080")

    // Health check
    if err := client.HealthCheck(); err != nil {
        log.Fatalf("Health check failed: %v", err)
    }

    // List models
    models, err := client.ListModels()
    if err != nil {
        log.Fatalf("Error listing models: %v", err)
    }
    fmt.Printf("Available models: %+v\n", models)

    // Generate content using the llama2 model
    response, err := client.GenerateContent(
        sdk.ProviderOllama,
        "llama2",
        []sdk.Message{
            {
                Role:    sdk.RoleSystem,
                Content: "You are a helpful assistant.",
            },
            {
                Role:    sdk.RoleUser,
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
// List all models from all providers
models, err := client.ListModels()
if err != nil {
    log.Fatalf("Error listing models: %v", err)
}
fmt.Printf("All available models: %+v\n", models)

// List models for a specific provider
providerModels, err := client.ListProviderModels(sdk.ProviderGroq)
if err != nil {
    log.Fatalf("Error listing provider models: %v", err)
}
fmt.Printf("Available Groq models: %+v\n", providerModels)
```

### Generating Content

To generate content using a model, use the GenerateContent method:

```go
response, err := client.GenerateContent(
    sdk.ProviderOllama,
    "llama2",
    []sdk.Message{
        {
            Role:    sdk.RoleSystem,
            Content: "You are a helpful assistant.",
        },
        {
            Role:    sdk.RoleUser,
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

## Supported Providers

The SDK supports the following LLM providers:

-   Ollama (sdk.ProviderOllama)
-   Groq (sdk.ProviderGroq)
-   OpenAI (sdk.ProviderOpenAI)
-   Google (sdk.ProviderGoogle)
-   Cloudflare (sdk.ProviderCloudflare)
-   Cohere (sdk.ProviderCohere)
-   Anthropic (sdk.ProviderAnthropic

## Documentation

1. Run: `task docs`
2. Open: `http://localhost:6060/pkg/github.com/inference-gateway/sdk`

## Contributing

Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for information about how to get involved. We welcome issues, questions, and pull requests.

## License

This SDK is distributed under the MIT License, see [LICENSE](LICENSE) for more information.
