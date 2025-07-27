package main

import (
	"context"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit
	g, err := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithPromptDir("prompts"),
	)
	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Look up the cooking instructor prompt
	cookingInstructor := genkit.LookupPrompt(g, "cooking-instructor")
	if cookingInstructor == nil {
		log.Fatal("no prompt named 'cooking-instructor' found")
	}

	// Execute the prompt with a cooking topic
	resp, err := cookingInstructor.Execute(ctx,
		ai.WithInput(map[string]any{
			"topic": "how to properly dice onions",
		}))
	if err != nil {
		log.Printf("Error executing prompt: %v", err)
	} else {
		log.Printf("Explanation: %s", resp.Text())
	}

	// Keep the application running for Developer UI
	// This allows you to test prompts visually without defining flows
	select {}
}
