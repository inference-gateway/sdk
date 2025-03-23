package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	sdk "github.com/inference-gateway/sdk"
)

// WeatherData represents the structure for the weather function return value
type WeatherData struct {
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Description string  `json:"description"`
}

// GetWeatherInfo simulates a weather API call
func GetWeatherInfo(location string) WeatherData {
	// This would normally call a real weather API
	// For this example, we're just returning mock data
	weatherData := map[string]WeatherData{
		"san francisco": {Temperature: 14, Unit: "celsius", Description: "Foggy"},
		"new york":      {Temperature: 22, Unit: "celsius", Description: "Sunny"},
		"london":        {Temperature: 10, Unit: "celsius", Description: "Rainy"},
		"tokyo":         {Temperature: 28, Unit: "celsius", Description: "Clear"},
		"sydney":        {Temperature: 20, Unit: "celsius", Description: "Partly cloudy"},
	}

	// Default weather if location not found
	result := WeatherData{Temperature: 15, Unit: "celsius", Description: "Unknown"}

	// Check for the location in our mock data
	if data, exists := weatherData[location]; exists {
		result = data
	}

	return result
}

func main() {
	// Get API URL and provider from environment variables or use defaults
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	// OpenAI is best for function calling examples
	provider := sdk.Openai
	modelName := "gpt-4o"

	// Create a new client
	client := sdk.NewClient(apiURL)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Define the function we want the model to potentially call
	weatherFunction := sdk.FunctionObject{
		Name:        "get_current_weather",
		Description: stringPtr("Get the current weather in a given location"),
		Parameters: &sdk.FunctionParameters{
			Type: stringPtr("object"),
			Properties: &map[string]interface{}{
				"location": map[string]interface{}{
					"type":        "string",
					"description": "The city and state, e.g. San Francisco, CA",
				},
				"unit": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"celsius", "fahrenheit"},
					"description": "The temperature unit to use",
				},
			},
			Required: &[]string{"location"},
		},
	}

	// Define messages for the conversation
	messages := []sdk.Message{
		{
			Role:    sdk.System,
			Content: "You are a helpful weather assistant.",
		},
		{
			Role:    sdk.User,
			Content: "What's the weather like in San Francisco right now?",
		},
	}

	// Create tools array with our function
	tools := []sdk.ChatCompletionTool{
		{
			Type:     sdk.Function,
			Function: weatherFunction,
		},
	}

	// Create request with tool specification
	requestBody := sdk.CreateChatCompletionRequest{
		Model:    modelName,
		Messages: messages,
		Tools:    &tools,
	}

	// Make the request directly since we need to customize it
	fmt.Printf("Asking about weather with function calling using %s...\n\n", modelName)

	// Use custom HTTP client for this specific example
	httpClient := client.(*sdk.ClientImpl).GetHTTPClient()
	resp, err := httpClient.R().
		SetContext(ctx).
		SetQueryParam("provider", string(provider)).
		SetBody(requestBody).
		SetResult(&sdk.CreateChatCompletionResponse{}).
		Post(fmt.Sprintf("%s/v1/chat/completions", apiURL))

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	if resp.IsError() {
		log.Fatalf("API error: %s", resp.Body())
	}

	// Parse the response
	result, ok := resp.Result().(*sdk.CreateChatCompletionResponse)
	if !ok || result == nil {
		log.Fatalf("Failed to parse response")
	}

	// Check if the model wants to call our function
	if len(result.Choices) > 0 && result.Choices[0].Message.ToolCalls != nil && len(*result.Choices[0].Message.ToolCalls) > 0 {
		fmt.Println("Model is calling a function:")

		// Extract the function call
		toolCall := (*result.Choices[0].Message.ToolCalls)[0]
		fmt.Printf("Function: %s\n", toolCall.Function.Name)
		fmt.Printf("Arguments: %s\n\n", toolCall.Function.Arguments)

		// Parse arguments
		var args struct {
			Location string `json:"location"`
			Unit     string `json:"unit"`
		}
		if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
			log.Fatalf("Failed to parse arguments: %v", err)
		}

		// Call the function with the provided arguments
		location := args.Location
		if location == "" {
			location = "san francisco" // Default
		}

		weather := GetWeatherInfo(location)
		weatherJSON, _ := json.Marshal(weather)

		// Create a message to send back with the function result
		messages = append(messages, result.Choices[0].Message)
		messages = append(messages, sdk.Message{
			Role:       sdk.Tool,
			Content:    string(weatherJSON),
			ToolCallId: &toolCall.Id,
		})

		// Make another request to get the final response
		fmt.Println("Sending function result back to the model...\n")
		response, err := client.GenerateContent(ctx, provider, modelName, messages)
		if err != nil {
			log.Fatalf("Error generating content: %v", err)
		}

		// Print the final response
		fmt.Printf("Final response: %s\n", response.Choices[0].Message.Content)
	} else {
		// Model responded directly without calling function
		fmt.Printf("Model response: %s\n", result.Choices[0].Message.Content)
	}
}

func stringPtr(s string) *string {
	return &s
}
