package tools

import (
	"examples/chapter-05/internal/schemas"
	"fmt"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// NewCreateOrder creates a tool for order creation with structured output
func NewCreateOrder(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, "create_order",
		"Create order and return order details",
		func(ctx *ai.ToolContext, input schemas.OrderCreationInput) (schemas.OrderResult, error) {
			result := schemas.OrderResult{}

			// Validate order amount
			if err := validateOrderAmount(input.TotalAmount); err != nil {
				return result, err
			}

			// Validate items
			if len(input.Items) == 0 {
				return result, fmt.Errorf("order must contain at least one item")
			}

			// Create order with structured response
			result.OrderID = fmt.Sprintf("ORD-%d", time.Now().Unix())
			result.Status = "confirmed"
			result.EstimatedDelivery = calculateDeliveryDate(len(input.Items))
			result.TrackingNumber = fmt.Sprintf("TRK%d", time.Now().Unix())
			result.TotalAmount = input.TotalAmount
			result.PaymentMethod = selectPaymentMethod(input.TotalAmount)

			return result, nil
		})
}

// Business logic functions (testable in isolation)

// ValidateOrderAmount validates the order amount against business rules
func ValidateOrderAmount(amount float64) error {
	if amount > 10000 {
		return fmt.Errorf("order amount exceeds maximum limit")
	}
	if amount < 0 {
		return fmt.Errorf("order amount cannot be negative")
	}
	return nil
}

// Export for testing from chapter
func validateOrderAmount(amount float64) error {
	return ValidateOrderAmount(amount)
}

func calculateDeliveryDate(itemCount int) string {
	// Add more days for larger orders
	daysToAdd := 3
	if itemCount > 5 {
		daysToAdd = 5
	}
	if itemCount > 10 {
		daysToAdd = 7
	}
	
	return time.Now().AddDate(0, 0, daysToAdd).Format("2006-01-02")
}

func selectPaymentMethod(amount float64) string {
	// Select payment method based on amount
	if amount > 1000 {
		return "credit_card"
	}
	return "debit_card"
}