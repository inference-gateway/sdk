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
		providerName = "groq" // Default provider for streaming example
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Define messages for the conversation
	messages := []sdk.Message{
		{
			Role:    sdk.System,
			Content: sdk.NewMessageContent("You are a creative storyteller. Write engaging short stories."),
		},
		{
			Role:    sdk.User,
			Content: sdk.NewMessageContent("Write a short story about a programmer who discovers an AI that can predict the future."),
		},
	}

	fmt.Printf("Streaming content from %s %s...\n\n", provider, modelName)

	// Generate content with streaming
	eventCh, err := client.GenerateContentStream(ctx, provider, modelName, messages)
	if err != nil {
		log.Fatalf("Error initiating content stream: %v", err)
	}

	// Process the stream of events
	var usageStats *sdk.CompletionUsage
	var fullContent string
	// Track thinking animation state
	var isThinking bool
	var reasoningText string

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
						// Handle reasoning (both reasoning and reasoning_content fields)
						hasReasoning := (choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "") ||
							(choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "")

						if hasReasoning {
							if !isThinking {
								// Start thinking animation if we weren't already thinking
								isThinking = true
								fmt.Printf("\n\033[33m💭 Thinking...\033[0m\n")
								reasoningText = ""
							}

							// Store and display the reasoning content with special formatting
							// Handle both reasoning and reasoning_content fields
							if choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "" {
								reasoningText += *choice.Delta.Reasoning
								fmt.Printf("\033[90m%s\033[0m", *choice.Delta.Reasoning)
							}
							if choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "" {
								reasoningText += *choice.Delta.ReasoningContent
								fmt.Printf("\033[90m%s\033[0m", *choice.Delta.ReasoningContent)
							}

						} else if choice.Delta.Content != "" {
							if isThinking {
								isThinking = false
								fmt.Printf("\n\n")
							}

							// Display the content as it arrives
							fmt.Print(choice.Delta.Content)
							fullContent += choice.Delta.Content
						}
					}

					// Extract usage statistics if present
					if streamResponse.Usage != nil {
						usageStats = streamResponse.Usage
					}
				}

			case sdk.StreamEnd:
				fmt.Println("\n\nStream ended")
			}
		}
	}

	// Display usage statistics if available
	if usageStats != nil {
		fmt.Printf("\nUsage Statistics:\n")
		fmt.Printf("  Prompt tokens: %d\n", usageStats.PromptTokens)
		fmt.Printf("  Completion tokens: %d\n", usageStats.CompletionTokens)
		fmt.Printf("  Total tokens: %d\n", usageStats.TotalTokens)
	}

	fmt.Printf("\n\nFull content length: %d characters\n", len(fullContent))
}
