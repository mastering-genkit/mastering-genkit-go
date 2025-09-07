package client

// RecipeRequest represents the input for streaming recipe generation flow
type RecipeRequest struct {
	Ingredients []string `json:"ingredients"`
}
