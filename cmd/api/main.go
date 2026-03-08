package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/AustinMCrane/toedoe/internal/api"
	"github.com/AustinMCrane/toedoe/internal/db"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "toedoe.db"
	}

	sqlDB, err := db.Open(dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer func(sqlDB *sql.DB) {
		if err := sqlDB.Close(); err != nil {
			log.Printf("close db: %v", err)
		}
	}(sqlDB)

	repo := db.NewRepository(sqlDB)
	sessions := api.NewSessionStore()

	// Register a dev session token for local development.
	// Start the server with DEV_SESSION_TOKEN=<token> to enable.
	if devToken := os.Getenv("DEV_SESSION_TOKEN"); devToken != "" {
		const devUserID = "dev-user"
		if err := repo.CreateUser(&db.User{ID: devUserID, Email: "dev@localhost"}); err != nil {
			log.Printf("warn: create dev user: %v", err)
		}
		sessions.Register(devToken, devUserID)
		log.Printf("DEV: registered session token for user %q", devUserID)
	}

	srv := api.NewServer(api.Config{Port: 8080}, repo, sessions)

	log.Println("API server listening on :8080")
	if err := srv.Start(); err != nil {
		log.Fatalf("server: %v", err)
	}
}
