package tools

// CheckIngredientInput represents the input for ingredient stock check tool
type CheckIngredientInput struct {
	Ingredients []string `json:"ingredients" jsonschema_description:"List of ingredients to check availability for"`
}