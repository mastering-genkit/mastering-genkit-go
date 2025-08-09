package flows

import (
	"examples/chapter-05/internal/schemas"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInvoice(t *testing.T) {
	// Test valid invoice
	validInvoice := schemas.InvoiceData{
		InvoiceNumber: "INV-001",
		Amount:        1500.00,
		DueDate:       "2024-02-15",
		CustomerName:  "ABC Corp",
	}
	assert.NoError(t, ValidateInvoice(validInvoice))

	// Test zero amount is valid
	zeroAmountInvoice := schemas.InvoiceData{
		InvoiceNumber: "INV-004",
		Amount:        0,
		DueDate:       "2024-02-15",
	}
	assert.NoError(t, ValidateInvoice(zeroAmountInvoice))

	// Test negative amount
	negativeInvoice := schemas.InvoiceData{
		InvoiceNumber: "INV-002",
		Amount:        -100.00,
		DueDate:       "2024-02-15",
	}
	err := ValidateInvoice(negativeInvoice)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invoice amount cannot be negative")

	// Test missing invoice number
	missingNumberInvoice := schemas.InvoiceData{
		Amount:  1500.00,
		DueDate: "2024-02-15",
	}
	err = ValidateInvoice(missingNumberInvoice)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invoice number is required")

	// Test missing due date
	missingDateInvoice := schemas.InvoiceData{
		InvoiceNumber: "INV-003",
		Amount:        1500.00,
	}
	err = ValidateInvoice(missingDateInvoice)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "due date is required")
}