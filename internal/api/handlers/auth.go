package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AustinMCrane/cranekit/auth"
	"github.com/AustinMCrane/cranestack/internal/db"
)

// Login handles POST /auth/login.
// It accepts an Apple identity token and exchanges it for a session.
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: parse Apple identity token from request body
	// TODO: verify token with Apple's public keys
	// TODO: upsert user in DB, create session, return session token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "not implemented"})
}

// GenerateMCPKey handles POST /auth/generate-mcp-key.
// It generates a PAT, stores a SHA-256 hash of it, and returns the raw token once.
func (h *Handlers) GenerateMCPKey(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	rawToken, err := generatePAT()
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	hash := sha256.Sum256([]byte(rawToken))
	keyHash := hex.EncodeToString(hash[:])

	id, err := newID()
	if err != nil {
		http.Error(w, "failed to generate id", http.StatusInternalServerError)
		return
	}

	if err := h.repo.StoreAPIKeyHash(&db.APIKey{
		ID:      id,
		UserID:  userID,
		KeyHash: keyHash,
		Label:   "MCP Token",
	}); err != nil {
		http.Error(w, "failed to store token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": rawToken})
}

// generatePAT creates a cryptographically random PAT with a "cranestack_" prefix.
func generatePAT() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("rand: %w", err)
	}
	return "cranestack_" + hex.EncodeToString(b), nil
}

// newID generates a random hex ID.
func newID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("rand: %w", err)
	}
	return hex.EncodeToString(b), nil
}
