package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-11/internal/evals"
	"mastering-genkit-go/example/chapter-11/internal/flows"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/plugins/compat_oai/anthropic"
	"github.com/firebase/genkit/go/plugins/evaluators"
	"github.com/firebase/genkit/go/plugins/server"
	"github.com/openai/openai-go/option"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit
	metrics := []evaluators.MetricConfig{
		{
			MetricType: evaluators.EvaluatorDeepEqual,
		},
	}

	g, err := genkit.Init(ctx,
		genkit.WithPlugins(
			&anthropic.Anthropic{Opts: []option.RequestOption{
				option.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
			},
			},
			&evaluators.GenkitEval{
				Metrics: metrics,
			},
		),
		genkit.WithDefaultModel("anthropic/claude-3-7-sonnet-20250219"),
	)

	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Create the chat flow with tools (empty for now)
	chatFlow := flows.NewChatFlow(g, []ai.ToolRef{})

	// Run evaluator examples
	runNonBilledEvaluator(ctx, g)
	runBilledEvaluator(ctx, g, chatFlow)
	runGenkitBuiltinEvaluator(ctx, g)

	log.Println("Setting up and starting server...")

	mux := http.NewServeMux()
	mux.HandleFunc("POST /chatFlow", genkit.Handler(chatFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}

// runNonBilledEvaluator demonstrates a non-billed custom evaluator
func runNonBilledEvaluator(ctx context.Context, g *genkit.Genkit) {
	log.Println("Running non-billed custom evaluator...")

	nonBilledEvaluatorService := evals.NewNonBilledEvaluatorService(g)

	nonBilledEvaluator, err := nonBilledEvaluatorService.NewResponseQualityEvaluator()
	if err != nil {
		log.Fatalf("could not define evaluator: %v", err)
	}

	log.Println("Custom evaluator defined:", nonBilledEvaluator.Name())

	dataset := nonBilledEvaluatorService.GetResponseQualityEvaluatorDataset()

	_, err = nonBilledEvaluatorService.RunResponseQualityEvaluator(nonBilledEvaluator, ctx, dataset)
	if err != nil {
		log.Fatalf("could not evaluate: %v", err)
	}
}

// runBilledEvaluator demonstrates a billed custom evaluator
func runBilledEvaluator(ctx context.Context, g *genkit.Genkit, chatFlow *core.Flow[flows.ChatMessage, flows.ChatMessage, struct{}]) {
	log.Println("Running billed custom evaluator...")

	billedEvaluatorService := evals.NewBilledEvaluatorService(g)

	maliciousnessEvaluator, err := billedEvaluatorService.NewMaliciousnessEvaluator()
	if err != nil {
		log.Fatalf("could not define evaluator: %v", err)
	}

	log.Println("Custom evaluator defined:", maliciousnessEvaluator.Name())

	// Auto-generate dataset using the chat flow
	log.Println("Auto-generating maliciousness evaluation dataset...")
	maliciousnessDataset, err := billedEvaluatorService.GenerateMaliciousnessEvaluatorDatasetWithChatFlow(ctx, chatFlow)
	if err != nil {
		log.Fatalf("could not generate dataset: %v", err)
	}

	_, err = billedEvaluatorService.RunMaliciousnessEvaluator(maliciousnessEvaluator, ctx, maliciousnessDataset)
	if err != nil {
		log.Fatalf("could not evaluate: %v", err)
	}
}

// runGenkitBuiltinEvaluator demonstrates a Genkit built-in evaluator
func runGenkitBuiltinEvaluator(ctx context.Context, g *genkit.Genkit) {
	log.Println("Running Genkit built-in evaluator...")

	genkitEvaluatorService := evals.NewGenkitEvaluatorService(g)

	genkitEvaluator, err := genkitEvaluatorService.GetDeepEqualEvaluator()
	if err != nil {
		log.Fatalf("could not get Genkit evaluator: %v", err)
	}

	log.Println("Genkit evaluator found:", genkitEvaluator.Name())

	genkitDataset := genkitEvaluatorService.GetSampleDataset()

	_, err = genkitEvaluatorService.RunGenkitEvaluator(genkitEvaluator, ctx, genkitDataset)
	if err != nil {
		log.Fatalf("could not evaluate with Genkit evaluator: %v", err)
	}
}
