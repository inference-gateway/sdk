# Middleware Bypass Example

This example demonstrates how to bypass MCP (Model Context Protocol) middleware when making requests through the Inference Gateway SDK.

## Overview

The Inference Gateway supports several types of middleware that can process requests before they reach the underlying AI models. Sometimes you may want to bypass these middlewares for direct access to the providers.

## Middleware Types

-   **MCP Middleware**: Handles Model Context Protocol integration

## Bypass Methods

### 1. Using MiddlewareOptions (Recommended)

The SDK provides a convenient `MiddlewareOptions` struct to control middleware behavior:

```go
// Bypass MCP middleware only
middlewareOpts := &sdk.MiddlewareOptions{
    SkipMCP: true,
}

response, err := client.
    WithMiddlewareOptions(middlewareOpts).
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

```go
// Direct provider access (no effect: the gateway ignores X-Direct-Provider)
middlewareOpts := &sdk.MiddlewareOptions{
    DirectProvider: true,
}

response, err := client.
    WithMiddlewareOptions(middlewareOpts).
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

### 2. Using Custom Headers

Alternatively, you can use the MCP bypass header directly:

```go
response, err := client.
    WithHeader("X-MCP-Bypass", "true").
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

## Supported Headers

-   `X-MCP-Bypass: true` - Bypasses MCP middleware processing
-   `X-Direct-Provider: true` - Intended to route directly to the provider without middleware, but the gateway currently ignores it, so it has no effect

## When to Use Middleware Bypass

Consider bypassing middleware when:

-   You need direct access to provider APIs without additional processing
-   You're experiencing latency issues and want to minimize processing overhead
-   You're testing or debugging and want to isolate provider responses
-   You have specific requirements that conflict with middleware behavior

## Running the Example

```bash
cd examples/middleware-bypass
go run main.go
```

Note: Set `INFERENCE_GATEWAY_URL` to point the example at your gateway (defaults to `http://localhost:8080/v1`) and `INFERENCE_GATEWAY_API_KEY` if your gateway requires a key.
