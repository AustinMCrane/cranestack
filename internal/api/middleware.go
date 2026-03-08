package api

import (
	"net/http"
	"strings"

	"github.com/AustinMCrane/toedoe/internal/apictx"
)

// RequireAuth returns middleware that validates a Bearer session token,
// resolves the user ID from the session store, and injects it into the context.
func RequireAuth(sessions *SessionStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			userID, ok := sessions.Lookup(token)
			if !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(apictx.WithUserID(r.Context(), userID)))
		})
	}
}
