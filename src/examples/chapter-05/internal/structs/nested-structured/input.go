package nested

// RecipeRequest represents the input for recipe generation
type RecipeRequest struct {
	DishName            string   `json:"dish_name"`
	Servings            int      `json:"servings"`
	DietaryRestrictions []string `json:"dietary_restrictions"`
}