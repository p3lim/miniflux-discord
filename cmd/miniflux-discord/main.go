package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	LISTEN_PORT         int
	LISTEN_ADDR         string
	DISCORD_WEBHOOK_URL string
)

func parseEnv() error {
	if envPort := os.Getenv("LISTEN_PORT"); envPort != "" {
		var err error
		LISTEN_PORT, err = strconv.Atoi(envPort)
		if err != nil {
			return err
		}
	} else {
		LISTEN_PORT = 8080
	}

	if envAddr := os.Getenv("LISTEN_ADDR"); envAddr != "" {
		// TODO: validate
		LISTEN_ADDR = envAddr
	} else {
		LISTEN_ADDR = "0.0.0.0"
	}

	if discordFile := os.Getenv("DISCORD_WEBHOOK_URL_FILE"); discordFile != "" {
		fileData, err := os.ReadFile(discordFile)
		if err != nil {
			return err
		}

		DISCORD_WEBHOOK_URL = strings.TrimSpace(string(fileData))
	} else {
		DISCORD_WEBHOOK_URL = os.Getenv("DISCORD_WEBHOOK_URL")
	}

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
