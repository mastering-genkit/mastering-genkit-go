package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"mastering-genkit-go/example/chapter-15/internal/flows"
	"mastering-genkit-go/example/chapter-15/internal/tools"

	"cloud.google.com/go/firestore"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/firebase/genkit/go/plugins/server"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit with OpenAI plugin and default model using GPT-5
	g := genkit.Init(ctx,
		genkit.WithPlugins(
			&openai.OpenAI{
				APIKey: os.Getenv("OPENAI_API_KEY"),
			},
			&googlegenai.GoogleAI{
				APIKey: os.Getenv("GEMINI_API_KEY"),
			},
		),
		genkit.WithDefaultModel("openai/gpt-5-nano"),
	)

	// Get project ID from environment (required)
	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatal("PROJECT_ID environment variable is required")
	}

	// Create Firestore client once
	firestoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer firestoreClient.Close()

	// Initialize flows with tools (passing the shared Firestore client)
	tools := []ai.ToolRef{
		tools.NewSearchRecipeDatabase(g, firestoreClient),
		tools.NewCheckIngredientStock(g, firestoreClient),
		tools.NewCalculateNutrition(g, firestoreClient),
	}
	chatFlow := flows.NewCookingBattleChatFlow(g, tools)
	actionFlow := flows.NewCookingBattleActionFlow(g, tools)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /cookingBattleChat", genkit.Handler(chatFlow))
	mux.HandleFunc("POST /cookingBattleAction", genkit.Handler(actionFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
