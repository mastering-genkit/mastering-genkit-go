package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-13/cloud/internal/flows"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/plugins/googlecloud"
	"github.com/firebase/genkit/go/plugins/server"
	bedrock "github.com/xavidop/genkit-aws-bedrock-go"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit
	g, err := genkit.Init(ctx,
		genkit.WithPlugins(
			&bedrock.Bedrock{
				Region:                "us-east-1",
				DefineCommonModels:    true, // Automatically define common models
				DefineCommonEmbedders: true, // Automatically define common embedders
			},
			&googlecloud.GoogleCloud{
				ProjectID:   "my-project-id", // Replace with your Google Cloud project ID
				ForceExport: true,
			},
		),
		genkit.WithDefaultModel("bedrock/anthropic.claude-3-haiku-20240307-v1:0"), // Set default model
	)

	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Create the chat flow with tools (empty for now)
	chatFlow := flows.NewChatFlow(g, []ai.ToolRef{})

	log.Println("Setting up and starting server...")

	mux := http.NewServeMux()
	mux.HandleFunc("POST /chatFlow", genkit.Handler(chatFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
