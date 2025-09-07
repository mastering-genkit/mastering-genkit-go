package flows

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"mastering-genkit-go/example/chapter-15/internal/structs/client"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewCookingBattleChatFlow creates a streaming chat flow for cooking battle
func NewCookingBattleChatFlow(g *genkit.Genkit, tools []ai.ToolRef) *core.Flow[client.ChatRequest, client.ChatResponse, client.ChatResponse] {
	return genkit.DefineStreamingFlow(
		g,
		"cookingBattleChat",
		func(ctx context.Context, input client.ChatRequest, callback func(context.Context, client.ChatResponse) error) (client.ChatResponse, error) {
			log.Printf("cookingBattleChat flow called with %d messages", len(input.Messages))

			// System prompt for cooking battle
			systemPrompt := `You are a cooking battle assistant helping users create amazing dishes.
You have access to tools for searching recipes, checking ingredient stock, and calculating nutrition.
Guide users through the cooking process with enthusiasm and expertise.
When users complete a dish, evaluate it fairly and provide constructive feedback.`

			// Add context about ingredients and constraints if provided
			if len(input.Ingredients) > 0 || len(input.Constraints) > 0 {
				contextMsg := "\n\nContext for this cooking battle:\n"
				if len(input.Ingredients) > 0 {
					contextMsg += fmt.Sprintf("Available ingredients: %v\n", input.Ingredients)
				}
				if len(input.Constraints) > 0 {
					constraintsJSON, _ := json.Marshal(input.Constraints)
					contextMsg += fmt.Sprintf("Constraints: %s\n", string(constraintsJSON))
				}
				systemPrompt += contextMsg
			}

			// Build conversation prompt from messages
			prompt := ""
			for i, msg := range input.Messages {
				if i > 0 {
					prompt += "\n"
				}
				switch msg.Role {
				case "user":
					prompt += fmt.Sprintf("User: %s", msg.Content)
				case "assistant":
					prompt += fmt.Sprintf("Assistant: %s", msg.Content)
				}
			}

			// Buffer to accumulate chunks until we have a meaningful unit
			var buffer string

			// Generate streaming response with tools
			final, err := genkit.Generate(ctx, g,
				ai.WithSystem(systemPrompt),
				ai.WithPrompt(prompt),
				ai.WithTools(tools...),
				ai.WithStreaming(func(ctx context.Context, chunk *ai.ModelResponseChunk) error {
					// Accumulate chunk text
					for _, content := range chunk.Content {
						buffer += content.Text
					}

					// Stream when we have a complete sentence or paragraph
					// Look for sentence endings or newlines
					for {
						// Find a good breaking point (period, newline, or exclamation/question mark)
						breakPoint := -1
						for i, r := range buffer {
							if r == '.' || r == '\n' || r == '!' || r == '?' {
								// Make sure it's not an abbreviation (simple check)
								if i+1 < len(buffer) && (buffer[i+1] == ' ' || buffer[i+1] == '\n') {
									breakPoint = i + 1
									break
								}
							}
						}

						// If we found a breaking point, send that portion
						if breakPoint > 0 {
							toSend := buffer[:breakPoint]
							buffer = buffer[breakPoint:]

							if err := callback(ctx, client.ChatResponse{
								Type:    "content",
								Content: toSend,
							}); err != nil {
								return err
							}
						} else {
							// No complete sentence yet, wait for more
							break
						}
					}

					return nil
				}),
			)

			if err != nil {
				log.Printf("Generation error: %v", err)
				return client.ChatResponse{
					Type:  "error",
					Error: fmt.Sprintf("Failed to generate response: %v", err),
				}, err
			}

			// Send any remaining buffer content
			if buffer != "" {
				if err := callback(ctx, client.ChatResponse{
					Type:    "content",
					Content: buffer,
				}); err != nil {
					return client.ChatResponse{
						Type:  "error",
						Error: "Failed to send final buffer",
					}, err
				}
			}

			// Send done signal
			if err := callback(ctx, client.ChatResponse{
				Type: "done",
			}); err != nil {
				return client.ChatResponse{
					Type:  "error",
					Error: "Failed to send done signal",
				}, err
			}

			return client.ChatResponse{
				Type:    "content",
				Content: final.Text(),
			}, nil
		},
	)
}
