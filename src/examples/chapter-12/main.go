package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-12/internal/cli"
	"mastering-genkit-go/example/chapter-12/internal/flows"

	"github.com/firebase/genkit/go/plugins/ollama"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

func main() {
	ctx := context.Background()

	// Create the Ollama instance and define the model
	ollamaPlugin := &ollama.Ollama{
		ServerAddress: "http://localhost:11434",
	}

	// Initialize Genkit with Ollama plugin first
	g := genkit.Init(ctx,
		genkit.WithPlugins(ollamaPlugin),
	)

	model := ollamaPlugin.DefineModel(g,
		ollama.ModelDefinition{
			Name: "gemma3n:e4b",
			Type: "chat", // "chat" or "generate"
		},
		&ai.ModelOptions{
			Supports: &ai.ModelSupports{
				Media:      false,
				Multiturn:  true,
				Tools:      true,
				Context:    true,
				ToolChoice: true,
			},
		},
	)

	log.Println("Genkit initialized with Ollama plugin and model:", model.Name())

	// Create the chat flow with tools (empty for now)
	chatFlow := flows.NewChatFlow(g, []ai.ToolRef{}, ai.NewModelRef(model.Name(), nil))

	// Create and start the AI agent
	agent := cli.NewAgent(chatFlow)

	log.Println("Starting AI Agent...")
	if err := agent.Run(ctx); err != nil {
		log.Fatalf("AI agent error: %v", err)
	}
}
