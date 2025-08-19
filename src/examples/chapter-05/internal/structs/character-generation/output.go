package generation

// CharacterResult represents the complete result with image and metadata
type CharacterResult struct {
	ImageData string    `json:"image_data"` // base64 encoded image (data URI format)
	Character Character `json:"character"`  // Character details
}

// Character represents the detailed character information
type Character struct {
	Name        string   `json:"name"`        // AI-generated name
	Description string   `json:"description"` // Visual description
	Traits      []string `json:"traits"`      // Character traits
	Story       string   `json:"story"`       // Brief backstory
}