package db

import (
	"database/sql"

	cranedb "github.com/AustinMCrane/cranekit/db"
)

// Repository extends cranekit's auth repository with app-specific methods.
// User, APIKey, CreateUser, GetUserByID, StoreAPIKeyHash, ValidateAPIKey,
// and GenerateAPIKey are all inherited from cranedb.Repository.
type Repository struct {
	*cranedb.Repository
}

// Re-export cranekit types so callers don't need to import cranekit/db directly.
type User = cranedb.User
type APIKey = cranedb.APIKey

// NewRepository wraps an open *sql.DB.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{Repository: cranedb.NewRepository(db)}
}
