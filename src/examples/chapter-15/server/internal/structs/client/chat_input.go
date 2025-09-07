package client

// ChatRequest represents the input for the cookingBattleChat flow
type ChatRequest struct {
	Messages    []Message         `json:"messages"`
	Ingredients []string          `json:"ingredients,omitempty"`
	Constraints map[string]string `json:"constraints,omitempty"`
	Context     string            `json:"context,omitempty"`
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"` // "user" or "assistant"
	Content string `json:"content"`
}
