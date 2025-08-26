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
			// Buffer to accumulate chunks until we have a meaningful unit
			var buffer string
			
			// Generate with real-time streaming
			final, err := genkit.Generate(ctx, g,
				ai.WithSystem("You are a professional chef creating detailed, easy-to-follow cooking instructions."),
				ai.WithPrompt(fmt.Sprintf(`Create a detailed recipe for "%s".`, dish)),
				ai.WithStreaming(func(ctx context.Context, chunk *ai.ModelResponseChunk) error {
					// Accumulate chunk text
					for _, content := range chunk.Content {
						buffer += content.Text
					}
					
					// Stream when we have a complete sentence or paragraph
					// Look for sentence endings or newlines
					for {
						// Find a good breaking point (period, newline, or exclamation/question mark)
						breakPoint := -1
						for i, r := range buffer {
							if r == '.' || r == '\n' || r == '!' || r == '?' {
								// Make sure it's not an abbreviation (simple check)
								if i+1 < len(buffer) && (buffer[i+1] == ' ' || buffer[i+1] == '\n') {
									breakPoint = i + 1
									break
								}
							}
						}
						
						// If we found a breaking point, send that portion
						if breakPoint > 0 {
							toSend := buffer[:breakPoint]
							buffer = buffer[breakPoint:]
							
							if err := stream(ctx, toSend); err != nil {
								return err
							}
						} else {
							// No complete sentence yet, wait for more
							break
						}
					}
					
					return nil
				}),
			)

			if err != nil {
				return "", fmt.Errorf("failed to generate recipe: %w", err)
			}
			
			// Send any remaining buffer content
			if buffer != "" {
				if err := stream(ctx, buffer); err != nil {
					return "", fmt.Errorf("failed to send final buffer: %w", err)
				}
			}

			return final.Text(), nil
		})
}
