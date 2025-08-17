package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewRecipeFlow creates a streaming flow that generates recipe content chunk by chunk
func NewRecipeFlow(g *genkit.Genkit) *core.Flow[string, string, string] {
	return genkit.DefineStreamingFlow(g, "recipeStepsFlow",
		func(ctx context.Context, dish string, stream func(context.Context, string) error) (string, error) {
			// Generate with real-time streaming
			final, err := genkit.Generate(ctx, g,
				ai.WithSystem("You are a professional chef creating detailed, easy-to-follow cooking instructions."),
				ai.WithPrompt(fmt.Sprintf(`Create a detailed recipe for "%s". Include step-by-step instructions and at the end provide a JSON summary with total_steps count and summary message.`, dish)),
				ai.WithStreaming(func(ctx context.Context, chunk *ai.ModelResponseChunk) error {
					// Stream each chunk as pure text
					chunkText := ""
					for _, content := range chunk.Content {
						chunkText += content.Text
					}

					// If we have text, send it as a chunk
					if chunkText != "" {
						return stream(ctx, chunkText)
					}

					return nil
				}),
			)

			if err != nil {
				return "", fmt.Errorf("failed to generate recipe: %w", err)
			}

			return final.Text(), nil
		})
}
