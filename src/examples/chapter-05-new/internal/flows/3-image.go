package flows

import (
	"context"
	"fmt"
	image "mastering-genkit-go/example/chapter-05/internal/schemas/3-image"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewImageAnalysisFlow creates a flow that demonstrates image input with structured output.
// It analyzes product images and returns structured product information.
func NewImageAnalysisFlow(g *genkit.Genkit) *core.Flow[image.ImageAnalysisRequest, image.ProductInfo, struct{}] {
	return genkit.DefineFlow(g, "imageAnalysisFlow", func(ctx context.Context, input image.ImageAnalysisRequest) (image.ProductInfo, error) {
		// Use GenerateData with image URL
		result, _, err := genkit.GenerateData[image.ProductInfo](ctx, g,
			ai.WithSystem(fmt.Sprintf(`Analyze the product image and extract: category, features, colors, price range, target audience, and description.
Output language: %s`, input.Language)),
			ai.WithMessages(
				ai.NewUserMessage(
					ai.NewTextPart("Analyze this product image and provide structured information."),
					ai.NewMediaPart("", input.ImageURL), // Empty MIME type, URL directly
				),
			),
			ai.WithConfig(map[string]interface{}{
				"temperature": 0.4, // Lower temperature for factual analysis
			}),
		)
		if err != nil {
			return image.ProductInfo{}, fmt.Errorf("failed to analyze image: %w", err)
		}

		return *result, nil
	})
}
