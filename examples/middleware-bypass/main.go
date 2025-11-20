package main

import (
	"context"
	"fmt"
	"log"

	"github.com/inference-gateway/sdk"
)

func main() {
	// Create a new client
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: "http://localhost:8080/v1",
		APIKey:  "your-api-key",
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
		fmt.Printf("Response: %s\n", response1.Choices[0].Message.Content)
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
		fmt.Printf("Response: %s\n", response2.Choices[0].Message.Content)
	}

	// Example 3: Bypass both MCP middleware
	fmt.Println("\n=== Bypass Both MCP Middleware ===")
	middlewareOpts = &sdk.MiddlewareOptions{
		SkipMCP: true,
	}
	response3, err := client.
		WithMiddlewareOptions(middlewareOpts).
		GenerateContent(
			ctx,
			sdk.Openai,
			"gpt-4o",
			[]sdk.Message{
				{Role: sdk.User, Content: sdk.NewMessageContent("Hello, how are you?")},
			},
		)
	if err != nil {
		log.Printf("Error bypassing middleware: %v", err)
	} else {
		fmt.Printf("Response: %s\n", response3.Choices[0].Message.Content)
	}

	// Example 4: Direct provider access
	fmt.Println("\n=== Direct Provider Access ===")
	middlewareOpts = &sdk.MiddlewareOptions{
		DirectProvider: true,
	}
	response4, err := client.
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
		fmt.Printf("Response: %s\n", response4.Choices[0].Message.Content)
	}

	// Example 5: Using custom headers directly (alternative approach)
	fmt.Println("\n=== Custom Headers Approach ===")
	response5, err := client.
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
		fmt.Printf("Response: %s\n", response5.Choices[0].Message.Content)
	}
}
