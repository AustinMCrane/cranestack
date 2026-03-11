package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AustinMCrane/cranestack/internal/db"
)

// Config holds runtime configuration for the API server.
type Config struct {
	Port          int
	AppleClientID string
	PATPrefix     string
}

// Server is the HTTP API server.
type Server struct {
	cfg      Config
	repo     *db.Repository
	sessions *SessionStore
	http     *http.Server
}

// NewServer creates a new Server with the given config, repository, and session store.
func NewServer(cfg Config, repo *db.Repository, sessions *SessionStore) *Server {
	s := &Server{cfg: cfg, repo: repo, sessions: sessions}
	s.http = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: s.routes(),
	}
	return s
}

// Start begins listening for incoming HTTP connections.
func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

// HTTPServer returns the underlying *http.Server, e.g. to pass to
// server.RunWithGracefulShutdown.
func (s *Server) HTTPServer() *http.Server {
	return s.http
}

// Shutdown gracefully stops the HTTP server with the given context deadline.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
