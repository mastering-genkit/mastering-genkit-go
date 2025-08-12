package image

// ImageAnalysisRequest represents the input for image analysis
type ImageAnalysisRequest struct {
	ImageURL string `json:"image_url"` // URL of the image to analyze
	Language string `json:"language"` // Output language preference
}