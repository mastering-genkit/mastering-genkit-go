package structs

// RecommendationOutput represents the final server response for the recommendation flow.
// This is the data structure returned to clients after orchestrating multiple flows
// and AI-generated content.
//
// Example JSON response:
//
//	{
//	  "instrument": "Acoustic Guitar",
//	  "why": "The acoustic guitar is perfect for jazz beginners because...",
//	  "starter_items": ["Guitar picks", "Tuner", "Gig bag", "Chord chart"]
//	}
type RecommendationOutput struct {
	Instrument   string   `json:"instrument"`    // The recommended instrument name
	Why          string   `json:"why"`           // AI-generated explanation (2-3 sentences)
	StarterItems []string `json:"starter_items"` // Essential accessories and items to get started
}
