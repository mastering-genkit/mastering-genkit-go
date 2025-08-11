package flows

import (
	"context"
	"fmt"
	audio "mastering-genkit-go/example/chapter-05/internal/schemas/4-audio"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewAudioTranscriptionFlow creates a flow that demonstrates audio input with structured output.
// It transcribes meeting audio and returns structured meeting minutes.
func NewAudioTranscriptionFlow(g *genkit.Genkit) *core.Flow[audio.AudioRequest, audio.MeetingMinutes, struct{}] {
	return genkit.DefineFlow(g, "audioTranscriptionFlow", func(ctx context.Context, input audio.AudioRequest) (audio.MeetingMinutes, error) {
		// Note: Direct audio URL processing might not be supported by all models
		// This is a demonstration of the pattern - actual implementation may require
		// using specific audio transcription services first
		
		// For this example, we'll use a text prompt to simulate the process
		// In production, you'd first transcribe the audio using a service like
		// Google Speech-to-Text API, then process the transcription
		
		result, _, err := genkit.GenerateData[audio.MeetingMinutes](ctx, g,
			ai.WithSystem(fmt.Sprintf(`Analyze the meeting audio and extract: full transcription, speaker summaries, decisions, action items with assignees and due dates, and keywords.
Meeting topic: %s
Language: %s`, input.Topic, input.Language)),
			ai.WithMessages(
				ai.NewUserMessage(
					ai.NewTextPart("Process this meeting recording and provide structured minutes."),
					// Note: Audio URL processing would require a model that supports audio
					// or pre-processing with a transcription service
					ai.NewMediaPart("", input.AudioURL),
				),
			),
			ai.WithConfig(map[string]interface{}{
				"temperature": 0.3, // Lower temperature for accurate transcription
			}),
		)
		if err != nil {
			return audio.MeetingMinutes{}, fmt.Errorf("failed to process audio: %w", err)
		}

		return *result, nil
	})
}