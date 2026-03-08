package db

import "database/sql"

// migrate is no longer needed — schema SQL is now passed directly to cranedb.Open
// in db.go. This file is kept for any future migration helpers.
var _ *sql.DB // suppress unused import
