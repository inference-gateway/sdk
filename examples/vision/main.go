package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/inference-gateway/sdk"
)

func main() {
	// Create a new SDK client
	client := sdk.NewClient(&sdk.ClientOptions{
		BaseURL: os.Getenv("INFERENCE_GATEWAY_BASE_URL"),
		APIKey:  os.Getenv("INFERENCE_GATEWAY_API_KEY"),
	})

	// Example 1: Simple text message (backward compatible)
	fmt.Println("=== Example 1: Simple Text Message ===")
	textMessage, err := sdk.NewTextMessage(sdk.User, "What is Go programming language?")
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.GenerateContent(
		context.Background(),
		sdk.Openai,
		"gpt-4o",
		[]sdk.Message{textMessage},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Extract text from response
	responseText, err := response.Choices[0].Message.Content.AsMessageContent0()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n\n", responseText)

	// Example 2: Vision message with text and image URL
	fmt.Println("=== Example 2: Vision Message with Image URL ===")

	// Create content parts
	var contentParts []sdk.ContentPart

	// Add text part
	textPart, err := sdk.NewTextContentPart("What is in this image?")
	if err != nil {
		log.Fatal(err)
	}
	contentParts = append(contentParts, textPart)

	// Add image part with auto detail level
	imagePart, err := sdk.NewImageContentPart(
		"https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png",
		nil, // auto detail level (default)
	)
	if err != nil {
		log.Fatal(err)
	}
	contentParts = append(contentParts, imagePart)

	// Create vision message
	visionMessage, err := sdk.NewImageMessage(sdk.User, contentParts)
	if err != nil {
		log.Fatal(err)
	}

	response, err = client.GenerateContent(
		context.Background(),
		sdk.Openai,
		"gpt-4o", // or gpt-4o-mini, gpt-4-turbo, etc.
		[]sdk.Message{visionMessage},
	)
	if err != nil {
		log.Fatal(err)
	}

	responseText, err = response.Choices[0].Message.Content.AsMessageContent0()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n\n", responseText)

	// Example 3: Vision message with base64 encoded image
	fmt.Println("=== Example 3: Vision Message with Base64 Image ===")

	// Create content parts with base64 image
	contentParts = []sdk.ContentPart{}

	textPart, err = sdk.NewTextContentPart("Describe this image in detail")
	if err != nil {
		log.Fatal(err)
	}
	contentParts = append(contentParts, textPart)

	// Use high detail level for better analysis
	highDetail := sdk.High
	imagePart, err = sdk.NewImageContentPart(
		"data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD...", // truncated for example
		&highDetail,
	)
	if err != nil {
		log.Fatal(err)
	}
	contentParts = append(contentParts, imagePart)

	visionMessage, err = sdk.NewImageMessage(sdk.User, contentParts)
	if err != nil {
		log.Fatal(err)
	}

	response, err = client.GenerateContent(
		context.Background(),
		sdk.Openai,
		"gpt-4o",
		[]sdk.Message{visionMessage},
	)
	if err != nil {
		log.Fatal(err)
	}

	responseText, err = response.Choices[0].Message.Content.AsMessageContent0()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n\n", responseText)

	// Example 4: Multiple images in one message
	fmt.Println("=== Example 4: Multiple Images in One Message ===")

	contentParts = []sdk.ContentPart{}

	textPart, err = sdk.NewTextContentPart("Compare these two images and describe the differences:")
	if err != nil {
		log.Fatal(err)
	}
	contentParts = append(contentParts, textPart)

	// First image
	image1, err := sdk.NewImageContentPart("https://example.com/image1.jpg", nil)
	if err != nil {
		log.Fatal(err)
	}
	contentParts = append(contentParts, image1)

	// Second image
	image2, err := sdk.NewImageContentPart("https://example.com/image2.jpg", nil)
	if err != nil {
		log.Fatal(err)
	}
	contentParts = append(contentParts, image2)

	visionMessage, err = sdk.NewImageMessage(sdk.User, contentParts)
	if err != nil {
		log.Fatal(err)
	}

	response, err = client.GenerateContent(
		context.Background(),
		sdk.Openai,
		"gpt-4o",
		[]sdk.Message{visionMessage},
	)
	if err != nil {
		log.Fatal(err)
	}

	responseText, err = response.Choices[0].Message.Content.AsMessageContent0()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n\n", responseText)
}
