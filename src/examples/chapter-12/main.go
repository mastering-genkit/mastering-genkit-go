package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-12/internal/agent"
	"mastering-genkit-go/example/chapter-12/internal/cli"

	"github.com/firebase/genkit/go/ai/exp/localstore"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

func main() {
	ctx := context.Background()

	// Create the Ollama plugin pointing to the local Ollama server
	ollamaPlugin := &ollama.Ollama{
		ServerAddress: "http://localhost:11434",
		Timeout:       300, // local models can take a while to load
	}

	// Initialize Genkit with the Ollama plugin. Agents are part of Genkit's
	// experimental API, so we have to opt in with WithExperimental.
	g := genkit.Init(ctx,
		genkit.WithPlugins(ollamaPlugin),
		genkit.WithExperimental(),
	)

	// Keep session snapshots (the conversation state) in memory
	store := localstore.NewInMemorySessionStore[any]()

	// Define the agent. Ollama resolves any locally installed model by name.
	assistant := agent.NewAssistantAgent(g, "ollama/gemma4:e4b", store)
	log.Println("Genkit initialized with agent:", assistant.Name())

	// Start the interactive CLI
	if err := cli.New(assistant).Run(ctx); err != nil {
		log.Fatalf("AI agent error: %v", err)
	}
}
