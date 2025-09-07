package client

// ActionResponse represents the output for non-streaming action flows
type ActionResponse struct {
	Success  bool                   `json:"success"`
	Result   map[string]interface{} `json:"result,omitempty"`
	Error    string                 `json:"error,omitempty"`
}