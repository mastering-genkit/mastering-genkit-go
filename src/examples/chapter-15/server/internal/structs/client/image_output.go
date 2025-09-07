package client

// ImageResponse represents the output for image generation flow
type ImageResponse struct {
	Success  bool   `json:"success"`
	ImageUrl string `json:"imageUrl"`
	DishName string `json:"dishName"`
	Error    string `json:"error"`
}
