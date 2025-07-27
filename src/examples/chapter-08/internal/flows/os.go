package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewOperatingSystemFlow creates a flow that interacts with the operating system using AI tools.
// It can be used to perform tasks like listing directories or getting the current date.
func NewOperatingSystemFlow(g *genkit.Genkit, tools []ai.ToolRef) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "operatingSystemFlow", func(ctx context.Context, userRequest string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithSystem("You are an AI assistant that can interact with the operating system. Use the available tools to perform tasks."),
			ai.WithPrompt(`The user wants to: %s`, userRequest),
			ai.WithTools(tools...),
		)
		if err != nil {
			return "", fmt.Errorf("failed to generate response: %w", err)
		}

		return resp.Text(), nil
	})
}
