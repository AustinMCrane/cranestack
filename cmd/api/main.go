package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AustinMCrane/toedoe/internal/api"
	"github.com/AustinMCrane/toedoe/internal/db"
)

func main() {
	logger := newLogger()
	slog.SetDefault(logger)

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "toedoe.db"
	}

	sqlDB, err := db.Open(dbPath)
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

	go func() {
		slog.Info("API server listening", "port", 8080)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("forced shutdown", "err", err)
		os.Exit(1)
	}
	slog.Info("server stopped")
}

func newLogger() *slog.Logger {
	var handler slog.Handler
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	if os.Getenv("LOG_FORMAT") == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}
	return slog.New(handler)
}
