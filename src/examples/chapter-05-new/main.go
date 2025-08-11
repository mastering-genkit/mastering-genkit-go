package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-05/internal/flows"
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

	// flow: 1
	simpleStructuredFlow := flows.NewSimpleStructuredFlow(g)
	// flow: 2
	nestedStructuredFlow := flows.NewNestedStructuredFlow(g)
	// flow: 3
	imageAnalysisFlow := flows.NewImageAnalysisFlow(g)
	// flow: 4
	audioTranscriptionFlow := flows.NewAudioTranscriptionFlow(g)
	// flow: 5
	characterGenerationFlow := flows.NewCharacterGenerationFlow(g)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /simpleStructuredFlow", genkit.Handler(simpleStructuredFlow))
	mux.HandleFunc("POST /nestedStructuredFlow", genkit.Handler(nestedStructuredFlow))
	mux.HandleFunc("POST /imageAnalysisFlow", genkit.Handler(imageAnalysisFlow))
	mux.HandleFunc("POST /audioTranscriptionFlow", genkit.Handler(audioTranscriptionFlow))
	mux.HandleFunc("POST /characterGenerationFlow", genkit.Handler(characterGenerationFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
