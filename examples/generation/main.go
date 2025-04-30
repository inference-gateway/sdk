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
			Content: "What are the differences between goroutines and threads?",
		},
	}

	fmt.Printf("Generating content using %s %s...\n\n", provider, modelName)

	// Generate content
	reasoningFormat := "parsed"
	response, err := client.WithOptions(&sdk.CreateChatCompletionRequest{
		ReasoningFormat: &reasoningFormat,
	}).GenerateContent(ctx, provider, modelName, messages)
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
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
