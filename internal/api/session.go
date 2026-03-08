package api

import (
	"github.com/AustinMCrane/cranekit/auth"
)

// SessionStore is an alias for cranekit's SessionStore.
type SessionStore = auth.SessionStore

// NewSessionStore creates an empty SessionStore.
func NewSessionStore() *SessionStore {
	return auth.NewSessionStore()
}
