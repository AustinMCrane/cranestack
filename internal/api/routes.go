package api

import (
	"github.com/AustinMCrane/toedoe/internal/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes wires all HTTP routes and returns the root handler.
func (s *Server) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := handlers.New(s.repo)

	// Public auth endpoints
	r.Post("/auth/login", h.Login)

	// Protected endpoints — require a valid session token
	r.Group(func(r chi.Router) {
		r.Use(RequireAuth(s.sessions))
		r.Post("/auth/generate-mcp-key", h.GenerateMCPKey)
		r.Get("/api/data", h.GetData)
	})

	return r
}
