package main

import (
	"context"
	"fmt"
	"log"
	"mastering-genkit-go/example/chapter-10/internal/flows"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/firebase/genkit/go/plugins/localvec"
	"github.com/firebase/genkit/go/plugins/server"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit with OpenAI plugin and default model using GPT-4o.
	g, err := genkit.Init(ctx,
		genkit.WithPlugins(
			&openai.OpenAI{
				APIKey: os.Getenv("OPENAI_API_KEY"),
			},
		),
		genkit.WithDefaultModel("openai/gpt-4o"),
	)
	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Initialize localvec plugin
	err = localvec.Init()
	if err != nil {
		log.Fatalf("could not initialize localvec: %v", err)
	}

	// Get OpenAI embedder (text-embedding-3-large) using LookupEmbedder
	embedder := genkit.LookupEmbedder(g, "openai", "text-embedding-3-large")
	if embedder == nil {
		log.Println("failed to get text-embedding-3-large embedder")
	}

	// Define retriever with localvec
	docStore, _, err := localvec.DefineRetriever(
		g,
		"arduino",
		localvec.Config{
			Embedder: embedder,
			Dir:      ".genkit/indexes", // use the .genkit/indexes directory for localvec
		},
	)

	indexerFlow := flows.NewIndexerFlow(g, []ai.ToolRef{}, docStore)

	// Get the retriever
	retriever := localvec.Retriever(g, "arduino")
	if retriever == nil {
		fmt.Println("retriever 'arduino' not found. Make sure to run the indexer first")
	}

	retrievalFlow := flows.NewRetrievalFlow(g, []ai.ToolRef{}, retriever)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /indexerFlow", genkit.Handler(indexerFlow))
	mux.HandleFunc("POST /retrievalFlow", genkit.Handler(retrievalFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
