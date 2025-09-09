package flows

import (
	"context"
	"fmt"
	generation "mastering-genkit-go/example/chapter-05/internal/structs/character-generation"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
	"google.golang.org/genai"
)

// NewCharacterGenerationFlow creates a flow that demonstrates image generation with structured metadata.
// It generates an anime character image and returns structured character information.
func NewCharacterGenerationFlow(g *genkit.Genkit) *core.Flow[generation.CharacterRequest, generation.CharacterResult, struct{}] {
	return genkit.DefineFlow(g, "characterGenerationFlow", func(ctx context.Context, input generation.CharacterRequest) (generation.CharacterResult, error) {
		// Step 1: Generate the image using Gemini 2.5 Flash Image (aka Nano Banana)
		imagePrompt := fmt.Sprintf("%s style: %s",
			input.Style,
			input.Description)

		imageResp, err := genkit.Generate(ctx, g,
			ai.WithModelName("googleai/gemini-2.5-flash-image-preview"),
			ai.WithPrompt(imagePrompt),
			ai.WithConfig(&genai.GenerateContentConfig{
				ResponseModalities: []string{"IMAGE"},
			}),
		)
		if err != nil {
			return generation.CharacterResult{}, fmt.Errorf("failed to generate image: %w", err)
		}

		// Extract the generated image data
		var imageData string
		if len(imageResp.Message.Content) > 0 {
			// The image is returned as a data URI
			imageData = imageResp.Message.Content[0].Text
		}

		// Step 2: Generate structured character metadata
		characterData, _, err := genkit.GenerateData[generation.Character](ctx, g,
			ai.WithSystem(`Create character profile with name, visual description, traits, and brief story.`),
			ai.WithPrompt(`Create character details for: %s in %s style`,
				input.Description,
				input.Style),
			ai.WithConfig(map[string]interface{}{
				"temperature": 0.7,
			}),
		)
		if err != nil {
			return generation.CharacterResult{}, fmt.Errorf("failed to generate character data: %w", err)
		}

		// Combine image and metadata
		result := generation.CharacterResult{
			ImageData: imageData,
			Character: *characterData,
		}

		return result, nil
	})
}
