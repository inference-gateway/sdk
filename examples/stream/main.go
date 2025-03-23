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

func main() {
	// Get API URL and provider from environment variables or use defaults
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	providerName := os.Getenv("LLM_PROVIDER")
	if providerName == "" {
		providerName = "ollama" // Default provider for streaming example
	}

	modelName := os.Getenv("LLM_MODEL")
	if modelName == "" {
		modelName = "llama2" // Default model
	}

	// Map provider string to SDK Provider type
	var provider sdk.Provider
	switch providerName {
	case "openai":
		provider = sdk.Openai
	case "ollama":
		provider = sdk.Ollama
	case "groq":
		provider = sdk.Groq
	case "anthropic":
		provider = sdk.Anthropic
	case "cohere":
		provider = sdk.Cohere
	case "cloudflare":
		provider = sdk.Cloudflare
	default:
		log.Fatalf("Unsupported provider: %s", providerName)
	}

	// Create a new client
	client := sdk.NewClient(apiURL)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Define messages for the conversation
	messages := []sdk.Message{
		{
			Role:    sdk.System,
			Content: "You are a creative storyteller. Write engaging short stories.",
		},
		{
			Role:    sdk.User,
			Content: "Write a short story about a programmer who discovers an AI that can predict the future.",
		},
	}

	fmt.Printf("Streaming content from %s %s...\n\n", provider, modelName)

	// Generate content with streaming
	eventCh, err := client.GenerateContentStream(ctx, provider, modelName, messages)
	if err != nil {
		log.Fatalf("Error initiating content stream: %v", err)
	}

	// Process the stream of events
	var fullContent string
	for event := range eventCh {
		if event.Event != nil {
			switch *event.Event {
			case sdk.StreamStart:
				fmt.Println("Stream started")

			case sdk.ContentDelta:
				if event.Data != nil {
					var streamResponse sdk.CreateChatCompletionStreamResponse
					if err := json.Unmarshal(*event.Data, &streamResponse); err != nil {
						log.Printf("Error parsing stream data: %v", err)
						continue
					}

					for _, choice := range streamResponse.Choices {
						if choice.Delta.Content != "" {
							// Print the content as it arrives
							fmt.Print(choice.Delta.Content)
							fullContent += choice.Delta.Content
						}
					}
				}

			case sdk.StreamEnd:
				fmt.Println("\n\nStream ended")
			}
		}
	}

	fmt.Printf("\n\nFull content length: %d characters\n", len(fullContent))
}
