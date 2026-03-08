package db

import "database/sql"

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

// migrate runs CREATE TABLE statements to initialize the schema.
func migrate(db *sql.DB) error {
	for _, stmt := range []string{createUsers, createAPIKeys} {
		if _, err := db.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}
