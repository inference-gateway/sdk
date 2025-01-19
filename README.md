# Inference Gateway Go SDK

An SDK written in Go for the [Inference Gateway](https://github.com/edenreich/inference-gateway).

- [Inference Gateway Go SDK](#inference-gateway-go-sdk)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Client](#creating-a-client)
    - [Listing Models](#listing-models)
    - [Generating Content](#generating-content)
  - [License](#license)

## Installation

To install the SDK, use `go get`:

```sh
go get github.com/edenreich/inference-gateway-go-sdk
```

## Usage

### Creating a Client

To create a client, use the `NewClient` function:

```go
package main

import (
    "fmt"
    "log"

    "github.com/edenreich/inference-gateway-go-sdk"
)

func main() {
    client := sdk.NewClient("http://localhost:8080")

    // List models
    models, err := client.ListModels()
    if err != nil {
        log.Fatalf("Error listing models: %v", err)
    }
    fmt.Println("Available models:", models)

    // Generate content
    response, err := client.GenerateContent("providerName", "modelName", "your prompt here")
    if err != nil {
        log.Fatalf("Error generating content: %v", err)
    }
    fmt.Println("Generated content:", response.Response.Content)
}
```

### Listing Models

To list available models, use the ListModels method:

```go
models, err := client.ListModels()
if err != nil {
    log.Fatalf("Error listing models: %v", err)
}
fmt.Println("Available models:", models)
```

### Generating Content

To generate content using a model, use the GenerateContent method:

```go
response, err := client.GenerateContent("providerName", "modelName", "your prompt here")
if err != nil {
    log.Fatalf("Error generating content: %v", err)
}
fmt.Println("Generated content:", response.Response.Content)
```

## License

This SDK is distributed under the MIT License, see [LICENSE](LICENSE) for more information.
