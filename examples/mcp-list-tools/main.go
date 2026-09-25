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

	// List available MCP tools through the gateway's MCP JSON-RPC endpoint
	fmt.Println("Listing available MCP tools...")
	request, err := sdk.NewMCPJSONRPCRequest(1, sdk.ToolsList, nil)
	if err != nil {
		log.Fatalf("Error building request: %v", err)
	}

	response, err := client.MCPJSONRPC(ctx, request)
	if err != nil {
		log.Fatalf("Error listing tools: %v", err)
	}

	if response.Error != nil {
		log.Fatalf("JSON-RPC error %d: %s", response.Error.Code, response.Error.Message)
	}

	// The result is the MCP spec's ListToolsResult - decode the bits we display
	var result struct {
		Tools []struct {
			Name        string         `json:"name"`
			Description string         `json:"description"`
			InputSchema map[string]any `json:"inputSchema"`
		} `json:"tools"`
	}

	raw, err := json.Marshal(response.Result)
	if err != nil {
		log.Fatalf("Error encoding result: %v", err)
	}
	if err := json.Unmarshal(raw, &result); err != nil {
		log.Fatalf("Error decoding result: %v", err)
	}

	fmt.Printf("Found %d MCP tools:\n\n", len(result.Tools))

	// Display each tool
	for i, tool := range result.Tools {
		fmt.Printf("Tool %d:\n", i+1)
		fmt.Printf("  Name: %s\n", tool.Name)
		fmt.Printf("  Description: %s\n", tool.Description)

		if tool.InputSchema != nil {
			fmt.Printf("  Input Schema: %+v\n", tool.InputSchema)
		}
		fmt.Println()
	}
}
