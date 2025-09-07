package domain

// BattleResult represents the outcome of a cooking battle
type BattleResult struct {
	PlayerDish   CookedDish `json:"playerDish"`
	OpponentDish CookedDish `json:"opponentDish"`
	Winner       string     `json:"winner"`
	Evaluation   string     `json:"evaluation"`
}

// CookedDish represents a completed dish with image
type CookedDish struct {
	RecipeName  string `json:"recipeName"`
	ImageURL    string `json:"imageUrl,omitempty"`
	ImageBase64 string `json:"imageBase64,omitempty"`
	Score       int    `json:"score"`
	Comments    string `json:"comments"`
}