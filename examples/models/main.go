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

	// List all models from all providers
	fmt.Println("Listing all available models...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	allModels, err := client.ListModels(ctx)
	cancel() // Clean up resources
	if err != nil {
		log.Fatalf("Error listing models: %v", err)
	}

	fmt.Printf("Found %d models\n", len(allModels.Data))
	for i, model := range allModels.Data {
		fmt.Printf("%d. %s (owned by %s)\n", i+1, model.Id, model.OwnedBy)
	}

	// List models for specific providers
	providers := []struct {
		name     string
		provider sdk.Provider
	}{
		{"Groq", sdk.Groq},
		{"DeepSeek", sdk.Deepseek},
		{"Google", sdk.Google},
		{"Ollama", sdk.Ollama},
		{"Ollama Cloud", sdk.OllamaCloud},
	}

	for _, p := range providers {
		fmt.Printf("\nListing models from %s...\n", p.name)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		models, err := client.ListProviderModels(ctx, p.provider)
		cancel()

		if err != nil {
			log.Printf("Error listing %s models: %v", p.name, err)
			continue
		}

		fmt.Printf("Provider: %s\n", *models.Provider)
		for i, model := range models.Data {
			fmt.Printf("%d. %s\n", i+1, model.Id)
		}
	}
}
