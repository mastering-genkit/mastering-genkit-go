package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"mastering-genkit-go/example/chapter-07/internal/flows"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/firebase/genkit/go/plugins/server"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit with OpenAI plugin and default model using GPT-5
	g := genkit.Init(ctx,
		genkit.WithPlugins(
			&openai.OpenAI{
				APIKey: os.Getenv("OPENAI_API_KEY"),
			},
		),
		genkit.WithDefaultModel("openai/gpt-5"),
	)

	// Define the streaming flow
	recipeStepsFlow := flows.NewRecipeFlow(g)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /recipeStepsFlow", genkit.Handler(recipeStepsFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
