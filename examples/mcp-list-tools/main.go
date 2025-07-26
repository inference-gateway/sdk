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
	// Get API URL from environment or use default
	apiURL := os.Getenv("INFERENCE_GATEWAY_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/v1"
	}

	// Get API key from environment (optional)
	apiKey := os.Getenv("INFERENCE_GATEWAY_API_KEY")

	// Create a new client
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: apiURL,
		APIKey:  apiKey,
	})

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// List available MCP tools
	fmt.Println("Listing available MCP tools...")
	tools, err := client.ListTools(ctx)
	if err != nil {
		log.Fatalf("Error listing tools: %v", err)
	}

	fmt.Printf("Found %d MCP tools:\n\n", len(tools.Data))

	// Display each tool
	for i, tool := range tools.Data {
		fmt.Printf("Tool %d:\n", i+1)
		fmt.Printf("  Name: %s\n", tool.Name)
		fmt.Printf("  Description: %s\n", tool.Description)
		fmt.Printf("  Server: %s\n", tool.Server)

		if tool.InputSchema != nil {
			fmt.Printf("  Input Schema: %+v\n", *tool.InputSchema)
		}
		fmt.Println()
	}
}
