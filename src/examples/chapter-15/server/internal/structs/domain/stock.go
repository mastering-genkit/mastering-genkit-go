package domain

// IngredientStock represents available ingredients in stock
type IngredientStock struct {
	ID         string `json:"id" firestore:"-"`
	Ingredient string `json:"ingredient" firestore:"ingredient"`
	Quantity   int    `json:"quantity" firestore:"quantity"`
	Unit       string `json:"unit" firestore:"unit"`
	Available  bool   `json:"available" firestore:"available"`
}