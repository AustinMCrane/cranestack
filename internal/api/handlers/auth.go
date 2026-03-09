package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/AustinMCrane/cranekit/auth"
	"github.com/AustinMCrane/cranestack/internal/db"
)

// Login handles POST /auth/login.
// It accepts an Apple identity token, upserts the user, creates a session,
// and returns a session token.
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		IdentityToken string `json:"identity_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.IdentityToken == "" {
		http.Error(w, "identity_token required", http.StatusBadRequest)
		return
	}

	claims, err := auth.ParseAppleIDToken(r.Context(), req.IdentityToken, os.Getenv("APPLE_CLIENT_ID"))
	if err != nil {
		http.Error(w, "invalid identity token", http.StatusBadRequest)
		return
	}

	userID := "apple_" + claims.Sub
	if err := h.repo.CreateUser(&db.User{ID: userID, Email: claims.Email}); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	sessionToken, err := auth.GeneratePAT("session_")
	if err != nil {
		http.Error(w, "failed to generate session", http.StatusInternalServerError)
		return
	}
	if err := h.sessions.Register(sessionToken, userID, time.Now().Add(24*time.Hour)); err != nil {
		http.Error(w, "failed to store session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": sessionToken})
}

// GenerateMCPKey handles POST /auth/generate-mcp-key.
// It generates a PAT, stores a SHA-256 hash of it, and returns the raw token once.
func (h *Handlers) GenerateMCPKey(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	rawToken, err := h.repo.GenerateAPIKey(userID, "MCP Token", "cranestack_")
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": rawToken})
}
