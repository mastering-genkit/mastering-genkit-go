package audio

// MeetingMinutes represents the structured output of meeting transcription
type MeetingMinutes struct {
	Transcription    string           `json:"transcription"`
	SpeakerSummaries []SpeakerSummary `json:"speaker_summaries"`
	Decisions        []string         `json:"decisions"`
	ActionItems      []ActionItem     `json:"action_items"`
	Keywords         []string         `json:"keywords"`
}

// SpeakerSummary represents a summary of points made by each speaker
type SpeakerSummary struct {
	Speaker string   `json:"speaker"`
	Points  []string `json:"points"`
}

// ActionItem represents a task or action item from the meeting
type ActionItem struct {
	Task     string `json:"task"`
	Assignee string `json:"assignee"`
	DueDate  string `json:"due_date"`
	Priority string `json:"priority"` // high, medium, low
}