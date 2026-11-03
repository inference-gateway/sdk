// Usage:
//
//	INFERENCE_GATEWAY_URL=http://localhost:8080/v1 go run main.go
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

	var input sdk.CreateDecisionRequest_Input
	if err := input.FromCreateDecisionRequestInput0("Buy now!!!"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	decision, err := client.CreateDecision(ctx, sdk.Openai, sdk.CreateDecisionRequest{
		Model: "openai/gpt-6-luna",
		Input: input,
		Questions: []sdk.DecisionQuestion{{
			Name:         "is_spam",
			Type:         sdk.Predicate,
			Instructions: "Is this message unsolicited spam?",
		}},
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if decision.Answers != nil {
		for _, answer := range *decision.Answers {
			fmt.Printf("%v: %v\n", answer["name"], answer["answer"])
		}
	}
}
