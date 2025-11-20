package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	sdk "github.com/inference-gateway/sdk"
)

// WeatherData represents the structure for the weather function return value
type WeatherData struct {
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Description string  `json:"description"`
}

// CalculatorResult represents the result of a calculation
type CalculatorResult struct {
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}

// GetWeatherInfo simulates a weather API call
func GetWeatherInfo(location string) WeatherData {
	weatherData := map[string]WeatherData{
		"san francisco": {Temperature: 14, Unit: "celsius", Description: "Foggy"},
		"new york":      {Temperature: 22, Unit: "celsius", Description: "Sunny"},
		"london":        {Temperature: 10, Unit: "celsius", Description: "Rainy"},
		"tokyo":         {Temperature: 28, Unit: "celsius", Description: "Clear"},
		"sydney":        {Temperature: 20, Unit: "celsius", Description: "Partly cloudy"},
	}

	result := WeatherData{Temperature: 15, Unit: "celsius", Description: "Unknown"}
	if data, exists := weatherData[strings.ToLower(location)]; exists {
		result = data
	}
	return result
}

// Calculate performs basic arithmetic operations
func Calculate(operation string, a, b float64) CalculatorResult {
	var result float64
	switch operation {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		if b != 0 {
			result = a / b
		} else {
			result = 0
		}
	default:
		result = 0
	}

	return CalculatorResult{
		Operation: fmt.Sprintf("%.2f %s %.2f", a, operation, b),
		Result:    result,
	}
}

// executeToolCall executes a tool call and returns the result as JSON string
func executeToolCall(functionName, arguments string) string {
	switch functionName {
	case "get_current_weather":
		var args struct {
			Location string `json:"location"`
			Unit     string `json:"unit"`
		}
		if err := json.Unmarshal([]byte(arguments), &args); err != nil {
			return fmt.Sprintf(`{"error": "Failed to parse arguments: %v"}`, err)
		}

		weather := GetWeatherInfo(args.Location)
		result, _ := json.Marshal(weather)
		return string(result)

	case "calculate":
		var args struct {
			A         float64 `json:"a"`
			B         float64 `json:"b"`
			Operation string  `json:"operation"`
		}
		if err := json.Unmarshal([]byte(arguments), &args); err != nil {
			return fmt.Sprintf(`{"error": "Failed to parse arguments: %v"}`, err)
		}

		calcResult := Calculate(args.Operation, args.A, args.B)
		result, _ := json.Marshal(calcResult)
		return string(result)

	default:
		return fmt.Sprintf(`{"error": "Unknown function: %s"}`, functionName)
	}
}

