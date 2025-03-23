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
	// Get API URL from environment variable or use default
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	// Create a new client
	client := sdk.NewClient(apiURL)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// List all models from all providers
	fmt.Println("Listing all available models...")
	allModels, err := client.ListModels(ctx)
	if err != nil {
		log.Fatalf("Error listing models: %v", err)
	}

	fmt.Printf("Found %d models\n", len(*allModels.Data))
	for i, model := range *allModels.Data {
		fmt.Printf("%d. %s (owned by %s)\n", i+1, *model.Id, *model.OwnedBy)
	}

	// List models for a specific provider
	fmt.Println("\nListing models from OpenAI...")
	openaiModels, err := client.ListProviderModels(ctx, sdk.Openai)
	if err != nil {
		log.Printf("Error listing OpenAI models: %v", err)
	} else {
		fmt.Printf("Provider: %s\n", *openaiModels.Provider)
		for i, model := range *openaiModels.Data {
			fmt.Printf("%d. %s\n", i+1, *model.Id)
		}
	}

	// List models for another provider
	fmt.Println("\nListing models from Ollama...")
	ollamaModels, err := client.ListProviderModels(ctx, sdk.Ollama)
	if err != nil {
		log.Printf("Error listing Ollama models: %v", err)
	} else {
		fmt.Printf("Provider: %s\n", *ollamaModels.Provider)
		for i, model := range *ollamaModels.Data {
			fmt.Printf("%d. %s\n", i+1, *model.Id)
		}
	}
}
