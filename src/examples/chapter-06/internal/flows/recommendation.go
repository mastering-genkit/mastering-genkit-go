package flows

import (
	"context"
	"examples/chapter-06/internal/structs"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewRecommendationFlow creates an orchestrator flow that coordinates multiple flows
// to provide instrument recommendations based on music genre.
func NewRecommendationFlow(
	g *genkit.Genkit,
	analyzeFlow *core.Flow[string, string, struct{}],
	acousticFlow *core.Flow[string, string, struct{}],
	electronicFlow *core.Flow[string, string, struct{}],
) *core.Flow[structs.RecommendationInput, structs.RecommendationOutput, struct{}] {
	return genkit.DefineFlow(g, "recommendationFlow", func(ctx context.Context, input structs.RecommendationInput) (structs.RecommendationOutput, error) {
		// Step 1: Analyze the genre using Flow_A
		genreCategory, err := analyzeFlow.Run(ctx, input.Genre)
		if err != nil {
			return structs.RecommendationOutput{}, fmt.Errorf("failed to analyze genre: %w", err)
		}

		// Step 2: Get instrument recommendation based on category (Flow_B or Flow_C)
		var instrument string
		if genreCategory == "acoustic" {
			// Invoke Flow_B
			instrument, err = acousticFlow.Run(ctx, input.Genre)
			if err != nil {
				return structs.RecommendationOutput{}, fmt.Errorf("failed to get acoustic instrument: %w", err)
			}
		} else {
			// Invoke Flow_C
			instrument, err = electronicFlow.Run(ctx, input.Genre)
			if err != nil {
				return structs.RecommendationOutput{}, fmt.Errorf("failed to get electronic instrument: %w", err)
			}
		}

		// Step 3: Generate detailed recommendation with reasoning and starter items
		experience := input.Experience
		if experience == "" {
			experience = "beginner"
		}

		// Load the details prompt
		detailsPrompt := genkit.LookupPrompt(g, "recommendation-details")

		// Execute the prompt to get detailed recommendation
		resp, err := detailsPrompt.Execute(ctx,
			ai.WithInput(map[string]any{
				"genre":      input.Genre,
				"category":   genreCategory,
				"instrument": instrument,
				"experience": experience,
			}))
		if err != nil {
			return structs.RecommendationOutput{}, fmt.Errorf("failed to generate detailed recommendation: %w", err)
		}

		// Manual unmarshaling
		var output structs.RecommendationOutput
		if err := resp.Output(&output); err != nil {
			return structs.RecommendationOutput{}, fmt.Errorf("failed to parse recommendation details: %w", err)
		}

		return output, nil
	})
}
