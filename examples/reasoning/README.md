# DeepSeek Reasoner Example

This example demonstrates how to properly display the reasoning process from DeepSeek Reasoner models using the SDK.

## Features

-   **Reasoning Display**: Shows the step-by-step thinking process of DeepSeek Reasoner
-   **Streaming Support**: Displays reasoning content as it streams in real-time
-   **Visual Formatting**: Clean, readable formatting with visual separators
-   **Dual Field Support**: Handles both `reasoning` and `reasoning_content` fields

## Key Differences from Basic Examples

Unlike basic chat examples, this demonstrates:

1. **Reasoning Detection**: Checks for both `choice.Delta.Reasoning` and `choice.Delta.ReasoningContent`
2. **Visual Separation**: Clear visual distinction between reasoning and final response
3. **Real-time Streaming**: Shows reasoning as it's generated, not just the final result
4. **Enhanced Formatting**: Uses colors and borders to make reasoning more readable

## Running the Example

```bash
# Set up environment variables
export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"
export LLM_PROVIDER="deepseek"
export LLM_MODEL="deepseek-reasoner"

# Run the example
go run main.go
```

## Expected Output

```
>DeepSeek Reasoner - Reasoning Display Example
Provider: deepseek, Model: deepseek-reasoner

S Question 1: What's the best way to learn a new programming language?

> Assistant:

= Reasoning Process:
DeepSeek Reasoner Thinking
This is a great question about learning methodologies. I should consider...
1. Different learning styles (visual, hands-on, theoretical)
2. Time constraints and available resources
3. Prior programming experience level
4. Specific goals (hobby, career, specific project)
...

=Final Response:
The best approach to learning a new programming language depends on several factors...
```

## Code Explanation

### Reasoning Detection

```go
hasReasoning := (choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "") ||
    (choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "")
```

### Formatting

-   **Reasoning**: Displayed in dim gray (`\033[90m`) with visual borders
-   **Response**: Clear, normal formatting with section headers
-   **Visual Separators**: Unicode box-drawing characters for clean presentation

## Troubleshooting

If you don't see reasoning content:

1. **Check Model**: Ensure you're using `deepseek-reasoner`, not `deepseek-chat`
2. **Verify Backend**: Confirm your inference gateway supports reasoning fields
3. **Check Prompts**: Use questions that encourage step-by-step thinking
4. **Debug Fields**: Add logging to see which fields contain reasoning data

## Customization

You can customize the reasoning display by:

-   Modifying the visual formatting (colors, borders)
-   Changing the questions to focus on specific topics
-   Adding reasoning content to conversation history
-   Implementing reasoning content filtering or highlighting
