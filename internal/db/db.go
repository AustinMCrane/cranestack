package db

import (
	"database/sql"
	"fmt"

	cranedb "github.com/AustinMCrane/cranekit/db"
)

// Open initializes the CraneStack SQLite database and runs app-specific
// migrations on top of the base auth schema (users + api_keys).
func Open(path string, extra []string) (*sql.DB, error) {
	conn, err := cranedb.OpenWithBaseSchemas(path, extra)
	if err != nil {
		return nil, fmt.Errorf("cranestack db: %w", err)
	}
	return conn, nil
}
