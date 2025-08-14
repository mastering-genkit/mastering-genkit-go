package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewBasicGenerationFlow creates a simple text generation flow
func NewBasicGenerationFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "basicGenerationFlow", func(ctx context.Context, userRequest string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithPrompt(fmt.Sprintf("As a helpful cooking instructor, explain %s in simple terms that a beginner can understand.", userRequest)),
			ai.WithMiddleware(LoggingMiddleware),
		)
		if err != nil {
			return "", fmt.Errorf("failed to generate response: %w", err)
		}
		return resp.Text(), nil
	})
}