package tools

// CheckCompatibilityInput represents input for ingredient compatibility checking
type CheckCompatibilityInput struct {
	Ingredients []string `json:"ingredients" jsonschema_description:"List of ingredients to check compatibility for"`
}

// CompatibilityResult represents the compatibility analysis result
type CompatibilityResult struct {
	Ingredients        []string `json:"ingredients"`
	CompatibilityScore int      `json:"compatibilityScore"`
	FlavorProfile      string   `json:"flavorProfile"`
	CuisineStyle       string   `json:"cuisineStyle"`
	Tips               string   `json:"tips"`
	DifficultyBonus    int      `json:"difficultyBonus"`
	OverallRating      string   `json:"overallRating"`
}
