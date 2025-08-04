package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// ChatRequest represents the input for the chat flow
type ChatRequest struct {
	Message string        `json:"message,omitempty"`
	History []*ai.Message `json:"history,omitempty"`
}

// ChatResponse represents the output from the chat flow
type ChatResponse struct {
	Response string        `json:"response"`
	History  []*ai.Message `json:"history"`
}

// NewChatFlow creates a flow that reads user messages, processes them with conversation history, and returns a response.
func NewChatFlow(g *genkit.Genkit, tools []ai.ToolRef, model ai.ModelRef) *core.Flow[ChatRequest, ChatResponse, struct{}] {
	return genkit.DefineFlow(g, "chatFlow", func(ctx context.Context, req ChatRequest) (ChatResponse, error) {
		// Prepare the conversation history with the new user message
		messages := make([]*ai.Message, len(req.History))
		copy(messages, req.History)

		// Add the new user message to the conversation
		userMessage := &ai.Message{
			Content: []*ai.Part{ai.NewTextPart(req.Message)},
			Role:    ai.RoleUser,
		}
		messages = append(messages, userMessage)

		systemPrompt := "You're a helpful AI assistant. Respond to the user's message in a helpful manner."

		// Generate response using Genkit
		resp, err := genkit.Generate(ctx, g,
			ai.WithSystem(systemPrompt),
			ai.WithMessages(messages...),
			ai.WithTools(tools...),
			ai.WithModel(model),
		)
		if err != nil {
			return ChatResponse{}, fmt.Errorf("failed to generate response: %w", err)
		}

		// Extract the response text
		responseText := resp.Text()

		// Add the assistant's response to the conversation history
		assistantMessage := &ai.Message{
			Content: []*ai.Part{ai.NewTextPart(responseText)},
			Role:    ai.RoleModel,
		}
		messages = append(messages, assistantMessage)

		return ChatResponse{
			Response: responseText,
			History:  messages,
		}, nil
	})
}
