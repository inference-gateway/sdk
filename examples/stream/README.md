# Streaming Content Generation Example

This example demonstrates how to use the Inference Gateway SDK to stream content from LLMs. Streaming allows you to receive and display content tokens as they're generated, rather than waiting for the entire response.

## What This Example Shows

-   How to set up a streaming request with the SDK
-   How to process streaming events as they arrive
-   How to handle different event types in the stream
-   How to concatenate content parts into a complete response
-   How to access usage statistics in a streaming context

## Running the Example

```sh
# Set your API URL (optional)
export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"

# Set your preferred provider (optional, default: groq)
export LLM_PROVIDER="groq"

# Set your preferred model (optional, default: deepseek-r1-distill-llama-70b)
export LLM_MODEL="qwen-2.5-coder-32b"

# Run the example
go run main.go
```

## Example Output

```markdown
Streaming content from groq deepseek-r1-distill-llama-70b...

💭 Thinking...
Here I need to write a short story about a programmer who discovers an AI that can predict the future. I'll build the tension around the moment the program realizes the predictions are coming true...

In the quiet hours before dawn, Mara watched her terminal fill with predictions she had never asked it for...

Stream ended

Usage Statistics:
  Prompt tokens: 38
  Completion tokens: 612
  Total tokens: 650


Full content length: 612 characters
```

The thinking text and streamed story appear exactly as the model produces them. Reasoning deltas (shown under `💭 Thinking...`) appear only for models that emit reasoning, such as the default `deepseek-r1-distill-llama-70b`.

## How It Works

1. The example creates a client connected to the Inference Gateway API
2. It creates a context with a timeout to manage the request lifetime
3. It defines a conversation with system and user messages
4. It initiates a streaming request using the GenerateContentStream method
5. It processes the stream events as they arrive:

-   ContentDelta: Contains the incremental content chunks
-   StreamEnd: Marks the end of the stream

6. It prints content incrementally as it's received (reasoning deltas arrive inside ContentDelta events and are shown under a `💭 Thinking...` marker with dim gray text)
7. At the end, it displays the total content length and usage statistics

## Customization

You can modify this example to:

-   Change the system prompt to give the LLM different instructions
-   Change the user prompt to ask different questions
-   Use different providers and models through environment variables
-   Process the streamed content in different ways (e.g., updating a UI)
-   Track token usage for budgeting or rate limiting purposes
