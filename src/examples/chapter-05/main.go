package main

import (
	"context"
	"examples/chapter-05/internal/flows"
	"examples/chapter-05/internal/tools"
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

	// Initialize Genkit with Google AI plugin and gemini-2.5-pro model
	g, err := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-pro"),
	)
	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Create flows
	extractInvoiceDataFlow := flows.NewExtractInvoiceDataFlow(g)

	processCompleteOrderFlow := flows.NewProcessCompleteOrderFlow(g, []ai.ToolRef{
		tools.NewValidateCustomer(g),
		tools.NewCheckInventory(g),
		tools.NewCreateOrder(g),
	})

	// Set up HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("POST /extractInvoiceData", genkit.Handler(extractInvoiceDataFlow))
	mux.HandleFunc("POST /processCompleteOrder", genkit.Handler(processCompleteOrderFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
