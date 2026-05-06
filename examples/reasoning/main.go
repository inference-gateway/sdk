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

// isReasoningModel checks if the given model supports reasoning capabilities
func isReasoningModel(provider sdk.Provider, model string) bool {
	reasoningModels := map[sdk.Provider][]string{
		sdk.Deepseek: {"deepseek-v4-pro"},
		// Add more providers and their reasoning models here as they become available
		// Example: sdk.Openai: {"o1", "o1-preview", "o1-mini"},
	}

	if models, exists := reasoningModels[provider]; exists {
		for _, reasoningModel := range models {
			if model == reasoningModel {
				return true
			}
		}
	}
	return false
}

func main() {
	// Environment setup - specifically configured for DeepSeek Reasoner
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	// Default to DeepSeek Reasoner
	providerName := os.Getenv("LLM_PROVIDER")
	if providerName == "" {
		providerName = "deepseek"
	}

	modelName := os.Getenv("LLM_MODEL")
	if modelName == "" {
		modelName = "deepseek-v4-pro"
	}

	provider := sdk.Provider(providerName)
	client := sdk.NewClient(&sdk.ClientOptions{BaseURL: apiURL})
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// Validate that the selected model supports reasoning
	if !isReasoningModel(provider, modelName) {
		log.Fatalf("Error: Model '%s' from provider '%s' does not support reasoning capabilities. Please use a reasoning model like 'deepseek-v4-pro'.", modelName, provider)
	}

	// Conversation setup with prompts that encourage step-by-step reasoning
	conversationHistory := []sdk.Message{
		{
			Role: sdk.System,
			Content: sdk.NewMessageContent(`You are a helpful assistant. When answering questions, think through your response step by step.
Consider multiple approaches and explain your reasoning process clearly.`),
		},
	}

	// Example questions that encourage reasoning
	questions := []string{
		"What's the best way to learn a new programming language? Consider different learning styles and time constraints.",
		"Explain why quicksort has O(n log n) average time complexity. Walk me through the mathematical reasoning.",
		"If I have a 1000-piece jigsaw puzzle and I can place 2 pieces per minute on average, but I slow down as the puzzle gets more complex, how should I estimate the total time needed?",
		"Compare the trade-offs between microservices and monolithic architecture. What factors should influence this decision?",
	}

	fmt.Printf("🧠 DeepSeek Reasoner - Reasoning Display Example\n")
	fmt.Printf("Provider: %s, Model: %s\n", provider, modelName)
	fmt.Printf("═══════════════════════════════════════════════════════════════════════════════\n")

	// Process each question
	for i, question := range questions {
		fmt.Printf("❓ Question %d: %s\n", i+1, question)

		// Add user message to conversation
		conversationHistory = append(conversationHistory, sdk.Message{
			Role:    sdk.User,
			Content: sdk.NewMessageContent(question),
		})

		// Generate response with streaming
		eventCh, err := client.GenerateContentStream(ctx, provider, modelName, conversationHistory)
		if err != nil {
			log.Fatalf("Error initiating content stream: %v", err)
		}

		// Process stream with enhanced reasoning display
		assistantMessage := sdk.Message{Role: sdk.Assistant}
		var assistantContent strings.Builder
		var isThinking bool
		var reasoningBuffer strings.Builder

		fmt.Printf("🤖 Assistant: ")

		for event := range eventCh {
			if event.Event != nil && *event.Event == sdk.ContentDelta && event.Data != nil {
				var streamResponse sdk.CreateChatCompletionStreamResponse
				if err := json.Unmarshal(*event.Data, &streamResponse); err != nil {
					continue
				}

				for _, choice := range streamResponse.Choices {
					// Enhanced reasoning handling for DeepSeek Reasoner
					hasReasoning := (choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "") ||
						(choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "")

					if hasReasoning {
						if !isThinking {
							isThinking = true
							fmt.Printf("\n💭 Reasoning: ")
						}

						// Handle both reasoning and reasoning_content fields
						var reasoningText string
						if choice.Delta.Reasoning != nil && *choice.Delta.Reasoning != "" {
							reasoningText = *choice.Delta.Reasoning
						}
						if choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "" {
							reasoningText = *choice.Delta.ReasoningContent
						}

						// Display and buffer reasoning with nice formatting
						if reasoningText != "" {
							reasoningBuffer.WriteString(reasoningText)
							// Display reasoning text directly as it streams
							fmt.Printf("\033[90m%s\033[0m", reasoningText)
						}
					}

					// Handle content (the final response)
					if choice.Delta.Content != "" {
						if isThinking {
							isThinking = false
							fmt.Printf("\n\n📝 Response: ")
						}
						fmt.Print(choice.Delta.Content)
						assistantContent.WriteString(choice.Delta.Content)
					}
				}
			}
		}

		// Add assistant response to conversation history
		assistantMessage.Content = sdk.NewMessageContent(assistantContent.String())
		conversationHistory = append(conversationHistory, assistantMessage)

		fmt.Print("\n" + strings.Repeat("═", 79) + "\n")

		// Small delay between questions for readability
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("✅ DeepSeek Reasoner reasoning demonstration complete!\n")
	fmt.Printf("💡 Key Features Demonstrated:\n")
	fmt.Printf("   • Step-by-step reasoning display with visual formatting\n")
	fmt.Printf("   • Proper handling of both 'reasoning' and 'reasoning_content' fields\n")
	fmt.Printf("   • Clear separation between thinking process and final response\n")
	fmt.Printf("   • Streaming reasoning content as it's generated\n")
}
