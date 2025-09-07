package tools

import (
	"context"
	"fmt"
	"log"
	"strings"

	"mastering-genkit-go/example/chapter-15/internal/structs/domain"
	toolstructs "mastering-genkit-go/example/chapter-15/internal/structs/tools"

	"cloud.google.com/go/firestore"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"google.golang.org/api/iterator"
)

// NewSearchRecipeDatabase creates a tool that searches for recipes in Firestore
func NewSearchRecipeDatabase(genkitClient *genkit.Genkit, firestoreClient *firestore.Client) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"searchRecipe",
		"Search for recipes in the database that match given ingredients, tags, or category",
		func(ctx *ai.ToolContext, input toolstructs.SearchRecipeInput) ([]domain.Recipe, error) {
			log.Printf("Tool 'searchRecipeDatabase' called with ingredients: %v, tags: %v, category: %s",
				input.Ingredients, input.Tags, input.Category)

			// Start with recipes collection as a Query
			query := firestoreClient.Collection("recipes").Query

			// Set max results (default to 5)
			maxResults := input.MaxResults
			if maxResults <= 0 {
				maxResults = 5
			}

			// Apply category filter if specified
			if input.Category != "" {
				query = query.Where("category", "==", input.Category)
			}

			// Get all recipes and filter by ingredients
			iter := query.Documents(context.Background())
			var recipes []domain.Recipe

			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Printf("Failed to iterate documents: %v", err)
					return nil, fmt.Errorf("failed to iterate documents: %w", err)
				}

				var recipe domain.Recipe
				if err := doc.DataTo(&recipe); err != nil {
					log.Printf("Failed to parse recipe document: %v", err)
					continue // Skip malformed documents
				}
				recipe.ID = doc.Ref.ID

				// Check if recipe contains any of the requested ingredients
				if len(input.Ingredients) > 0 {
					if !containsIngredients(recipe.Ingredients, input.Ingredients) {
						continue
					}
				}

				// Check if recipe has any of the requested tags
				if len(input.Tags) > 0 {
					if !containsTags(recipe.Tags, input.Tags) {
						continue
					}
				}

				recipes = append(recipes, recipe)

				// Limit results
				if len(recipes) >= maxResults {
					break
				}
			}

			log.Printf("Found %d recipes matching criteria", len(recipes))
			return recipes, nil
		})
}

// containsIngredients checks if recipe contains any of the requested ingredients
func containsIngredients(recipeIngredients []domain.Ingredient, requestedIngredients []string) bool {
	for _, recipeIng := range recipeIngredients {
		for _, requested := range requestedIngredients {
			if strings.Contains(strings.ToLower(recipeIng.Name), strings.ToLower(requested)) {
				return true
			}
		}
	}
	return false
}

// containsTags checks if recipe has any of the requested tags
func containsTags(recipeTags []string, requestedTags []string) bool {
	for _, recipeTag := range recipeTags {
		for _, requested := range requestedTags {
			if strings.EqualFold(recipeTag, requested) {
				return true
			}
		}
	}
	return false
}
