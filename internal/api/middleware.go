package api

import (
	"net/http"

	"github.com/AustinMCrane/cranekit/auth"
	"github.com/AustinMCrane/cranestack/internal/db"
)

// RequireAuth returns middleware that validates a Bearer session token only.
func RequireAuth(sessions *auth.SessionStore) func(http.Handler) http.Handler {
	return auth.RequireAuth(sessions)
}

// RequireAnyAuth returns middleware that accepts either a session token or a
// PAT (stored as a hash in the api_keys table).
func RequireAnyAuth(sessions *auth.SessionStore, repo *db.Repository) func(http.Handler) http.Handler {
	return auth.RequireAnyAuth(sessions, repo)
}
