package flows

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"mastering-genkit-go/example/chapter-15/internal/structs/client"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewCreateRecipeFlow creates a streaming recipe generation flow
func NewCreateRecipeFlow(g *genkit.Genkit, tools []ai.ToolRef) *core.Flow[client.RecipeRequest, client.RecipeResponse, client.RecipeResponse] {
	return genkit.DefineStreamingFlow(
		g,
		"createRecipe",
		func(ctx context.Context, input client.RecipeRequest, callback func(context.Context, client.RecipeResponse) error) (client.RecipeResponse, error) {
			log.Printf("createRecipe flow called with ingredients: %v", input.Ingredients)

			systemPrompt := `You are a professional chef creating recipes for a cooking battle.
Create detailed, creative recipes that make the best use of available ingredients.
Use the available tools to check ingredient availability and calculate nutrition.`

			prompt := fmt.Sprintf(`Create a recipe using the following ingredients:
Ingredients: %v

Provide a complete recipe with:
1. Recipe name
2. Ingredients list with quantities
3. Step-by-step instructions
4. Cooking time
5. Nutritional information (use the available tools)`,
				input.Ingredients)

			// Send streaming updates
			if err := callback(ctx, client.RecipeResponse{
				Type:    "content",
				Content: "🍳 Starting recipe creation...\n",
			}); err != nil {
				return client.RecipeResponse{
					Type:  "error",
					Error: "Failed to send initial message",
				}, err
			}

			if err := callback(ctx, client.RecipeResponse{
				Type:    "content",
				Content: "👨‍🍳 Analyzing ingredients and creating custom recipe...\n",
			}); err != nil {
				return client.RecipeResponse{
					Type:  "error",
					Error: "Failed to send progress message",
				}, err
			}

			// Generate recipe
			resp, err := genkit.Generate(ctx, g,
				ai.WithSystem(systemPrompt),
				ai.WithPrompt(prompt),
				ai.WithTools(tools...),
			)
			if err != nil {
				log.Printf("Failed to generate recipe: %v", err)
				return client.RecipeResponse{
					Type:  "error",
					Error: fmt.Sprintf("Failed to generate recipe: %v", err),
				}, err
			}

			// Stream the recipe content in smaller chunks
			if err := streamRecipeContent(ctx, callback, resp.Text()); err != nil {
				return client.RecipeResponse{
					Type:  "error",
					Error: "Failed to send recipe content",
				}, err
			}

			// Send completion signal
			if err := callback(ctx, client.RecipeResponse{
				Type: "done",
			}); err != nil {
				return client.RecipeResponse{
					Type:  "error",
					Error: "Failed to send completion signal",
				}, err
			}

			return client.RecipeResponse{
				Type:    "content",
				Content: resp.Text(),
			}, nil
		},
	)
}

// streamRecipeContent streams recipe content in smaller, more digestible chunks
func streamRecipeContent(ctx context.Context, callback func(context.Context, client.RecipeResponse) error, content string) error {
	// Split content into lines for better streaming experience
	lines := strings.Split(content, "\n")
	currentChunk := ""

	for i, line := range lines {
		// Add current line to chunk
		currentChunk += line + "\n"

		// Send chunk when we have enough content or reach certain markers
		shouldSend := false

		// Send if chunk is getting long (around 200 characters)
		if len(currentChunk) > 200 {
			shouldSend = true
		}

		// Send at natural break points (empty lines, numbered lists, headers)
		if strings.TrimSpace(line) == "" ||
			strings.HasPrefix(strings.TrimSpace(line), "#") ||
			strings.Contains(line, ":") && len(line) < 50 {
			shouldSend = true
		}

		// Send if this is the last line
		if i == len(lines)-1 {
			shouldSend = true
		}

		if shouldSend && strings.TrimSpace(currentChunk) != "" {
			// Add a small delay for better streaming experience
			time.Sleep(100 * time.Millisecond)

			if err := callback(ctx, client.RecipeResponse{
				Type:    "content",
				Content: currentChunk,
			}); err != nil {
				return err
			}

			currentChunk = ""
		}
	}

	return nil
}
