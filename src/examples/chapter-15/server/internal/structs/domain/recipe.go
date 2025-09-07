package domain

// Recipe represents a cooking recipe in the system
type Recipe struct {
	ID          string       `json:"id" firestore:"-"`
	Name        string       `json:"name" firestore:"name"`
	Category    string       `json:"category" firestore:"category"`
	Difficulty  string       `json:"difficulty" firestore:"difficulty"`
	PrepTime    int          `json:"prepTime" firestore:"prepTime"`
	CookTime    int          `json:"cookTime" firestore:"cookTime"`
	Ingredients []Ingredient `json:"ingredients" firestore:"ingredients"`
	Nutrition   Nutrition    `json:"nutrition" firestore:"nutrition"`
	Tags        []string     `json:"tags" firestore:"tags"`
}

// Ingredient represents an ingredient in a recipe
type Ingredient struct {
	Name   string `json:"name" firestore:"name"`
	Amount int    `json:"amount" firestore:"amount"`
	Unit   string `json:"unit" firestore:"unit"`
}

// Nutrition represents nutritional information per serving
type Nutrition struct {
	Calories int `json:"calories" firestore:"calories"`
	Protein  int `json:"protein" firestore:"protein"`
	Carbs    int `json:"carbs" firestore:"carbs"`
	Fat      int `json:"fat" firestore:"fat"`
}