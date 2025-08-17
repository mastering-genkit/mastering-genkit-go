package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewAudioTranscriptionFlow creates a flow for transcribing audio from URLs
func NewAudioTranscriptionFlow(g *genkit.Genkit) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "audioTranscriptionFlow", func(ctx context.Context, audioURL string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithModelName("googleai/gemini-2.5-flash"),
			ai.WithMessages(
				ai.NewUserMessage(
					ai.NewTextPart("Transcribe this audio accurately."),
					ai.NewMediaPart("", audioURL),
				),
			),
			ai.WithMiddleware(LoggingMiddleware),
		)
		if err != nil {
			return "", fmt.Errorf("failed to process audio: %w", err)
		}
		return resp.Text(), nil
	})
}