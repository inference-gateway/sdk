// Usage:
//
//	INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
//
//	# Optional first-frame image for avatar models:
//	VIDEO_INPUT_REFERENCE_PATH=/path/to/frame.png \
//	  INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
//
// Requires a gateway running with VIDEOS_ENABLED=true and a provider that
// serves the Videos API (currently elevenlabs and mirelo).
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"

	sdk "github.com/inference-gateway/sdk"
)

func main() {
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: os.Getenv("INFERENCE_GATEWAY_URL"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	request := sdk.CreateVideoRequest{
		Model:   "elevenlabs/creatify-aurora",
		Prompt:  new("A person waving enthusiastically at the camera"),
		Seconds: new("4"),
		Size:    new("720x1280"),
	}
	if path := os.Getenv("VIDEO_INPUT_REFERENCE_PATH"); path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading input reference: %v", err)
		}
		var image openapi_types.File
		image.InitFromBytes(data, path)
		request.InputReference = &image
	}

	job, err := client.CreateVideo(ctx, sdk.Elevenlabs, request)
	if err != nil {
		log.Fatalf("Error creating video job: %v", err)
	}
	fmt.Printf("Job created: id=%s status=%s\n", job.ID, job.Status)

	for {
		job, err = client.RetrieveVideo(ctx, sdk.Elevenlabs, job.ID)
		if err != nil {
			log.Fatalf("Error polling video job: %v", err)
		}
		if job.Progress != nil {
			fmt.Printf("Polled: id=%s status=%s progress=%d%%\n", job.ID, job.Status, *job.Progress)
		} else {
			fmt.Printf("Polled: id=%s status=%s\n", job.ID, job.Status)
		}
		if job.Status == sdk.VideoJobStatusCompleted {
			break
		}
		if job.Status == sdk.VideoJobStatusFailed {
			log.Fatalf("Video job failed: %s", jobErrorMessage(job))
		}
		time.Sleep(20 * time.Second)
	}

	video, err := client.DownloadVideoContent(ctx, sdk.Elevenlabs, job.ID)
	if err != nil {
		log.Fatalf("Error downloading video: %v", err)
	}
	if err := os.WriteFile("output.mp4", video, 0o644); err != nil {
		log.Fatalf("Error saving video: %v", err)
	}
	fmt.Printf("Saved %d bytes to output.mp4\n", len(video))
}

func jobErrorMessage(job *sdk.VideoJob) string {
	if job.Error == nil || job.Error.Message == nil {
		return "unknown error"
	}
	return *job.Error.Message
}
