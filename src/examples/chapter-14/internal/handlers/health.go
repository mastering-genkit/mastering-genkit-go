package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthHandler handles the /health endpoint
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status:  "healthy",
		Service: "genkit-app",
	}

	json.NewEncoder(w).Encode(response)
}
