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
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	providerName := os.Getenv("LLM_PROVIDER")
	if providerName == "" {
		providerName = "anthropic" // The Messages API is Anthropic-compatible
	}

	modelName := os.Getenv("LLM_MODEL")
	if modelName == "" {
		modelName = "claude-sonnet-5"
	}

	provider := sdk.Provider(providerName)

	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: apiURL,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var content sdk.MessagesMessage_Content
	if err := content.FromMessagesMessageContent0("What are the differences between goroutines and threads? Keep it short."); err != nil {
		log.Fatalf("Failed to build message content: %v", err)
	}

	request := sdk.CreateMessagesRequest{
		Model:     modelName,
		MaxTokens: 1024,
		Messages: []sdk.MessagesMessage{
			{Role: sdk.MessagesMessageRoleUser, Content: content},
		},
	}

	// --- Synchronous message ---
	fmt.Printf("Creating a message using %s %s...\n\n", provider, modelName)

	response, err := client.CreateMessage(ctx, provider, request)
	if err != nil {
		log.Fatalf("Failed to create message: %v", err)
	}

	fmt.Printf("Model: %s\n", response.Model)
	for _, block := range response.Content {
		if text, err := block.AsMessagesTextBlock(); err == nil {
			fmt.Printf("Response: %s\n", text.Text)
		}
	}
	fmt.Printf("Stop reason: %s\n", response.StopReason)
	fmt.Printf("Usage: %d input tokens, %d output tokens\n\n",
		response.Usage.InputTokens, response.Usage.OutputTokens)

	// --- Streaming message ---
	fmt.Println("Now streaming the same message...")

	events, err := client.CreateMessageStream(ctx, provider, request)
	if err != nil {
		log.Fatalf("Failed to create message stream: %v", err)
	}

	for event := range events {
		if event.Event == nil {
			continue
		}

		switch *event.Event {
		case sdk.ContentDelta:
			if event.Data == nil {
				continue
			}
			var streamEvent sdk.MessagesStreamEvent
			if err := json.Unmarshal(*event.Data, &streamEvent); err != nil {
				log.Printf("Error parsing stream event: %v", err)
				continue
			}
			if streamEvent.Delta != nil && streamEvent.Delta.Text != nil {
				fmt.Print(*streamEvent.Delta.Text)
			}

		case sdk.StreamEnd:
			fmt.Println("\nStream ended")
		}
	}
}
