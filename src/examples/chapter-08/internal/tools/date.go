package tools

import (
	"fmt"
	"log"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// DateRequest represents the input structure for the date tool
type DateRequest struct {
	Format string `json:"format" jsonschema_description:"Date format to use"` // Optional format string
}

// NewGetCurrentDate creates a tool that returns the current date in a specified format.
// If no format is provided, it defaults to RFC3339 format.
func NewGetCurrentDate(genkitClient *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"getCurrentDate",
		"Gets the current date and time in a specified format. Accepts a string with an optional 'format' field. Common formats: 'RFC3339' (default), 'Kitchen', 'Stamp', 'DateTime', or custom Go time format like '2006-01-02 15:04:05'.",
		func(ctx *ai.ToolContext, input DateRequest) (string, error) {
			log.Printf("Tool 'getCurrentDate' called with input: %s", input)

			// Set default format if none provided
			if input.Format == "" {
				input.Format = time.RFC3339
			}

			now := time.Now()
			var formattedDate string

			// Handle common named formats
			switch input.Format {
			case "RFC3339":
				formattedDate = now.Format(time.RFC3339)
			case "Kitchen":
				formattedDate = now.Format(time.Kitchen)
			case "Stamp":
				formattedDate = now.Format(time.Stamp)
			case "DateTime":
				formattedDate = now.Format(time.DateTime)
			case "DateOnly":
				formattedDate = now.Format(time.DateOnly)
			case "TimeOnly":
				formattedDate = now.Format(time.TimeOnly)
			case "RFC822":
				formattedDate = now.Format(time.RFC822)
			case "RFC1123":
				formattedDate = now.Format(time.RFC1123)
			default:
				// Treat as custom format
				formattedDate = now.Format(input.Format)
			}

			result := fmt.Sprintf("Current date and time: %s (format: %s)", formattedDate, input.Format)
			log.Printf("Returning formatted date: %s", result)

			return result, nil
		})
}
