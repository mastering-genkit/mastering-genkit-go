package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewElectronicInstrumentFlow creates a flow that recommends electronic instruments
// based on the input genre characteristics.
func NewElectronicInstrumentFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "electronicInstrumentFlow", func(ctx context.Context, genre string) (string, error) {
		// Load the dotprompt
		electronicPrompt := genkit.LookupPrompt(g, "electronic-instrument")
		
		// Execute the prompt
		resp, err := electronicPrompt.Execute(ctx,
			ai.WithInput(map[string]any{
				"genre": genre,
			}))
		if err != nil {
			return "", fmt.Errorf("failed to recommend electronic instrument: %w", err)
		}

		return resp.Text(), nil
	})
}