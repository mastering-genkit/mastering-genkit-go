package flows

import (
	"context"
	"fmt"
	simple "mastering-genkit-go/example/chapter-05/internal/structs/simple-structured"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewSimpleStructuredFlow creates a flow that demonstrates basic structured output.
// It analyzes product reviews and returns structured sentiment analysis.
func NewSimpleStructuredFlow(g *genkit.Genkit) *core.Flow[simple.ReviewInput, simple.ReviewAnalysis, struct{}] {
	return genkit.DefineFlow(g, "simpleStructuredFlow", func(ctx context.Context, input simple.ReviewInput) (simple.ReviewAnalysis, error) {
		result, _, err := genkit.GenerateData[simple.ReviewAnalysis](ctx, g,
			ai.WithSystem(`Analyze the review and extract: sentiment (positive/negative/neutral), score (1-5), keywords, and summary.`),
			ai.WithPrompt("Analyze this review: %s", input.ReviewText),
			ai.WithConfig(map[string]interface{}{
				"temperature": 0.3,
			}),
		)
		if err != nil {
			return simple.ReviewAnalysis{}, fmt.Errorf("failed to analyze review: %w", err)
		}

		return *result, nil
	})
}
