package agent

import (
	"context"
	"time"

	"github.com/firebase/genkit/go/ai"
	aix "github.com/firebase/genkit/go/ai/exp"
	"github.com/firebase/genkit/go/genkit"
	genkitx "github.com/firebase/genkit/go/genkit/exp"
)

// DateTimeInput is the (empty) input of the getCurrentDateTime tool.
type DateTimeInput struct{}

// NewAssistantAgent registers the assistant's tools and the agent itself.
// The agent keeps no custom state (State is any): its memory is the
// conversation history that Genkit stores in every session snapshot.
func NewAssistantAgent(g *genkit.Genkit, model string, store aix.SessionStore[any]) *aix.Agent[any] {
	// A tool the model can call to answer questions about the current time
	getCurrentDateTime := genkitx.DefineTool(g, "getCurrentDateTime",
		"Returns the current local date and time.",
		func(ctx context.Context, _ DateTimeInput) (string, error) {
			return time.Now().Format(time.RFC1123), nil
		})

	return genkitx.DefineAgent(g, "assistant",
		aix.InlinePrompt{
			ai.WithModelName(model),
			ai.WithSystem("You're a helpful AI assistant. Respond to the user's message in a helpful and concise manner. " +
				"When the user asks about the current date or time, call getCurrentDateTime."),
			ai.WithTools(getCurrentDateTime),
			// Cap the tool-calling loop of a single turn
			ai.WithMaxTurns(5),
		},
		aix.WithSessionStore(store),
		aix.WithDescription[any]("Helpful assistant with conversation memory"),
	)
}
