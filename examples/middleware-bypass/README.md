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
// Bypass both MCP middleware
middlewareOpts := &sdk.MiddlewareOptions{
    SkipMCP: true,
}

response, err := client.
    WithMiddlewareOptions(middlewareOpts).
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

```go
// Direct provider access (bypasses all middleware)
middlewareOpts := &sdk.MiddlewareOptions{
    DirectProvider: true,
}

response, err := client.
    WithMiddlewareOptions(middlewareOpts).
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

### 2. Using Custom Headers

Alternatively, you can use custom headers directly:

```go
response, err := client.
    WithHeader("X-MCP-Bypass", "true").
    GenerateContent(ctx, sdk.Openai, "gpt-4o", messages)
```

## Supported Headers

-   `X-MCP-Bypass: true` - Bypasses MCP middleware processing
-   `X-Direct-Provider: true` - Routes directly to provider without any middleware

## When to Use Middleware Bypass

Consider bypassing middleware when:

-   You need direct access to provider APIs without additional processing
-   You're experiencing latency issues and want to minimize processing overhead
-   You're testing or debugging and want to isolate provider responses
-   You have specific requirements that conflict with middleware behavior

## Running the Example

```bash
cd /workspaces/sdk/examples/middleware-bypass
go run main.go
```

Note: Make sure the Inference Gateway is running on `http://localhost:8080` with your desired configuration.
