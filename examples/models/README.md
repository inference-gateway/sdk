# Listing Models Example

This example demonstrates how to use the Inference Gateway SDK to list available language models from different providers.

## What This Example Shows

-   How to list all available models across providers
-   How to list models from specific providers
-   How to handle and display model information

## Running the Example

```sh
# Set your API URL (optional)
export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"

# Run the example
go run main.go
```

## Example Output

```sh
Listing all available models...
Found 15 models
1. gpt-4o (owned by openai)
2. gpt-4-turbo (owned by openai)
3. gpt-3.5-turbo (owned by openai)
4. claude-3-opus-20240229 (owned by anthropic)
5. claude-3-sonnet-20240229 (owned by anthropic)
6. llama-3.3-70b-versatile (owned by groq)
7. llama-3.3-8b-versatile (owned by groq)
8. phi3:3.8b (owned by ollama)
9. phi3:14b (owned by ollama)
10. mistral (owned by ollama)
11. llama2 (owned by ollama)
12. command-r (owned by cohere)
13. command-r-plus (owned by cohere)
14. mistral-7b (owned by cloudflare)
15. mixtral-8x7b (owned by cloudflare)

Listing models from OpenAI...
Provider: openai
1. gpt-4o
2. gpt-4-turbo
3. gpt-3.5-turbo

Listing models from Ollama...
Provider: ollama
1. phi3:3.8b
2. phi3:14b
3. mistral
4. llama2
```

## How It Works

-   The example creates a client connected to the Inference Gateway
-   It calls ListModels() to get all available models across providers
-   It then calls ListProviderModels() for specific providers like OpenAI and Ollama
-   The code processes and displays the model information in a readable format
