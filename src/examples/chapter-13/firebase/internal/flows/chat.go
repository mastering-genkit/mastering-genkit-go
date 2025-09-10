package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// ChatMessage represents the input for the chat flow
type ChatMessage struct {
	Message string `json:"message,omitempty"`
}

// ChatResponse represents the output from the chat flow
type ChatResponse struct {
	Response string        `json:"response"`
	History  []*ai.Message `json:"history"`
}

// NewChatFlow creates a flow that reads user messages, processes them with conversation history, and returns a response.
func NewChatFlow(g *genkit.Genkit, tools []ai.ToolRef) *core.Flow[ChatMessage, ChatMessage, struct{}] {
	return genkit.DefineFlow(g, "chatFlow", func(ctx context.Context, req ChatMessage) (ChatMessage, error) {

		systemPrompt := "You're a helpful AI assistant. Respond to the user's message in a helpful manner."
		prompt := fmt.Sprintf("User: %s\nAI:", req.Message)

		// Generate response using Genkit
		resp, err := genkit.Generate(ctx, g,
			ai.WithSystem(systemPrompt),
			ai.WithPrompt(prompt),
			ai.WithTools(tools...),
		)
		if err != nil {
			return ChatMessage{}, fmt.Errorf("failed to generate response: %w", err)
		}

		return ChatMessage{
			Message: resp.Text(),
		}, nil
	})
}
