package flows

import (
	"context"
	"fmt"
	nested "mastering-genkit-go/example/chapter-05/internal/schemas/2-nested"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewNestedStructuredFlow creates a flow that demonstrates complex nested structures.
// It generates recipes with nested ingredients, steps, and nutrition maps.
func NewNestedStructuredFlow(g *genkit.Genkit) *core.Flow[nested.RecipeRequest, nested.Recipe, struct{}] {
	return genkit.DefineFlow(g, "nestedStructuredFlow", func(ctx context.Context, input nested.RecipeRequest) (nested.Recipe, error) {
		// Use GenerateData to get structured output with nested structures
		result, _, err := genkit.GenerateData[nested.Recipe](ctx, g,
			ai.WithSystem(`Create a detailed recipe with ingredients, steps, nutrition info, cooking times, and difficulty level. Consider dietary restrictions.`),
			ai.WithPrompt(`Create a recipe for: %s
Servings: %d
Dietary restrictions: %v`,
				input.DishName,
				input.Servings,
				input.DietaryRestrictions),
			ai.WithConfig(map[string]interface{}{
				"temperature": 0.5,
			}),
		)
		if err != nil {
			return nested.Recipe{}, fmt.Errorf("failed to generate recipe: %w", err)
		}

		return *result, nil
	})
}
