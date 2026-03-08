package handlers

import (
	"github.com/AustinMCrane/cranekit/auth"
	"github.com/AustinMCrane/cranestack/internal/db"
)

// Handlers holds shared dependencies for all HTTP handlers.
type Handlers struct {
	repo     *db.Repository
	sessions *auth.SessionStore
}

// New creates a Handlers instance.
func New(repo *db.Repository, sessions *auth.SessionStore) *Handlers {
	return &Handlers{repo: repo, sessions: sessions}
}
