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

// NewCheckIngredientStock creates a tool that checks ingredient availability in stock
func NewCheckIngredientStock(genkitClient *genkit.Genkit, firestoreClient *firestore.Client) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"checkIngredientStock",
		"Check the availability and quantity of ingredients in stock",
		func(ctx *ai.ToolContext, input toolstructs.CheckIngredientInput) ([]domain.IngredientStock, error) {
			log.Printf("Tool 'checkIngredientStock' called with ingredients: %v", input.Ingredients)

			// Query ingredient_stock collection
			iter := firestoreClient.Collection("ingredient_stock").Documents(context.Background())
			var stockItems []domain.IngredientStock

			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Printf("Failed to iterate documents: %v", err)
					return nil, fmt.Errorf("failed to iterate documents: %w", err)
				}

				var stock domain.IngredientStock
				if err := doc.DataTo(&stock); err != nil {
					log.Printf("Failed to parse stock document: %v", err)
					continue // Skip malformed documents
				}
				stock.ID = doc.Ref.ID

				// If specific ingredients are requested, filter
				if len(input.Ingredients) > 0 {
					found := false
					for _, requested := range input.Ingredients {
						if strings.Contains(strings.ToLower(stock.Ingredient), strings.ToLower(requested)) {
							found = true
							break
						}
					}
					if !found {
						continue
					}
				}

				stockItems = append(stockItems, stock)
			}

			// Add "not found" entries for requested ingredients not in stock
			for _, requested := range input.Ingredients {
				found := false
				for _, stock := range stockItems {
					if strings.Contains(strings.ToLower(stock.Ingredient), strings.ToLower(requested)) {
						found = true
						break
					}
				}
				if !found {
					stockItems = append(stockItems, domain.IngredientStock{
						Ingredient: requested,
						Quantity:   0,
						Unit:       "unknown",
						Available:  false,
					})
					log.Printf("Ingredient '%s' not found in stock", requested)
				}
			}

			log.Printf("Checked %d ingredients, %d items returned", len(input.Ingredients), len(stockItems))
			return stockItems, nil
		})
}
