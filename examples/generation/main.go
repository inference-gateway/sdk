package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	sdk "github.com/inference-gateway/sdk"
)

func main() {
	// Get API URL and provider from environment variables or use defaults
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	providerName := os.Getenv("LLM_PROVIDER")
	if providerName == "" {
		providerName = "groq" // Default provider
	}

	modelName := os.Getenv("LLM_MODEL")
	if modelName == "" {
		modelName = "deepseek-r1-distill-llama-70b" // Default model
	}

	// Map provider string to SDK Provider type
	provider := sdk.Provider(providerName)

	// Create a new client
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: apiURL,
	})

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Define messages for the conversation
	messages := []sdk.Message{
		{
			Role:    sdk.System,
			Content: "You are a helpful assistant with expertise in Go programming.",
		},
		{
			Role:    sdk.User,
			Content: "What are the differences between goroutines and threads? Keep it short.",
		},
	}

	fmt.Printf("Generating content using %s %s...\n\n", provider, modelName)

	// Generate content
	// Note: Reasoning format is only supported by certain models like:
	// - DeepSeek R1 models (deepseek-r1, deepseek-r1-distill)
	// - OpenAI o1 models (o1-preview, o1-mini)
	// Uncomment the following lines if using a reasoning-capable model:
	//
	// reasoningFormat := "parsed"
	// response, err := client.WithOptions(&sdk.CreateChatCompletionRequest{
	//     ReasoningFormat: &reasoningFormat,
	// }).GenerateContent(ctx, provider, modelName, messages)

	response, err := client.GenerateContent(ctx, provider, modelName, messages)
	if err != nil {
		log.Printf("Error generating content with %s: %v", provider, err)

		// Suggest alternative providers if Google fails
		if provider == sdk.Google {
			log.Printf("Tip: Google's Gemini API can be slow or have connectivity issues.")
			log.Printf("Try alternative providers like:")
			log.Printf("  export LLM_PROVIDER=openai")
			log.Printf("  export LLM_PROVIDER=anthropic")
			log.Printf("  export LLM_PROVIDER=groq")
		}

		log.Fatalf("Failed to generate content")
	}

	// Print the response
	fmt.Printf("Model: %s\n", response.Model)
	if response.Choices[0].Message.Reasoning != nil {
		fmt.Printf("Reasoning: %s\n", *response.Choices[0].Message.Reasoning)
	}
	fmt.Printf("Response: %s\n", response.Choices[0].Message.Content)

	// Print usage information if available
	if response.Usage != nil {
		fmt.Printf("\nUsage Statistics:\n")
		fmt.Printf("  Prompt tokens: %d\n", response.Usage.PromptTokens)
		fmt.Printf("  Completion tokens: %d\n", response.Usage.CompletionTokens)
		fmt.Printf("  Total tokens: %d\n", response.Usage.TotalTokens)
	}
}
