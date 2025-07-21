package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/avast/retry-go/v4"
)

type DiscordMessage struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

func send(msg DiscordMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	slog.Debug("message details", "message", msg)
	req, err := http.NewRequest("POST", DISCORD_WEBHOOK_URL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// retry logic
	return retry.Do(func() error {
		res, err := client.Do(req)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode > 399 {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}

			slog.Debug("discord response", "body", string(body))

			return fmt.Errorf("unexpected message response: %d - %s", res.StatusCode, res.Status)
		} else {
			slog.Info("message sent", "status", res.Status, "code", res.StatusCode)
		}

		return nil
	}, retry.UntilSucceeded(), retry.Delay(10*time.Second))
}
