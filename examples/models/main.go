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
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: apiURL,
	})

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// List all models from all providers
	fmt.Println("Listing all available models...")
	allModels, err := client.ListModels(ctx)
	if err != nil {
		log.Fatalf("Error listing models: %v", err)
	}

	fmt.Printf("Found %d models\n", len(allModels.Data))
	for i, model := range allModels.Data {
		fmt.Printf("%d. %s (owned by %s)\n", i+1, model.Id, model.OwnedBy)
	}

	// List models for a specific provider
	fmt.Println("\nListing models from Groq Cloud...")
	groqModels, err := client.ListProviderModels(ctx, sdk.Groq)
	if err != nil {
		log.Printf("Error listing Groq Cloud models: %v", err)
	} else {
		fmt.Printf("Provider: %s\n", *groqModels.Provider)
		for i, model := range groqModels.Data {
			fmt.Printf("%d. %s\n", i+1, model.Id)
		}
	}

	// List models for a specific provider
	fmt.Println("\nListing models from DeepSeek...")
	deepseekModels, err := client.ListProviderModels(ctx, sdk.Deepseek)
	if err != nil {
		log.Printf("Error listing DeepSeek models: %v", err)
	} else {
		fmt.Printf("Provider: %s\n", *deepseekModels.Provider)
		for i, model := range deepseekModels.Data {
			fmt.Printf("%d. %s\n", i+1, model.Id)
		}
	}

	// List models for another provider
	fmt.Println("\nListing models from Ollama...")
	ollamaModels, err := client.ListProviderModels(ctx, sdk.Ollama)
	if err != nil {
		log.Printf("Error listing Ollama models: %v", err)
	} else {
		fmt.Printf("Provider: %s\n", *ollamaModels.Provider)
		for i, model := range ollamaModels.Data {
			fmt.Printf("%d. %s\n", i+1, model.Id)
		}
	}
}
