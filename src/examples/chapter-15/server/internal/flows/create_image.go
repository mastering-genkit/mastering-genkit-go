package flows

import (
	"context"
	"fmt"
	"log"

	"mastering-genkit-go/example/chapter-15/internal/structs/client"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
	"google.golang.org/genai"
)

// NewCreateImageFlow creates an image generation flow
func NewCreateImageFlow(g *genkit.Genkit) *core.Flow[client.ImageRequest, client.ImageResponse, struct{}] {
	return genkit.DefineFlow(
		g,
		"createImage",
		func(ctx context.Context, input client.ImageRequest) (client.ImageResponse, error) {
			log.Printf("createImage flow called for dish: %s", input.DishName)

			prompt := fmt.Sprintf(`Create a high-quality, appetizing image of a completed dish:
Dish Name: %s
Description: %s

The image should be:
- Photorealistic and appetizing
- Well-plated and professionally presented
- Showing the dish from an attractive angle
- With good lighting and composition`,
				input.DishName, input.Description)

			// Generate image using Imagen3
			resp, err := genkit.Generate(ctx, g,
				ai.WithModelName("googleai/imagen-3.0-generate-002"),
				ai.WithPrompt(prompt),
				ai.WithConfig(&genai.GenerateImagesConfig{
					NumberOfImages:    1,
					AspectRatio:       "1:1",
					SafetyFilterLevel: genai.SafetyFilterLevelBlockLowAndAbove,
					PersonGeneration:  genai.PersonGenerationAllowAll,
					OutputMIMEType:    "image/png",
				}),
			)
			if err != nil {
				log.Printf("Failed to generate image: %v", err)
				return client.ImageResponse{
					Success: false,
					Error:   fmt.Sprintf("Failed to generate image: %v", err),
				}, err
			}

			// Extract image data from response
			var imageData string
			for _, part := range resp.Message.Content {
				if part.IsMedia() {
					// The Text field contains the image data (base64 encoded data URI)
					imageData = part.Text
					break
				}
			}

			if imageData == "" {
				return client.ImageResponse{
					Success: false,
					Error:   "No image generated in response",
				}, fmt.Errorf("no image generated")
			}

			return client.ImageResponse{
				Success:  true,
				ImageUrl: imageData,
				DishName: input.DishName,
				Error:    "",
			}, nil
		},
	)
}
