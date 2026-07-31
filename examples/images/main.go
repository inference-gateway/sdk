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
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	providerName := os.Getenv("LLM_PROVIDER")
	if providerName == "" {
		providerName = "openai"
	}

	modelName := os.Getenv("LLM_MODEL")
	if modelName == "" {
		modelName = "dall-e-3"
	}

	prompt := os.Getenv("IMAGE_PROMPT")
	if prompt == "" {
		prompt = "A cute cat sitting on a windowsill, digital art style"
	}

	provider := sdk.Provider(providerName)

	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: apiURL,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Printf("Generating image using %s %s...\n\n", provider, modelName)
	fmt.Printf("Prompt: %s\n\n", prompt)

	response, err := client.CreateImage(ctx, provider, sdk.CreateImageRequest{
		Prompt: prompt,
		Model:  &modelName,
		N:      intPtr(1),
		Size:   strPtr("1024x1024"),
	})
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}

	for i, img := range response.Data {
		fmt.Printf("Image %d:\n", i+1)
		if img.URL != nil {
			fmt.Printf("  URL: %s\n", *img.URL)
		}
		if img.B64Json != nil {
			fmt.Printf("  B64 JSON: %s\n", truncate(*img.B64Json, 80))
		}
		if img.RevisedPrompt != nil {
			fmt.Printf("  Revised prompt: %s\n", *img.RevisedPrompt)
		}
	}

	if response.Usage != nil {
		fmt.Printf("\nUsage:\n")
		if response.Usage.InputTokens != nil {
			fmt.Printf("  Input tokens: %d\n", *response.Usage.InputTokens)
		}
		if response.Usage.OutputTokens != nil {
			fmt.Printf("  Output tokens: %d\n", *response.Usage.OutputTokens)
		}
		if response.Usage.TotalTokens != nil {
			fmt.Printf("  Total tokens: %d\n", *response.Usage.TotalTokens)
		}
	}
}

func intPtr(n int) *int { return &n }

func strPtr(s string) *string { return &s }

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
