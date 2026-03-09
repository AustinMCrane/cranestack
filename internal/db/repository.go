package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	cranedb "github.com/AustinMCrane/cranekit/db"
)

// SubscriptionSchema creates the subscriptions table.
const SubscriptionSchema = `
CREATE TABLE IF NOT EXISTS subscriptions (
    id         TEXT PRIMARY KEY,
    user_id    TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_id TEXT NOT NULL,
    status     TEXT NOT NULL DEFAULT 'active',
    created_at DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at DATETIME NOT NULL DEFAULT (datetime('now'))
);
CREATE INDEX IF NOT EXISTS idx_subscriptions_user_id ON subscriptions(user_id);
`

// Subscription represents a subscriptions record.
type Subscription struct {
	ID        string
	UserID    string
	ProductID string
	Status    string
	CreatedAt string
	UpdatedAt string
}

// Repository extends cranekit's auth repository with app-specific methods.
// User, APIKey, CreateUser, GetUserByID, StoreAPIKeyHash, ValidateAPIKey,
// and GenerateAPIKey are all inherited from cranedb.Repository.
type Repository struct {
	*cranedb.Repository
	db *sql.DB
}

// Re-export cranekit types so callers don't need to import cranekit/db directly.
type User = cranedb.User
type APIKey = cranedb.APIKey

// NewRepository wraps an open *sql.DB.
func NewRepository(sqlDB *sql.DB) *Repository {
	return &Repository{Repository: cranedb.NewRepository(sqlDB), db: sqlDB}
}

// GetSubscriptionByUserID returns the active subscription for a user.
// Returns nil, nil if no subscription exists.
func (r *Repository) GetSubscriptionByUserID(userID string) (*Subscription, error) {
	var s Subscription
	err := r.db.QueryRow(
		`SELECT id, user_id, product_id, status, created_at, updated_at FROM subscriptions WHERE user_id = ? ORDER BY created_at DESC LIMIT 1`,
		userID,
	).Scan(&s.ID, &s.UserID, &s.ProductID, &s.Status, &s.CreatedAt, &s.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// CreateSubscription inserts a new subscription record.
func (r *Repository) CreateSubscription(sub *Subscription) error {
	if sub.ID == "" {
		sub.ID = fmt.Sprintf("sub_%d", time.Now().UnixNano())
	}
	_, err := r.db.Exec(
		`INSERT INTO subscriptions (id, user_id, product_id, status) VALUES (?, ?, ?, ?)`,
		sub.ID, sub.UserID, sub.ProductID, sub.Status,
	)
	return err
}
