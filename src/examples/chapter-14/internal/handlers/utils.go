package handlers

// Response represents the response structure for health checks
type Response struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}
