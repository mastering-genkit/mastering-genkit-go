package client

// EvaluateResponse represents the output for evaluation flow
type EvaluateResponse struct {
	Success         bool   `json:"success"`
	Score           int    `json:"score"`
	Feedback        string `json:"feedback"`
	CreativityScore int    `json:"creativityScore"`
	TechniqueScore  int    `json:"techniqueScore"`
	AppealScore     int    `json:"appealScore"`
	Title           string `json:"title"`       // Chef title based on score
	Achievement     string `json:"achievement"` // Special achievement description
	Error           string `json:"error"`
}
