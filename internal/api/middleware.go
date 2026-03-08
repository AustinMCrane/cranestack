package api

import (
	"net/http"

	"github.com/AustinMCrane/cranekit/auth"
)

// RequireAuth returns middleware that validates a Bearer session token,
// resolves the user ID from the session store, and injects it into the context.
func RequireAuth(sessions *auth.SessionStore) func(http.Handler) http.Handler {
	return auth.RequireAuth(sessions)
}
