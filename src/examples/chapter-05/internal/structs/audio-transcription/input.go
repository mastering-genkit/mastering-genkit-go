package audio

// AudioRequest represents the input for audio transcription and analysis
type AudioRequest struct {
	AudioURL string `json:"audio_url"` // URL of the audio file
	Topic    string `json:"topic"`     // Meeting topic/context
	Language string `json:"language"`  // Audio language
}