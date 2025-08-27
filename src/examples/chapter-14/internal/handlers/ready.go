package handlers

import (
	"encoding/json"
	"net/http"
)

// ReadyHandler handles the /ready endpoint
func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status:  "ready",
		Service: "genkit-app",
	}

	json.NewEncoder(w).Encode(response)
}
