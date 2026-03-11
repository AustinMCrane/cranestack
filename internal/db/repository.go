package db

import (
	"database/sql"

	cranedb "github.com/AustinMCrane/cranekit/db"
)

// SubscriptionSchema re-exports cranekit's subscription schema for use in
// db.Open migrations.
const SubscriptionSchema = cranedb.SubscriptionSchema

// Repository extends cranekit's base repository with app-specific methods.
// User, APIKey, Subscription, CreateUser, GetUserByID, StoreAPIKeyHash,
// ValidateAPIKey, GenerateAPIKey, GetSubscriptionByUserID, and
// CreateSubscription are all inherited from cranedb.Repository.
type Repository struct {
	*cranedb.Repository
	db *sql.DB
}

// Re-export cranekit types so callers don't need to import cranekit/db directly.
type User = cranedb.User
type APIKey = cranedb.APIKey
type Subscription = cranedb.Subscription

// NewRepository wraps an open *sql.DB.
func NewRepository(sqlDB *sql.DB) *Repository {
	return &Repository{Repository: cranedb.NewRepository(sqlDB), db: sqlDB}
}
