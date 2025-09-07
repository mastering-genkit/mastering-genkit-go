package tools

// CalculateNutritionInput represents the input for nutrition calculation tool
type CalculateNutritionInput struct {
	RecipeID    string   `json:"recipeId,omitempty" jsonschema_description:"Recipe ID to calculate nutrition for"`
	Ingredients []string `json:"ingredients,omitempty" jsonschema_description:"List of ingredients to calculate nutrition"`
	Servings    int      `json:"servings,omitempty" jsonschema_description:"Number of servings (default: 1)"`
}