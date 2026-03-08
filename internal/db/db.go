package db

import (
	"database/sql"
	"fmt"

	cranedb "github.com/AustinMCrane/cranekit/db"
)

const createUsers = `
CREATE TABLE IF NOT EXISTS users (
    id         TEXT PRIMARY KEY,
    email      TEXT NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT (datetime('now'))
);`

const createAPIKeys = `
CREATE TABLE IF NOT EXISTS api_keys (
    id         TEXT PRIMARY KEY,
    user_id    TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    key_hash   TEXT NOT NULL UNIQUE,
    label      TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT (datetime('now'))
);`

// Open initializes the ToeDoe SQLite database and runs schema migrations.
func Open(path string) (*sql.DB, error) {
	conn, err := cranedb.Open(path, []string{createUsers, createAPIKeys})
	if err != nil {
		return nil, fmt.Errorf("toedoe db: %w", err)
	}
	return conn, nil
}
