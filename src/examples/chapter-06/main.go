package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"examples/chapter-06/internal/flows"

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
		genkit.WithPromptDir("prompts"),
	)

	// Create flow orchestrator
	recommendationFlow := flows.NewRecommendationFlow(g,
		flows.NewAnalyzeGenreFlow(g),
		flows.NewAcousticInstrumentFlow(g),
		flows.NewElectronicInstrumentFlow(g),
	)

	// Setup HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("POST /recommendationFlow", genkit.Handler(recommendationFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
