package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

var (
	LISTEN_PORT         int
	LISTEN_ADDR         string
	DISCORD_WEBHOOK_URL string
)

func parseEnv() error {
	envPort := os.Getenv("LISTEN_PORT")
	if envPort == "" {
		LISTEN_PORT = 8080
	} else {
		var err error
		LISTEN_PORT, err = strconv.Atoi(envPort)
		if err != nil {
			return err
		}
	}

	envAddr := os.Getenv("LISTEN_ADDR")
	if envAddr == "" {
		LISTEN_ADDR = "0.0.0.0"
	} else {
		// TODO: validate
		LISTEN_ADDR = envAddr
	}

	DISCORD_WEBHOOK_URL = os.Getenv("DISCORD_WEBHOOK_URL")
	if DISCORD_WEBHOOK_URL == "" {
		return fmt.Errorf("missing DISCORD_WEBHOOK_URL")
	}

	return nil
}

func main() {
	if err := parseEnv(); err != nil {
		slog.Error("configuration issue", "error", err)
		os.Exit(1) // really sucks slog.Fatal doesn't exist
	}

	http.HandleFunc("/", serve)

	slog.Info("webhook server started", "addr", LISTEN_ADDR, "port", LISTEN_PORT)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", LISTEN_ADDR, LISTEN_PORT), nil); err != nil {
		slog.Error("webhook server error", "error", err)
		os.Exit(1)
	}
}
