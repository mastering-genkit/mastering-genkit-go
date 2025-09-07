package client

// RecipeResponse represents a streaming response chunk for recipe generation
type RecipeResponse struct {
	Type    string `json:"type"` // "content", "done", "error"
	Content string `json:"content,omitempty"`
	Error   string `json:"error,omitempty"`
}
