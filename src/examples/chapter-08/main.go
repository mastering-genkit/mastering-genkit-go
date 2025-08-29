package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-08/internal/flows"
	"mastering-genkit-go/example/chapter-08/internal/tools"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/firebase/genkit/go/plugins/server"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit with OpenAI plugin and default model using GPT-4o.
	g := genkit.Init(ctx,
		genkit.WithPlugins(&openai.OpenAI{
			APIKey: os.Getenv("OPENAI_API_KEY"),
		}),
		genkit.WithDefaultModel("openai/gpt-4o"),
	)

	operatingSystemFlow := flows.NewOperatingSystemFlow(g, []ai.ToolRef{
		tools.NewListDirectories(g),
		tools.NewCreateDirectory(g),
		tools.NewGetCurrentDate(g),
		tools.NewSystemInfo(g),
	})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /operatingSystemFlow", genkit.Handler(operatingSystemFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
