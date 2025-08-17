package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewDotpromptFlow creates a flow using Dotprompt templates
func NewDotpromptFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	cookingInstructor := genkit.LookupPrompt(g, "cooking-instructor")
	if cookingInstructor == nil {
		panic("no prompt named 'cooking-instructor' found")
	}

	return genkit.DefineFlow(g, "dotpromptFlow", func(ctx context.Context, userRequest string) (string, error) {
		resp, err := cookingInstructor.Execute(ctx,
			ai.WithInput(map[string]any{
				"topic": userRequest,
			}),
			ai.WithMiddleware(LoggingMiddleware),
		)
		if err != nil {
			return "", fmt.Errorf("error executing prompt: %w", err)
		}
		return resp.Text(), nil
	})
}