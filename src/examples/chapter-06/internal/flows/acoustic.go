package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewAcousticInstrumentFlow creates a flow that recommends acoustic instruments
// based on the input genre characteristics.
func NewAcousticInstrumentFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "acousticInstrumentFlow", func(ctx context.Context, genre string) (string, error) {
		// Load the dotprompt
		acousticPrompt := genkit.LookupPrompt(g, "acoustic-instrument")
		
		// Execute the prompt
		resp, err := acousticPrompt.Execute(ctx,
			ai.WithInput(map[string]any{
				"genre": genre,
			}))
		if err != nil {
			return "", fmt.Errorf("failed to recommend acoustic instrument: %w", err)
		}

		return resp.Text(), nil
	})
}