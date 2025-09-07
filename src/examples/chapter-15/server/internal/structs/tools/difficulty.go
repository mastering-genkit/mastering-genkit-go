package tools

// EstimateDifficultyInput represents input for cooking difficulty estimation
type EstimateDifficultyInput struct {
	Ingredients    []string `json:"ingredients" jsonschema_description:"List of ingredients being used"`
	CookingSteps   int      `json:"cookingSteps" jsonschema_description:"Number of cooking steps in the recipe"`
	CookingMethods []string `json:"cookingMethods" jsonschema_description:"List of cooking methods (boil, fry, grill, etc.)"`
}

// DifficultyEstimate represents the difficulty analysis result
type DifficultyEstimate struct {
	Level             string   `json:"level"`             // Easy, Medium, Hard
	Score             int      `json:"score"`             // 1-10
	Reasoning         string   `json:"reasoning"`         // Why this difficulty level
	TimeEstimate      int      `json:"timeEstimate"`      // Estimated cooking time in minutes
	SkillsRequired    []string `json:"skillsRequired"`    // Required cooking skills
	EquipmentRequired []string `json:"equipmentRequired"` // Required cooking equipment
	Tips              string   `json:"tips"`              // Helpful tips for this difficulty level
}
