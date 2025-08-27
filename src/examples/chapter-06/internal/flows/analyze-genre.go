package flows

import (
	"context"
	"fmt"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewAnalyzeGenreFlow creates a flow that analyzes a music genre and categorizes it
// as either "acoustic" or "electronic" based on its characteristics.
func NewAnalyzeGenreFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "analyzeGenreFlow", func(ctx context.Context, genre string) (string, error) {
		// Load the dotprompt
		analyzePrompt := genkit.LookupPrompt(g, "analyze-genre")

		// Execute the prompt
		resp, err := analyzePrompt.Execute(ctx,
			ai.WithInput(map[string]any{
				"genre": genre,
			}))
		if err != nil {
			return "", fmt.Errorf("failed to analyze genre: %w", err)
		}

		// Clean the response: remove quotes, trim whitespace, convert to lowercase
		result := strings.ToLower(strings.TrimSpace(resp.Text()))
		result = strings.Trim(result, `"'`) // Remove any quotes
		return result, nil
	})
}
