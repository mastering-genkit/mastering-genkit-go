package main

import (
	"context"
	"examples/chapter-04/internal/flows"
	"log"
	"net/http"
	"os"

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

	// Create flows
	basicGenerationFlow := flows.NewBasicGenerationFlow(g)
	dotpromptFlow := flows.NewDotpromptFlow(g)
	imageAnalysisFlow := flows.NewImageAnalysisFlow(g)
	audioTranscriptionFlow := flows.NewAudioTranscriptionFlow(g)
	imageGenerationFlow := flows.NewImageGenerationFlow(g)

	// Set up HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("POST /basicGenerationFlow", genkit.Handler(basicGenerationFlow))
	mux.HandleFunc("POST /dotpromptFlow", genkit.Handler(dotpromptFlow))
	mux.HandleFunc("POST /imageAnalysisFlow", genkit.Handler(imageAnalysisFlow))
	mux.HandleFunc("POST /audioTranscriptionFlow", genkit.Handler(audioTranscriptionFlow))
	mux.HandleFunc("POST /imageGenerationFlow", genkit.Handler(imageGenerationFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
