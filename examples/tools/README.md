# Tool-Use

This example demonstrates how to use the Inference Gateway SDK to implement Tool-Use calling with LLMs. The example creates a weather assistant that can call a function to get weather information for a specific location.

## What This Example Shows

-   How to define a function that an LLM can call
-   How to structure function parameters
-   How to handle function call responses
-   How to send function results back to the LLM
-   How to process the final assistant response

## Running the Example

```sh
# Set your API URL (optional)
export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"

# Run the example
go run [main.go](main.go)
```

## Example Output:

```sh
Asking about weather with function calling using gpt-4o...

Model is calling a function:
Function: get_current_weather
Arguments: {"location":"San Francisco"}

Sending function result back to the model...

Final response: Based on the current weather data for San Francisco, it's 14°C (celsius) and foggy.
```

## How It Works

1. A tool called get_current_weather is defined with parameters for location and unit
2. The function is provided to the LLM as a tool it can use
3. When the user asks about weather, the LLM, based on the context, decides to call the tool function
4. The example code extracts the function call arguments and calls the real function
5. The function result is sent back to the LLM as a tool message
6. The LLM uses the information to provide a natural language response

## Note

This example requires a provider that supports function calling, such as OpenAI's GPT-4 models.
