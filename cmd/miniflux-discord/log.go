package main

import (
	"log/slog"
	"os"
)

func init() {
	var verbosity slog.Level
	switch env := os.Getenv("LOG_LEVEL"); env {
	case "INFO":
		verbosity = slog.LevelInfo
	case "WARN":
		verbosity = slog.LevelWarn
	case "ERROR":
		verbosity = slog.LevelError
	case "DEBUG":
		verbosity = slog.LevelDebug
	default:
		verbosity = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: verbosity}))
	slog.SetDefault(logger)
}
