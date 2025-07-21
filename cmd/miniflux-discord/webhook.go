package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func serve(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	slog.Info("received webhook")

	var data WebhookNewEntriesEvent
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		slog.Error("failed to decode entry", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if data.EventType != NewEntriesEventType {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	encounteredError := false

	slog.Info("new entries", "count", len(data.Entries))
	for index, entry := range data.Entries {
		slog.Debug("entry data", "index", index, "entry", entry)

		message := DiscordMessage{
			Content:  fmt.Sprintf("__**%s**__\n%s", entry.Title, entry.URL),
			Username: data.Feed.Title,
		}

		if err := send(message); err != nil {
			slog.Error("failed to send message", "error", err, "entry", entry.ID)
			encounteredError = true
		}
	}

	if encounteredError {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
