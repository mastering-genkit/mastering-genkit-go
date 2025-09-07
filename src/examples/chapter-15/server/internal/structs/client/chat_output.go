package client

// ChatResponse represents a streaming response chunk for cookingBattleChat flow
type ChatResponse struct {
	Type     string    `json:"type"` // "content", "tool_call", "done", "error"
	Content  string    `json:"content,omitempty"`
	ToolCall *ToolCall `json:"toolCall,omitempty"`
	Error    string    `json:"error,omitempty"`
}

// ToolCall represents a tool invocation in the streaming response
type ToolCall struct {
	Name   string                 `json:"name"`
	Args   map[string]interface{} `json:"args"`
	Result interface{}            `json:"result,omitempty"`
}