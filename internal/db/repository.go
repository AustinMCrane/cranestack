package db

import "database/sql"

// Repository provides data access methods backed by SQLite.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// User represents a user record.
type User struct {
	ID        string
	Email     string
	CreatedAt string
}

// APIKey represents an api_keys record.
type APIKey struct {
	ID        string
	UserID    string
	KeyHash   string
	Label     string
	CreatedAt string
}

// GetUserByID retrieves a user by their ID.
func (r *Repository) GetUserByID(id string) (*User, error) {
	// TODO: implement
	return nil, nil
}

// CreateUser inserts a new user record, ignoring conflicts on id or email.
func (r *Repository) CreateUser(user *User) error {
	_, err := r.db.Exec(
		`INSERT OR IGNORE INTO users (id, email) VALUES (?, ?)`,
		user.ID, user.Email,
	)
	return err
}

// StoreAPIKeyHash persists a hashed API key for a user.
func (r *Repository) StoreAPIKeyHash(apiKey *APIKey) error {
	_, err := r.db.Exec(
		`INSERT INTO api_keys (id, user_id, key_hash, label) VALUES (?, ?, ?, ?)`,
		apiKey.ID, apiKey.UserID, apiKey.KeyHash, apiKey.Label,
	)
	return err
}

// ValidateAPIKey checks whether a given raw key matches a stored hash.
// Returns the associated user ID on success.
func (r *Repository) ValidateAPIKey(rawKey string) (string, error) {
	// TODO: implement
	return "", nil
}
