package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	ctx := context.Background()

	// Example 1: Regular request with MCP enabled (default)
	fmt.Println("=== Regular Request (with middleware) ===")
	response1, err := client.GenerateContent(
		ctx,
		sdk.Openai,
		"gpt-4o",
		[]sdk.Message{
			{Role: sdk.User, Content: sdk.NewMessageContent("What tools are available?")},
		},
	)
	if err != nil {
		log.Printf("Error with middleware: %v", err)
	} else {
		text, _ := response1.Choices[0].Message.Content.AsMessageContent0()
		fmt.Printf("Response: %s\n", text)
	}

	// Example 2: Bypass MCP middleware
	fmt.Println("\n=== Bypass MCP Middleware ===")
	middlewareOpts := &sdk.MiddlewareOptions{
		SkipMCP: true,
	}
	response2, err := client.
		WithMiddlewareOptions(middlewareOpts).
		GenerateContent(
			ctx,
			sdk.Openai,
			"gpt-4o",
			[]sdk.Message{
				{Role: sdk.User, Content: sdk.NewMessageContent("What tools are available?")},
			},
		)
	if err != nil {
		log.Printf("Error bypassing MCP: %v", err)
	} else {
		text, _ := response2.Choices[0].Message.Content.AsMessageContent0()
		fmt.Printf("Response: %s\n", text)
	}

	// Example 3: Direct provider access
	// DirectProvider sets X-Direct-Provider, but the gateway currently ignores
	// that header and only reads X-MCP-Bypass, so the option has no effect.
	fmt.Println("\n=== Direct Provider Access ===")
	middlewareOpts = &sdk.MiddlewareOptions{
		DirectProvider: true,
	}
	response3, err := client.
		WithMiddlewareOptions(middlewareOpts).
		GenerateContent(
			ctx,
			sdk.Openai,
			"gpt-4o",
			[]sdk.Message{
				{Role: sdk.User, Content: sdk.NewMessageContent("Simple question without middleware")},
			},
		)
	if err != nil {
		log.Printf("Error with direct provider: %v", err)
	} else {
		text, _ := response3.Choices[0].Message.Content.AsMessageContent0()
		fmt.Printf("Response: %s\n", text)
	}

	// Example 4: Using a custom header directly. Reset the middleware options
	// first so headers set by earlier examples are not left on the client.
	fmt.Println("\n=== Custom Headers Approach ===")
	response4, err := client.
		WithMiddlewareOptions(&sdk.MiddlewareOptions{}).
		WithHeader("X-MCP-Bypass", "true").
		GenerateContent(
			ctx,
			sdk.Openai,
			"gpt-4o",
			[]sdk.Message{
				{Role: sdk.User, Content: sdk.NewMessageContent("Using custom headers")},
			},
		)
	if err != nil {
		log.Printf("Error with custom headers: %v", err)
	} else {
		text, _ := response4.Choices[0].Message.Content.AsMessageContent0()
		fmt.Printf("Response: %s\n", text)
	}
}
