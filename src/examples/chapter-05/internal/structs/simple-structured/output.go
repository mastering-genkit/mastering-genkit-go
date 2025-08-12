package simple

// ReviewAnalysis represents the structured output of review analysis
type ReviewAnalysis struct {
	Sentiment string   `json:"sentiment"` // "positive", "negative", "neutral"
	Score     int      `json:"score"`     // 1-5
	Keywords  []string `json:"keywords"`  // Key topics/features mentioned
	Summary   string   `json:"summary"`   // Brief summary
}