func main() {
	// Environment setup
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	providerName := os.Getenv("LLM_PROVIDER")
	if providerName == "" {
		providerName = "deepseek"
	}

	modelName := os.Getenv("LLM_MODEL")
	if modelName == "" {
		modelName = "deepseek-chat"
	}

	provider := sdk.Provider(providerName)
	client := sdk.NewClient(&sdk.ClientOptions{BaseURL: apiURL})
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// Define tools
	weatherFunction := sdk.FunctionObject{
		Name:        "get_current_weather",
		Description: stringPtr("Get the current weather in a given location"),
		Parameters: &sdk.FunctionParameters{
			"type": "object",
			"properties": map[string]interface{}{
				"location": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"san francisco", "new york", "london", "tokyo", "sydney"},
					"description": "The city and state, e.g. San Francisco, CA",
				},
				"unit": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"celsius", "fahrenheit"},
					"description": "The temperature unit to use",
				},
			},
			"required": []string{"location"},
		},
	}

	calculatorFunction := sdk.FunctionObject{
		Name:        "calculate",
		Description: stringPtr("Perform basic arithmetic operations"),
		Parameters: &sdk.FunctionParameters{
			"type": "object",
			"properties": map[string]interface{}{
				"a": map[string]interface{}{
					"type":        "number",
					"description": "First number",
				},
				"b": map[string]interface{}{
					"type":        "number",
					"description": "Second number",
				},
				"operation": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"add", "subtract", "multiply", "divide"},
					"description": "The arithmetic operation to perform",
				},
			},
			"required": []string{"a", "b", "operation"},
		},
	}

	tools := []sdk.ChatCompletionTool{
		{Type: sdk.Function, Function: weatherFunction},
		{Type: sdk.Function, Function: calculatorFunction},
	}

	// Agent conversation setup
	conversationHistory := []sdk.Message{
		{
			Role:    sdk.System,
			Content: sdk.NewMessageContent("You are a helpful assistant with access to weather information and a calculator. Always use proper JSON format for function calls."),
		},
	}

	// Simulate a queue of hard-coded user's inputs / requests
	// Normally you would let the user type in the input, but for simplicity and deterministic example we simulate it
	messageQueue := []sdk.Message{
		{Role: sdk.User, Content: sdk.NewMessageContent("What's the weather in San Francisco and what's 15 multiplied by 7?")},
		{Role: sdk.User, Content: sdk.NewMessageContent("Now calculate 100 divided by 5")},
	}

	fmt.Printf("🤖 Starting agent conversation with %s %s...\n\n", provider, modelName)

	// Agent loop
	for len(messageQueue) > 0 {
		// Pop next message from queue
		nextMessage := messageQueue[0]
		messageQueue = messageQueue[1:]
		conversationHistory = append(conversationHistory, nextMessage)

		fmt.Printf("👤 User: %s\n\n", nextMessage.Content)

		// Keep processing until no more tool calls
		for {
			// Generate response with streaming
			eventCh, err := client.WithTools(&tools).GenerateContentStream(ctx, provider, modelName, conversationHistory)
			if err != nil {
				log.Fatalf("Error initiating content stream: %v", err)
			}

			// Process stream
			assistantMessage := sdk.Message{Role: sdk.Assistant}
			var isThinking bool
			toolCallBuffer := make(map[string]*sdk.ChatCompletionMessageToolCall)
			toolCallsExecuted := false

			fmt.Printf("🤖 Assistant: ")

			for event := range eventCh {
				if event.Event != nil && *event.Event == sdk.ContentDelta && event.Data != nil {
					var streamResponse sdk.CreateChatCompletionStreamResponse
					if err := json.Unmarshal(*event.Data, &streamResponse); err != nil {
						continue
					}

					for _, choice := range streamResponse.Choices {
						// Handle reasoning (both reasoning and reasoning_content fields)
						hasReasoning := (choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "") ||
							(choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "")

						if hasReasoning {
							if !isThinking {
								isThinking = true
								fmt.Printf("\n💭 Thinking...\n")
							}
							// Handle both reasoning and reasoning_content fields
							if choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "" {
								fmt.Printf("\033[90m%s\033[0m", *choice.Delta.Reasoning)
							}
							if choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "" {
								fmt.Printf("\033[90m%s\033[0m", *choice.Delta.ReasoningContent)
							}
						}

						// Handle content
						if choice.Delta.Content != "" {
							if isThinking {
								isThinking = false
								fmt.Printf("\n\n")
							}
							fmt.Print(choice.Delta.Content)
							assistantMessage.Content += choice.Delta.Content
						} // Handle tool calls - accumulate and execute when complete
						for _, toolCallChunk := range choice.Delta.ToolCalls {
							// Use index for consistent tracking across chunks
							id := fmt.Sprintf("tool_call_%d", toolCallChunk.Index)

							if toolCallBuffer[id] == nil {
								toolCallBuffer[id] = &sdk.ChatCompletionMessageToolCall{
									Id:   toolCallChunk.ID, // Keep original ID if available
									Type: sdk.ChatCompletionToolType(toolCallChunk.Type),
									Function: sdk.ChatCompletionMessageToolCallFunction{
										Name:      "",
										Arguments: "",
									},
								}
							}

							// Update ID if we get a new one
							if toolCallChunk.ID != "" {
								toolCallBuffer[id].Id = toolCallChunk.ID
							}

							// Accumulate function name and arguments
							if toolCallChunk.Function.Name != "" {
								toolCallBuffer[id].Function.Name += toolCallChunk.Function.Name
							}
							if toolCallChunk.Function.Arguments != "" {
								toolCallBuffer[id].Function.Arguments += toolCallChunk.Function.Arguments
							}

							// Check if JSON is complete and execute
							args := strings.TrimSpace(toolCallBuffer[id].Function.Arguments)
							funcName := strings.TrimSpace(toolCallBuffer[id].Function.Name)
							if args != "" && funcName != "" && strings.HasSuffix(args, "}") {
								var temp interface{}
								if json.Unmarshal([]byte(args), &temp) == nil {
									// Execute tool call immediately
									toolCall := toolCallBuffer[id]
									fmt.Printf("\n🔧 Executing: %s(%s)\n", toolCall.Function.Name, toolCall.Function.Arguments)

									result := executeToolCall(toolCall.Function.Name, toolCall.Function.Arguments)
									fmt.Printf("📋 Result: %s\n", result)

									// Add to conversation
									if assistantMessage.ToolCalls == nil {
										assistantMessage.ToolCalls = &[]sdk.ChatCompletionMessageToolCall{}
									}
									*assistantMessage.ToolCalls = append(*assistantMessage.ToolCalls, *toolCall)

									conversationHistory = append(conversationHistory, sdk.Message{
										Role:      sdk.Assistant,
										Content:   assistantMessage.Content,
										ToolCalls: assistantMessage.ToolCalls,
									})
									conversationHistory = append(conversationHistory, sdk.Message{
										Role:       sdk.Tool,
										Content:    result,
										ToolCallId: &toolCall.Id,
									})

									assistantMessage = sdk.Message{Role: sdk.Assistant}
									toolCallsExecuted = true

									// Remove this tool call from buffer to avoid re-execution
									delete(toolCallBuffer, id)
								}
							}
						}
					}
				}
			}

			fmt.Printf("\n")

			// Execute any remaining tool calls that weren't executed during streaming
			for toolCallId, toolCall := range toolCallBuffer {
				args := strings.TrimSpace(toolCall.Function.Arguments)
				funcName := strings.TrimSpace(toolCall.Function.Name)
				if args != "" && funcName != "" && strings.HasSuffix(args, "}") {
					var temp interface{}
					if json.Unmarshal([]byte(args), &temp) == nil {
						fmt.Printf("\n🔧 Executing: %s(%s)\n", toolCall.Function.Name, toolCall.Function.Arguments)

						result := executeToolCall(toolCall.Function.Name, toolCall.Function.Arguments)
						fmt.Printf("📋 Result: %s\n", result)

						// Add to conversation
						if assistantMessage.ToolCalls == nil {
							assistantMessage.ToolCalls = &[]sdk.ChatCompletionMessageToolCall{}
						}
						*assistantMessage.ToolCalls = append(*assistantMessage.ToolCalls, *toolCall)

						conversationHistory = append(conversationHistory, sdk.Message{
							Role:      sdk.Assistant,
							Content:   assistantMessage.Content,
							ToolCalls: assistantMessage.ToolCalls,
						})
						conversationHistory = append(conversationHistory, sdk.Message{
							Role:       sdk.Tool,
							Content:    result,
							ToolCallId: &toolCall.Id,
						})

						toolCallsExecuted = true
						delete(toolCallBuffer, toolCallId)
					}
				}
			}

			// If no tool calls were executed, add final message and break
			if !toolCallsExecuted {
				if assistantMessage.Content != "" {
					conversationHistory = append(conversationHistory, assistantMessage)
				}
				break
			}
		}

		fmt.Print("\n" + strings.Repeat("-", 50) + "\n\n")
	}

	fmt.Printf("✅ Agent conversation complete!\n")
}

func stringPtr(s string) *string {
	return &s
}
