package flows

import (
	"context"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// Message represents the input for the simple flow
type Message struct {
	Message string `json:"message,omitempty"`
}

// NewSimpleFlow creates a flow that reads user messages, processes them with conversation history, and returns a response.
func NewSimpleFlow(g *genkit.Genkit, tools []ai.ToolRef) *core.Flow[Message, Message, struct{}] {
	return genkit.DefineFlow(g, "simpleFlow", func(ctx context.Context, req Message) (Message, error) {

		return Message{
			Message: "Hello from the cloud! You said: " + req.Message,
		}, nil
	})
}
