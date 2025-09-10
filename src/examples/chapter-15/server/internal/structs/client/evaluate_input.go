package client

// EvaluateRequest represents the input for evaluation flow
type EvaluateRequest struct {
	DishName    string `json:"dishName"`
	Description string `json:"description"`
}
