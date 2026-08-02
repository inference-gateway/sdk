// Usage:
//
//	# Image generation (existing):
//	INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
//
//	# Image edit - set IMAGE_EDIT_PATH to a PNG/JPEG file:
//	IMAGE_EDIT_PATH=/path/to/image.png \
//	  INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
//
//	# Image variation - set IMAGE_VARIATION_PATH to a PNG/JPEG file:
//	IMAGE_VARIATION_PATH=/path/to/image.png \
//	  INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
//
//	# All three:
//	IMAGE_EDIT_PATH=/path/to/edit.png \
//	  IMAGE_VARIATION_PATH=/path/to/variation.png \
//	  INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	sdk "github.com/inference-gateway/sdk"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func main() {
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: os.Getenv("INFERENCE_GATEWAY_URL"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// --- CreateImage (generation) ---
	resp, err := client.CreateImage(ctx, sdk.Openai, sdk.CreateImageRequest{
		Prompt: "A cute cat sitting on a windowsill, digital art style",
		Model:  new("gpt-image-2"),
		N:      new(1),
		Size:   new(sdk.ImageSize1024X1024),
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, img := range resp.Data {
		if img.URL != nil {
			fmt.Println("Generated:", *img.URL)
		}
	}

	// --- CreateImageEdit ---
	if path := os.Getenv("IMAGE_EDIT_PATH"); path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading edit image: %v", err)
		}
		var imageFile openapi_types.File
		imageFile.InitFromBytes(data, path)

		resp, err := client.CreateImageEdit(ctx, sdk.Openai, sdk.CreateImageEditMultipartBody{
			Image:  imageFile,
			Prompt: "Add a smiling sun in the top-right corner",
			Model:  new("gpt-image-2"),
			N:      new(1),
			Size:   new(sdk.ImageSize1024X1024),
		})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for _, img := range resp.Data {
			if img.URL != nil {
				fmt.Println("Edited:", *img.URL)
			}
		}
	}

	// --- CreateImageVariation ---
	if path := os.Getenv("IMAGE_VARIATION_PATH"); path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading variation image: %v", err)
		}
		var imageFile openapi_types.File
		imageFile.InitFromBytes(data, path)

		resp, err := client.CreateImageVariation(ctx, sdk.Openai, sdk.CreateImageVariationMultipartBody{
			Image: imageFile,
			Model: new("gpt-image-2"),
			N:     new(1),
			Size:  new(sdk.ImageSize1024X1024),
		})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for _, img := range resp.Data {
			if img.URL != nil {
				fmt.Println("Variation:", *img.URL)
			}
		}
	}
}
