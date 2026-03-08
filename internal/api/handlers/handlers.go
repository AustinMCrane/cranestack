package handlers

import "github.com/AustinMCrane/cranestack/internal/db"

// Handlers holds shared dependencies for all HTTP handlers.
type Handlers struct {
	repo *db.Repository
}

// New creates a Handlers instance.
func New(repo *db.Repository) *Handlers {
	return &Handlers{repo: repo}
}
