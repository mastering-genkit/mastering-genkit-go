package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-14/internal/flows"
	"mastering-genkit-go/example/chapter-14/internal/handlers"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/plugins/server"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit
	g, err := genkit.Init(ctx)

	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Create the simple flow with tools (empty for now)
	simpleFlow := flows.NewSimpleFlow(g, []ai.ToolRef{})

	log.Println("Setting up and starting server...")

	mux := http.NewServeMux()

	// Health and readiness endpoints
	mux.HandleFunc("GET /health", handlers.HealthHandler)
	mux.HandleFunc("GET /ready", handlers.ReadyHandler)

	// Main application endpoint
	mux.HandleFunc("POST /simpleFlow", genkit.Handler(simpleFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
