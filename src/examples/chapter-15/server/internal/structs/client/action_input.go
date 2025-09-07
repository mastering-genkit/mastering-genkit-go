package client

// ActionRequest represents the input for non-streaming action flows
type ActionRequest struct {
	// Action to perform: "generate_recipe", "create_image", "evaluate"
	Action string `json:"action"`

	// For generate_recipe action
	Ingredients []string               `json:"ingredients,omitempty"`
	Constraints map[string]interface{} `json:"constraints,omitempty"`

	// For create_image action
	DishName    string `json:"dishName,omitempty"`
	Description string `json:"description,omitempty"`

	// For evaluate action
	ImageUrl string `json:"imageUrl,omitempty"`
}