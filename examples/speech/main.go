// Usage:
//
//	INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
//
// Requires a gateway running with AUDIO_ENABLED=true and an OpenAI key
// (OPENAI_API_KEY). The Audio API is also served by elevenlabs and
// mirelo.
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	audio, err := client.CreateSpeech(ctx, sdk.Openai, sdk.CreateSpeechRequest{
		Model:          "gpt-4o-mini-tts",
		Input:          "The Inference Gateway lets you generate speech, sound effects and music through a single API.",
		Voice:          "alloy",
		ResponseFormat: new(sdk.CreateSpeechRequestResponseFormatMp3),
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if err := os.WriteFile("output.mp3", audio, 0644); err != nil {
		log.Fatalf("Error writing audio: %v", err)
	}
	fmt.Printf("Saved %d bytes to output.mp3\n", len(audio))
}
