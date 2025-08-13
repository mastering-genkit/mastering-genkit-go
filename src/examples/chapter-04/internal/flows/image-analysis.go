package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewImageAnalysisFlow creates a flow for analyzing images from URLs
func NewImageAnalysisFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "imageAnalysisFlow", func(ctx context.Context, imageURL string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithModelName("googleai/gemini-2.5-flash"),
			ai.WithMessages(
				ai.NewUserMessage(
					ai.NewTextPart("What's in this image? Describe it in detail."),
					ai.NewMediaPart("", imageURL),
				),
			),
			ai.WithMiddleware(LoggingMiddleware),
		)
		if err != nil {
			return "", fmt.Errorf("failed to analyze image: %w", err)
		}
		return resp.Text(), nil
	})
}
