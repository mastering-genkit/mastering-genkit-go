package generation

// CharacterRequest represents the input for character generation
type CharacterRequest struct {
	Description string `json:"description"` // "blue-haired wizard"
	Style       string `json:"style"`       // "anime", "realistic", "cartoon"
}