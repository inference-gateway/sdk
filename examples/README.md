# Inference Gateway SDK Examples

This directory contains examples demonstrating how to use the Inference Gateway SDK in various scenarios.

## Available Examples

### [Models](models/)

Shows how to list available models from different providers using the SDK.

### [Generation](generation/)

Demonstrates basic content generation with different LLM providers.

### [Streaming](stream/)

Illustrates how to use streaming mode to get content as it's generated.

### [Tools & Function Calling](tools/)

Shows how to implement function calling and use tools with compatible models.

## Running the Examples

First you need to have an Inference Gateway instance running. You can use the [Inference Gateway Docker image](ghcr.io/inference-gateway/inference-gateway) to run a local instance.

1. Copy the `.env.example` file to `.env` and set the Inference Gateway API URL:

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

Each example directory contains a README.md with specific instructions, but the general pattern is:

1. Navigate to the example directory:

    ```sh
    cd examples/<example-name>
    ```

2. Run the example:
    ```sh
    go run main.go
    ```

## Prerequisites

-   Go 1.23 or later
-   Access to an Inference Gateway instance (local or remote)
-   Provider API keys configured in your Inference Gateway (for providers requiring authentication)

## Notes

-   Some examples may require specific LLM providers that support certain features (like function calling)
-   Each example can be modified to use different providers and models by setting environment variables
