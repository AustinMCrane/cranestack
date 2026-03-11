package api

import (
	"github.com/AustinMCrane/cranekit/auth"
	"github.com/AustinMCrane/cranekit/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes wires all HTTP routes and returns the root handler.
func (s *Server) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	kit := handlers.New(s.repo, s.repo, s.sessions, handlers.Config{
		AppleClientID: s.cfg.AppleClientID,
		PATPrefix:     s.cfg.PATPrefix,
	})

	// Public auth endpoints
	r.Post("/auth/login", kit.Login)

	// Protected endpoints — require a valid session token or PAT
	r.Group(func(r chi.Router) {
		r.Use(auth.RequireAnyAuth(s.sessions, s.repo))
		r.Post("/auth/generate-mcp-key", kit.GenerateMCPKey)
		r.Get("/api/subscriptions", kit.GetSubscription)
		r.Post("/api/subscriptions", kit.CreateSubscription)
	})

	return r
}
