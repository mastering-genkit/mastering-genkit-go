package tools

import (
	"examples/chapter-05/internal/schemas"
	"fmt"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// NewCheckInventory creates a tool for inventory checking with structured output
func NewCheckInventory(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, "check_inventory",
		"Check inventory availability for multiple items and return detailed stock information",
		func(ctx *ai.ToolContext, input schemas.InventoryCheckInput) (schemas.InventoryCheckResult, error) {
			result := schemas.InventoryCheckResult{
				ItemsInStock: make(map[string]int),
			}

			// Validation
			if len(input.Items) == 0 {
				return result, fmt.Errorf("no items specified")
			}

			// Check each item's availability
			allAvailable := true
			skus := []string{}

			for _, item := range input.Items {
				if item.Quantity <= 0 {
					return result, fmt.Errorf("quantity must be positive for SKU %s", item.SKU)
				}
				if item.Quantity > 100 {
					return result, fmt.Errorf("quantity exceeds maximum order limit for SKU %s", item.SKU)
				}

				// Get current stock for the item
				currentStock := getSimulatedStock(item.SKU)
				result.ItemsInStock[item.SKU] = currentStock

				// Check if we have enough stock
				if currentStock < item.Quantity {
					allAvailable = false
				}

				skus = append(skus, item.SKU)
			}

			result.Available = allAvailable
			result.ReservedUntil = time.Now().Add(15 * time.Minute).Format(time.RFC3339)
			result.Warehouse = selectWarehouse(skus)

			return result, nil
		})
}

// Business logic functions (testable in isolation)
func getSimulatedStock(item string) int {
	// Simulate different stock levels based on SKU
	stockMap := map[string]int{
		"SKU-001":  100,
		"SKU-002":  100,
		"SKU-003":  25,
		"SKU-004":  0,
		"SKU-005":  300,
		"laptop":   50,
		"mouse":    200,
		"keyboard": 150,
		"monitor":  75,
		"cable":    500,
	}

	if stock, exists := stockMap[item]; exists {
		return stock
	}
	return 100 // Default stock
}

func selectWarehouse(items []string) string {
	// Simple logic to select warehouse
	// In production, this would consider location, stock levels, etc.
	if len(items) > 3 {
		return "CENTRAL"
	}
	return "EAST-1"
}
