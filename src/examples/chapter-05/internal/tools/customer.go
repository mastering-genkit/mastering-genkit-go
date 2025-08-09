package tools

import (
	"examples/chapter-05/internal/schemas"
	"fmt"
	"strings"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// NewValidateCustomer creates a tool for customer validation with structured output
func NewValidateCustomer(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, "validate_customer",
		"Validate customer and return detailed status",
		func(ctx *ai.ToolContext, input schemas.CustomerValidationInput) (schemas.CustomerValidationResult, error) {
			result := schemas.CustomerValidationResult{}

			// Age validation
			if input.Age < 18 {
				return result, fmt.Errorf("customer must be 18 or older")
			}

			// Email validation
			if !strings.Contains(input.Email, "@") {
				return result, fmt.Errorf("invalid email format")
			}

			// Name validation
			if len(input.Name) < 2 {
				return result, fmt.Errorf("name too short")
			}

			// Return structured data
			result.CustomerID = fmt.Sprintf("CUST-%d", time.Now().Unix())
			result.IsValid = true
			result.CreditScore = calculateCreditScore(input.Age, input.CreditLimit)
			result.AccountTier = determineAccountTier(result.CreditScore)

			// Add restrictions if needed
			if input.CreditLimit < 1000 {
				result.Restrictions = []string{"low_credit_limit"}
			}

			return result, nil
		})
}

// Business logic functions (testable in isolation)

func calculateCreditScore(age int, creditLimit float64) int {
	baseScore := 600
	
	// Age bonus
	if age > 25 {
		baseScore += 50
	}
	if age > 40 {
		baseScore += 50
	}
	
	// Credit limit bonus
	if creditLimit > 5000 {
		baseScore += 100
	}
	
	return baseScore
}

func determineAccountTier(creditScore int) string {
	switch {
	case creditScore >= 750:
		return "Premium"
	case creditScore >= 650:
		return "Standard"
	default:
		return "Basic"
	}
}