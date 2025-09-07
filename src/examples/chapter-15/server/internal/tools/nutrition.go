package tools

import (
	"context"
	"fmt"
	"log"

	"mastering-genkit-go/example/chapter-15/internal/structs/domain"
	toolstructs "mastering-genkit-go/example/chapter-15/internal/structs/tools"

	"cloud.google.com/go/firestore"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// NutritionResult represents the calculated nutrition information
type NutritionResult struct {
	TotalCalories int              `json:"totalCalories"`
	TotalProtein  int              `json:"totalProtein"`
	TotalCarbs    int              `json:"totalCarbs"`
	TotalFat      int              `json:"totalFat"`
	PerServing    domain.Nutrition `json:"perServing"`
	Servings      int              `json:"servings"`
	RecipeName    string           `json:"recipeName,omitempty"`
}

// NewCalculateNutrition creates a tool that calculates nutrition information
func NewCalculateNutrition(genkitClient *genkit.Genkit, firestoreClient *firestore.Client) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"calculateNutrition",
		"Calculate nutritional information for a recipe or list of ingredients",
		func(ctx *ai.ToolContext, input toolstructs.CalculateNutritionInput) (*NutritionResult, error) {
			log.Printf("Tool 'calculateNutrition' called with recipeID: %s, ingredients: %v, servings: %d",
				input.RecipeID, input.Ingredients, input.Servings)

			// Default to 1 serving if not specified
			servings := input.Servings
			if servings <= 0 {
				servings = 1
			}

			var result NutritionResult
			result.Servings = servings

			// If recipe ID is provided, get nutrition from recipe
			if input.RecipeID != "" {
				doc, err := firestoreClient.Collection("recipes").Doc(input.RecipeID).Get(context.Background())
				if err != nil {
					log.Printf("Failed to get recipe document: %v", err)
					return nil, fmt.Errorf("failed to get recipe %s: %w", input.RecipeID, err)
				}

				var recipe domain.Recipe
				if err := doc.DataTo(&recipe); err != nil {
					log.Printf("Failed to parse recipe document: %v", err)
					return nil, fmt.Errorf("failed to parse recipe data: %w", err)
				}

				// Calculate total nutrition (recipe nutrition is per serving)
				result.RecipeName = recipe.Name
				result.TotalCalories = recipe.Nutrition.Calories * servings
				result.TotalProtein = recipe.Nutrition.Protein * servings
				result.TotalCarbs = recipe.Nutrition.Carbs * servings
				result.TotalFat = recipe.Nutrition.Fat * servings
				result.PerServing = recipe.Nutrition

				log.Printf("Calculated nutrition for recipe '%s': %d calories total", recipe.Name, result.TotalCalories)
				return &result, nil
			}

			// If ingredients are provided, calculate based on ingredients
			if len(input.Ingredients) > 0 {
				// For simplicity, use estimated values per ingredient
				// In a real application, this would query a nutrition database
				estimatedNutrition := estimateNutritionFromIngredients(input.Ingredients)

				result.TotalCalories = estimatedNutrition.Calories * servings
				result.TotalProtein = estimatedNutrition.Protein * servings
				result.TotalCarbs = estimatedNutrition.Carbs * servings
				result.TotalFat = estimatedNutrition.Fat * servings
				result.PerServing = estimatedNutrition

				log.Printf("Estimated nutrition for %d ingredients: %d calories per serving",
					len(input.Ingredients), estimatedNutrition.Calories)
				return &result, nil
			}

			return nil, fmt.Errorf("either recipeId or ingredients must be provided")
		})
}

// estimateNutritionFromIngredients provides rough nutrition estimates
// In a real application, this would query a nutrition database
func estimateNutritionFromIngredients(ingredients []string) domain.Nutrition {
	// Base values per ingredient (very simplified)
	baseCalories := 50
	baseProtein := 3
	baseCarbs := 8
	baseFat := 2

	// Calculate based on number of ingredients
	count := len(ingredients)

	// Check for high-protein ingredients
	proteinBoost := 0
	for _, ing := range ingredients {
		if contains(ing, []string{"chicken", "beef", "fish", "egg", "tofu"}) {
			proteinBoost += 10
		}
	}

	// Check for carb-heavy ingredients
	carbBoost := 0
	for _, ing := range ingredients {
		if contains(ing, []string{"rice", "pasta", "bread", "potato"}) {
			carbBoost += 15
		}
	}

	return domain.Nutrition{
		Calories: baseCalories * count,
		Protein:  baseProtein*count + proteinBoost,
		Carbs:    baseCarbs*count + carbBoost,
		Fat:      baseFat * count,
	}
}

// contains checks if text contains any of the keywords
func contains(text string, keywords []string) bool {
	for _, keyword := range keywords {
		if len(text) >= len(keyword) && text[:len(keyword)] == keyword {
			return true
		}
	}
	return false
}
