package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/firebase/genkit/go/plugins/server"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit with the Google AI plugin and Gemini 2.5 Flash.
	g, err := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
		genkit.WithPromptDir("prompts"),
	)
	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Define logging middleware
	loggingMiddleware := func(fn ai.ModelFunc) ai.ModelFunc {
		return func(ctx context.Context, req *ai.ModelRequest, cb ai.ModelStreamCallback) (*ai.ModelResponse, error) {
			start := time.Now()
			resp, err := fn(ctx, req, cb)
			if err != nil {
				log.Printf("Error after %v", time.Since(start))
				return nil, err
			}
			log.Printf("Success in %v", time.Since(start))
			return resp, nil
		}
	}

	// 1. Patterns that do not use dotprompt
	basicGenerationFlow := genkit.DefineFlow(g, "basicGenerationFlow", func(ctx context.Context, userRequest string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithPrompt(fmt.Sprintf("As a helpful cooking instructor, explain %s in simple terms that a beginner can understand.", userRequest)),
			ai.WithMiddleware(loggingMiddleware),
		)
		if err != nil {
			return "", fmt.Errorf("failed to generate response: %w", err)
		}
		return resp.Text(), nil
	})

	// 2. Patterns using dotprompt
	cookingInstructor := genkit.LookupPrompt(g, "cooking-instructor")
	if cookingInstructor == nil {
		log.Fatal("no prompt named 'cooking-instructor' found")
	}
	dotpromptFlow := genkit.DefineFlow(g, "dotpromptFlow", func(ctx context.Context, userRequest string) (string, error) {
		resp, err := cookingInstructor.Execute(ctx,
			ai.WithInput(map[string]any{
				"topic": userRequest,
			}),
			ai.WithMiddleware(loggingMiddleware),
		)
		if err != nil {
			return "", fmt.Errorf("error executing prompt: %w", err)
		}
		return resp.Text(), nil
	})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /basicGenerationFlow", genkit.Handler(basicGenerationFlow))
	mux.HandleFunc("POST /dotpromptFlow", genkit.Handler(dotpromptFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
