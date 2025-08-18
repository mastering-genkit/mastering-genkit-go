package nested

// Recipe represents the complete recipe with all details
type Recipe struct {
	Name        string        `json:"name"`
	Ingredients []Ingredient  `json:"ingredients"`
	Steps       []CookingStep `json:"steps"`
	Nutrition   Nutrition     `json:"nutrition"`
	PrepTime    int           `json:"prep_time"`  // minutes
	CookTime    int           `json:"cook_time"`  // minutes
	Difficulty  string        `json:"difficulty"` // easy, medium, hard
}

// Ingredient represents a single ingredient with measurements
type Ingredient struct {
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Unit     string  `json:"unit"`
	Optional bool    `json:"optional"`
}

// CookingStep represents a single step in the cooking process
type CookingStep struct {
	Number      int    `json:"number"`
	Instruction string `json:"instruction"`
	Duration    int    `json:"duration"` // minutes
	Tips        string `json:"tips,omitempty"`
}

// Nutrition represents nutritional information
type Nutrition struct {
	Calories float64 `json:"calories"`
	Protein  float64 `json:"protein"` // grams
	Carbs    float64 `json:"carbs"`   // grams
	Fat      float64 `json:"fat"`     // grams
	Fiber    float64 `json:"fiber"`   // grams
}
