package flows

import (
	"context"
	"fmt"
	"log"

	"mastering-genkit-go/example/chapter-15/internal/structs/client"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
	"google.golang.org/genai"
)

// NewCookingBattleActionFlow creates a non-streaming action flow for cooking battle
func NewCookingBattleActionFlow(g *genkit.Genkit, tools []ai.ToolRef) *core.Flow[client.ActionRequest, client.ActionResponse, struct{}] {
	return genkit.DefineFlow(
		g,
		"cookingBattleAction",
		func(ctx context.Context, input client.ActionRequest) (client.ActionResponse, error) {
			log.Printf("cookingBattleAction flow called with action: %s", input.Action)

			switch input.Action {
			case "generate_recipe":
				return generateRecipe(ctx, g, tools, input)
			case "create_image":
				return createDishImage(ctx, g, input)
			case "evaluate":
				return evaluateDish(ctx, g, input)
			default:
				return client.ActionResponse{
					Success: false,
					Error:   fmt.Sprintf("Unknown action: %s", input.Action),
				}, fmt.Errorf("unknown action: %s", input.Action)
			}
		},
	)
}

// generateRecipe generates a recipe based on ingredients and constraints
func generateRecipe(ctx context.Context, g *genkit.Genkit, tools []ai.ToolRef, input client.ActionRequest) (client.ActionResponse, error) {
	systemPrompt := `You are a professional chef creating recipes for a cooking battle.
Create detailed, creative recipes that make the best use of available ingredients.
Use the available tools to check ingredient availability and calculate nutrition.`

	prompt := fmt.Sprintf(`Create a recipe using the following:
Ingredients: %v
Constraints: %v

Provide a complete recipe with:
1. Recipe name
2. Ingredients list with quantities
3. Step-by-step instructions
4. Cooking time
5. Nutritional information (use the calculateNutrition tool)`,
		input.Ingredients, input.Constraints)

	resp, err := genkit.Generate(ctx, g,
		ai.WithSystem(systemPrompt),
		ai.WithPrompt(prompt),
		ai.WithTools(tools...),
	)
	if err != nil {
		log.Printf("Failed to generate recipe: %v", err)
		return client.ActionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to generate recipe: %v", err),
		}, err
	}

	return client.ActionResponse{
		Success: true,
		Result: map[string]interface{}{
			"recipe": resp.Text(),
		},
	}, nil
}

// createDishImage generates an image of the completed dish
func createDishImage(ctx context.Context, g *genkit.Genkit, input client.ActionRequest) (client.ActionResponse, error) {
	prompt := fmt.Sprintf(`Create a high-quality, appetizing image of a completed dish:
Dish Name: %s
Description: %s

The image should be:
- Photorealistic and appetizing
- Well-plated and professionally presented
- Showing the dish from an attractive angle
- With good lighting and composition`,
		input.DishName, input.Description)

	// Generate image using Imagen3
	resp, err := genkit.Generate(ctx, g,
		ai.WithModelName("googleai/imagen-3.0-generate-002"),
		ai.WithPrompt(prompt),
		ai.WithConfig(&genai.GenerateImagesConfig{
			NumberOfImages:    1,
			AspectRatio:       "1:1",
			SafetyFilterLevel: genai.SafetyFilterLevelBlockLowAndAbove,
			PersonGeneration:  genai.PersonGenerationAllowAll,
			OutputMIMEType:    "image/png",
		}),
	)
	if err != nil {
		log.Printf("Failed to generate image: %v", err)
		return client.ActionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to generate image: %v", err),
		}, err
	}

	// Extract image data from response
	var imageData string
	for _, part := range resp.Message.Content {
		if part.IsMedia() {
			// The Text field contains the image data (base64 encoded data URI)
			imageData = part.Text
			break
		}
	}

	if imageData == "" {
		return client.ActionResponse{
			Success: false,
			Error:   "No image generated in response",
		}, fmt.Errorf("no image generated in response")
	}

	return client.ActionResponse{
		Success: true,
		Result: map[string]interface{}{
			"imageUrl": imageData,
			"dishName": input.DishName,
		},
	}, nil
}

// evaluateDish analyzes and scores a dish image
func evaluateDish(ctx context.Context, g *genkit.Genkit, input client.ActionRequest) (client.ActionResponse, error) {
	systemPrompt := `You are a professional food critic and cooking competition judge.
Evaluate dishes based on presentation, creativity, and apparent quality.
Provide constructive feedback and fair scoring.`

	prompt := fmt.Sprintf(`Evaluate this dish for a cooking battle:
Dish Name: %s
Description: %s
Image URL: %s

Please provide:
1. Presentation score (1-10)
2. Creativity score (1-10)
3. Apparent taste/quality score (1-10)
4. Overall score (1-10)
5. Detailed feedback and suggestions for improvement

Consider the constraints: %v`,
		input.DishName, input.Description, input.ImageUrl, input.Constraints)

	resp, err := genkit.Generate(ctx, g,
		ai.WithSystem(systemPrompt),
		ai.WithPrompt(prompt),
	)
	if err != nil {
		log.Printf("Failed to evaluate dish: %v", err)
		return client.ActionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to evaluate dish: %v", err),
		}, err
	}

	return client.ActionResponse{
		Success: true,
		Result: map[string]interface{}{
			"evaluation": resp.Text(),
			"dishName":   input.DishName,
		},
	}, nil
}
