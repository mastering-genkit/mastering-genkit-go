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
	g := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
		genkit.WithPromptDir("prompts"),
	)

	simpleStructuredFlow := flows.NewSimpleStructuredFlow(g)
	nestedStructuredFlow := flows.NewNestedStructuredFlow(g)
	imageAnalysisFlow := flows.NewImageAnalysisFlow(g)
	audioTranscriptionFlow := flows.NewAudioTranscriptionFlow(g)
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
