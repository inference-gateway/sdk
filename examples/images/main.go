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
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: os.Getenv("INFERENCE_GATEWAY_URL"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	resp, err := client.CreateImage(ctx, sdk.Openai, sdk.CreateImageRequest{
		Prompt: "A cute cat sitting on a windowsill, digital art style",
		Model:  strPtr("dall-e-3"),
		N:      intPtr(1),
		Size:   strPtr("1024x1024"),
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, img := range resp.Data {
		if img.URL != nil {
			fmt.Println(*img.URL)
		}
	}
}

func intPtr(n int) *int  { return &n }
func strPtr(s string) *string { return &s }
