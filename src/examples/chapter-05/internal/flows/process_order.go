package flows

import (
	"context"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewProcessCompleteOrderFlow demonstrates multi-tool orchestration
func NewProcessCompleteOrderFlow(g *genkit.Genkit, toolRefs []ai.ToolRef) *core.Flow[string, string, struct{}] {
	return genkit.DefineFlow(g, "processCompleteOrder",
		func(ctx context.Context, request string) (string, error) {
			response, err := genkit.Generate(ctx, g,
				ai.WithSystem(`You are an order processing assistant. You MUST use ALL relevant tools to complete the order processing:

1. ALWAYS use validate_customer first when customer information is provided
2. ALWAYS use check_inventory when product SKUs are mentioned
3. ALWAYS use create_order after validation succeeds

Available tools:
- validate_customer: requires name (string), email (string), age (number), credit_limit (number)
- check_inventory: requires items (array of objects with 'sku' and 'quantity' fields)
- create_order: requires customer_id (string), items (array of strings like ["SKU-001", "SKU-002"]), total_amount (number, not 'total')

IMPORTANT: For create_order, use 'total_amount' not 'total', and items must be an array of strings.

You must call multiple tools in sequence to complete the order.`),
				ai.WithPrompt("Process this order request completely: %s", request),
				ai.WithTools(toolRefs...),
				ai.WithConfig(map[string]interface{}{
					"temperature": 0.2,
				}),
			)

			if err != nil {
				return "", err
			}

			return response.Text(), nil
		})
}