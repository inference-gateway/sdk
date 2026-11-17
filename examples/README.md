# Inference Gateway SDK Examples

This directory contains examples demonstrating how to use the Inference Gateway SDK in various scenarios.

## Available Examples

### [List Available Models](models/)

Shows how to list available models from different providers using the SDK.

### [List MCP Tools](mcp-list-tools/)

Demonstrates how to list available MCP (Model Context Protocol) tools via the gateway's MCP JSON-RPC endpoint (`tools/list`), when the server has `MCP_ENABLED=true` and `MCP_EXPOSE=true` configured.

### [Tokens Generation](generation/)

Demonstrates basic content generation with different LLM providers.

### [Tokens Streaming](stream/)

Illustrates how to use streaming mode to get content as it's generated.

### [Stream Tools](stream-tools/)

Demonstrates advanced streaming with tool usage and agent-like interactions.

### [Tools-Use](tools/)

Shows how to implement function calling and use tools with compatible models.

### [Decisions](decisions/)

Asks structured questions (predicate, choice, score) over shared evidence through the OpenAI Decisions API relay (`POST /v1/decisions`), when the server has `DECISIONS_ENABLED=true` configured.

### [Images](images/)

Generates and edits images with `CreateImage` and `CreateImageEdit`, when the server has `IMAGES_ENABLED=true` configured.

### [Messages](messages/)

Calls the Anthropic-compatible Messages API (sync and streaming) with `CreateMessage` and `CreateMessageStream`.

### [Middleware Bypass](middleware-bypass/)

Bypasses middleware such as MCP with `WithMiddlewareOptions` or raw headers for direct provider access.

### [Reasoning](reasoning/)

Streams reasoning output alongside the final response using `GenerateContentStream`.

### [Vision](vision/)

Sends multimodal messages with image content parts to vision-capable models.

## Running the Examples

First you need to have an Inference Gateway instance running. You can use the [Inference Gateway Docker image](ghcr.io/inference-gateway/inference-gateway) to run a local instance.

1. Copy the `.env.example` file to `.env` and set the provider API keys for the providers you plan to use:

    ```sh
    cp .env.example .env
    ```

2. Run the Inference Gateway instance:

    ```sh
    docker run --rm -it -p 8080:8080 --env-file .env ghcr.io/inference-gateway/inference-gateway:latest
    ```

3. Set the Inference Gateway API URL, to let the SDK examples know where to send requests:

    ```sh
    export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"
    ```

    The vision example reads `INFERENCE_GATEWAY_BASE_URL` (plus `INFERENCE_GATEWAY_API_KEY`) instead of `INFERENCE_GATEWAY_URL`:

    ```sh
    export INFERENCE_GATEWAY_BASE_URL="http://localhost:8080/v1"
    export INFERENCE_GATEWAY_API_KEY="your-api-key"
    ```

Most examples include their own README with specific instructions, and the rest (`decisions`, `images`, `messages` and `vision`) are documented inline in their `main.go`. The general pattern is:

1. Navigate to the example directory:

    ```sh
    cd examples/<example-name>
    ```

2. Run the example:

    ```sh
    go run .
    ```

## Prerequisites

-   Go 1.27 or later
-   Access to an Inference Gateway instance (local or remote)
-   Provider API keys configured in your Inference Gateway (for providers requiring authentication)

## Notes

-   Each example can be modified to use different providers and models by setting environment variables
