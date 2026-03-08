package main

import (
	"database/sql"
	"log/slog"
	"os"
	"time"

	"github.com/AustinMCrane/cranekit/server"
	"github.com/AustinMCrane/cranestack/internal/api"
	"github.com/AustinMCrane/cranestack/internal/db"
)

func main() {
	slog.SetDefault(server.NewLogger(os.Getenv("LOG_FORMAT")))

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "cranestack.db"
	}

	sqlDB, err := db.Open(dbPath, nil)
	if err != nil {
		slog.Error("open db", "err", err)
		os.Exit(1)
	}
	defer func(sqlDB *sql.DB) {
		if err := sqlDB.Close(); err != nil {
			slog.Error("close db", "err", err)
		}
	}(sqlDB)

	repo := db.NewRepository(sqlDB)
	sessions := api.NewSessionStore()

	if devToken := os.Getenv("DEV_SESSION_TOKEN"); devToken != "" {
		const devUserID = "dev-user"
		if err := repo.CreateUser(&db.User{ID: devUserID, Email: "dev@localhost"}); err != nil {
			slog.Warn("create dev user", "err", err)
		}
		sessions.Register(devToken, devUserID)
		slog.Warn("DEV session token registered — do not use in production", "userID", devUserID)
	}

	srv := api.NewServer(api.Config{Port: 8080}, repo, sessions)

	slog.Info("API server listening", "port", 8080)
	if err := server.RunWithGracefulShutdown(srv.HTTPServer(), 10*time.Second); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
