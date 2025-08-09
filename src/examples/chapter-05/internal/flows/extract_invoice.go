package flows

import (
	"context"
	"examples/chapter-05/internal/schemas"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewExtractInvoiceDataFlow demonstrates GenerateData for structured extraction
func NewExtractInvoiceDataFlow(g *genkit.Genkit) *core.Flow[string, schemas.InvoiceData, struct{}] {
	return genkit.DefineFlow(g, "extractInvoiceData",
		func(ctx context.Context, documentText string) (schemas.InvoiceData, error) {
			invoicePtr, _, err := genkit.GenerateData[schemas.InvoiceData](ctx, g,
				ai.WithPrompt("Extract invoice data from the following document:\n"+documentText))

			if err != nil {
				return schemas.InvoiceData{}, fmt.Errorf("failed to extract invoice data: %w", err)
			}

			// Always validate extracted data
			invoice := *invoicePtr
			if err := ValidateInvoice(invoice); err != nil {
				return schemas.InvoiceData{}, fmt.Errorf("invalid invoice data: %w", err)
			}

			return invoice, nil
		})
}

// ValidateInvoice validates invoice data
func ValidateInvoice(invoice schemas.InvoiceData) error {
	if invoice.Amount < 0 {
		return fmt.Errorf("invoice amount cannot be negative")
	}
	if invoice.InvoiceNumber == "" {
		return fmt.Errorf("invoice number is required")
	}
	if invoice.DueDate == "" {
		return fmt.Errorf("due date is required")
	}
	return nil
}
