package structs

// RecommendationInput represents the client request input for the recommendation flow.
// This is the data structure that clients send to the /recommendationFlow endpoint.
//
// Example JSON request:
//
//	{
//	  "genre": "jazz",
//	  "experience": "beginner"
//	}
type RecommendationInput struct {
	Genre      string `json:"genre"`      // Music genre (e.g., "jazz", "rock", "edm")
	Experience string `json:"experience"` // Experience level: "beginner" or "intermediate" (defaults to "beginner")
}
