package image

// ProductInfo represents the structured output of product image analysis
type ProductInfo struct {
	Category       string   `json:"category"`
	Features       []string `json:"features"`
	Colors         []string `json:"colors"`
	EstimatedPrice string   `json:"estimated_price"` // price range
	TargetAudience string   `json:"target_audience"`
	Description    string   `json:"description"`
}