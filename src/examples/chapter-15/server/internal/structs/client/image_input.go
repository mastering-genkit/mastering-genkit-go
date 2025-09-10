package client

// ImageRequest represents the input for image generation flow
type ImageRequest struct {
	DishName    string `json:"dishName"`
	Description string `json:"description"`
}
