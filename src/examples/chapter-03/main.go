package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/firebase/genkit/go/plugins/server"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit with the Google AI plugin and Gemini 2.5 Flash.
	g := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	// Define a greeting flow that accepts a user's name and generates a personalized greeting
	greetingFlow := genkit.DefineFlow(g, "greetingFlow", func(ctx context.Context, userRequest string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithSystem("You are a friendly assistant that creates warm, personalized greetings. When given a name, respond with a welcoming message that makes the person feel valued."),
			ai.WithPrompt(fmt.Sprintf("Create a greeting for %s", userRequest)),
		)
		if err != nil {
			return "", fmt.Errorf("failed to generate response: %w", err)
		}

		return resp.Text(), nil
	})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /greetingFlow", genkit.Handler(greetingFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
