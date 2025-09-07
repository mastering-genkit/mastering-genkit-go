package tools

import (
	"context"
	"log"
	"sort"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"

	toolstructs "mastering-genkit-go/example/chapter-15/internal/structs/tools"
)

// NewCheckIngredientCompatibility creates a tool that checks ingredient combination compatibility
func NewCheckIngredientCompatibility(genkitClient *genkit.Genkit, firestoreClient *firestore.Client) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"checkIngredientCompatibility",
		"Check how well ingredients work together and get compatibility insights",
		func(ctx *ai.ToolContext, input toolstructs.CheckCompatibilityInput) (*toolstructs.CompatibilityResult, error) {
			log.Printf("Checking compatibility for ingredients: %v", input.Ingredients)

			// Query Firestore for ingredient combinations
			bestMatch, err := findBestCompatibilityMatch(ctx, firestoreClient, input.Ingredients)
			if err != nil {
				log.Printf("Error querying compatibility: %v", err)
				// Return default compatibility result
				return createDefaultCompatibilityResult(input.Ingredients), nil
			}

			if bestMatch != nil {
				return bestMatch, nil
			}

			// No direct match found, return estimated compatibility
			return estimateCompatibility(input.Ingredients), nil
		},
	)
}

// findBestCompatibilityMatch searches for the best ingredient combination in Firestore
func findBestCompatibilityMatch(ctx context.Context, client *firestore.Client, ingredients []string) (*toolstructs.CompatibilityResult, error) {
	// Normalize ingredients for matching (lowercase, sorted)
	normalizedIngredients := make([]string, len(ingredients))
	for i, ingredient := range ingredients {
		normalizedIngredients[i] = strings.ToLower(strings.TrimSpace(ingredient))
	}
	sort.Strings(normalizedIngredients)

	// Query all combinations from Firestore
	iter := client.Collection("ingredient_combinations").Documents(ctx)
	defer iter.Stop()

	var bestMatch *toolstructs.CompatibilityResult
	highestScore := 0

	for {
		doc, err := iter.Next()
		if err != nil {
			break // End of documents
		}

		data := doc.Data()

		// Extract ingredients from Firestore document
		dbIngredients, ok := data["ingredients"].([]interface{})
		if !ok {
			continue
		}

		// Convert to string slice and normalize
		var dbIngredientsStr []string
		for _, ing := range dbIngredients {
			if ingStr, ok := ing.(string); ok {
				dbIngredientsStr = append(dbIngredientsStr, strings.ToLower(strings.TrimSpace(ingStr)))
			}
		}
		sort.Strings(dbIngredientsStr)

		// Check for exact match or subset match
		matchCount := countMatchingIngredients(normalizedIngredients, dbIngredientsStr)

		if matchCount > 0 {
			score, _ := data["compatibility_score"].(int64)
			if int(score) > highestScore {
				highestScore = int(score)
				bestMatch = &toolstructs.CompatibilityResult{
					Ingredients:        ingredients,
					CompatibilityScore: int(score),
					FlavorProfile:      getStringField(data, "flavor_profile"),
					CuisineStyle:       getStringField(data, "cuisine_style"),
					Tips:               getStringField(data, "tips"),
					DifficultyBonus:    int(getIntField(data, "difficulty_bonus")),
					OverallRating:      getRatingFromScore(int(score)),
				}
			}
		}
	}

	return bestMatch, nil
}

// countMatchingIngredients counts how many ingredients match between two lists
func countMatchingIngredients(list1, list2 []string) int {
	count := 0
	for _, ing1 := range list1 {
		for _, ing2 := range list2 {
			if ing1 == ing2 {
				count++
			}
		}
	}
	return count
}

// estimateCompatibility provides a fallback compatibility estimate
func estimateCompatibility(ingredients []string) *toolstructs.CompatibilityResult {
	// Simple heuristic based on ingredient count and common pairings
	score := 6 // Base score

	if len(ingredients) <= 4 {
		score += 1 // Bonus for simple combinations
	}

	profile := "balanced"
	style := "fusion"
	tips := "Experiment with cooking methods to bring out the best in these ingredients"

	return &toolstructs.CompatibilityResult{
		Ingredients:        ingredients,
		CompatibilityScore: score,
		FlavorProfile:      profile,
		CuisineStyle:       style,
		Tips:               tips,
		DifficultyBonus:    0,
		OverallRating:      getRatingFromScore(score),
	}
}

// createDefaultCompatibilityResult creates a safe fallback result
func createDefaultCompatibilityResult(ingredients []string) *toolstructs.CompatibilityResult {
	return &toolstructs.CompatibilityResult{
		Ingredients:        ingredients,
		CompatibilityScore: 5,
		FlavorProfile:      "unknown",
		CuisineStyle:       "fusion",
		Tips:               "Creative combination - experiment and discover!",
		DifficultyBonus:    0,
		OverallRating:      "Average",
	}
}

// Helper functions for Firestore data extraction
func getStringField(data map[string]interface{}, field string) string {
	if value, ok := data[field].(string); ok {
		return value
	}
	return ""
}

func getIntField(data map[string]interface{}, field string) int64 {
	if value, ok := data[field].(int64); ok {
		return value
	}
	return 0
}

// getRatingFromScore converts numeric score to descriptive rating
func getRatingFromScore(score int) string {
	switch {
	case score >= 9:
		return "Perfect Match"
	case score >= 7:
		return "Great Combination"
	case score >= 5:
		return "Good Pairing"
	case score >= 3:
		return "Acceptable"
	default:
		return "Challenging"
	}
}
