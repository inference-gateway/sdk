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
		Model:  new("gpt-image-2"),
		N:      new(1),
		Size:   new(sdk.ImageSize1024x1024),
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

