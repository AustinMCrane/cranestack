package api

import "sync"

// SessionStore is a thread-safe in-memory map of session token → user ID.
type SessionStore struct {
	mu       sync.RWMutex
	sessions map[string]string
}

// NewSessionStore creates an empty SessionStore.
func NewSessionStore() *SessionStore {
	return &SessionStore{sessions: make(map[string]string)}
}

// Register adds a token → userID mapping.
func (s *SessionStore) Register(token, userID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[token] = userID
}

// Lookup returns the userID for the given token, and whether it was found.
func (s *SessionStore) Lookup(token string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	userID, ok := s.sessions[token]
	return userID, ok
}
