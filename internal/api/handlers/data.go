package handlers

import (
	"encoding/json"
	"net/http"
)

// GetData handles GET /api/data.
// This is a protected endpoint that requires a valid Bearer token.
func (h *Handlers) GetData(w http.ResponseWriter, r *http.Request) {
	// TODO: extract authenticated user from context
	// TODO: fetch and return user-specific data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "not implemented"})
}
