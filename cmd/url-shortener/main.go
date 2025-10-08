package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Print(cfg)
	log := setupLogger(cfg.Env)
	log.Debug("Hello")
}

func setupLogger(env string) slog.Logger {
	var log *slog.Logger
	var handler slog.Handler
	switch env {
	case envLocal:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	log = slog.New(handler)
	return *log
}
