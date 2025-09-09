package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
	"google.golang.org/genai"
)

// NewImageGenerationFlow creates a flow for generating images from text descriptions
func NewImageGenerationFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "imageGenerationFlow", func(ctx context.Context, description string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithModelName("googleai/gemini-2.5-flash-image-preview"),
			ai.WithPrompt(description),
			ai.WithConfig(&genai.GenerateContentConfig{
				ResponseModalities: []string{"IMAGE"},
			}),
		)
		if err != nil {
			return "", fmt.Errorf("failed to generate image: %w", err)
		}

		for _, part := range resp.Message.Content {
			if part.IsMedia() {
				// The Text field contains the image data (base64 encoded data URI)
				return part.Text, nil
			}
		}

		return "", fmt.Errorf("no image generated in response")
	})
}
