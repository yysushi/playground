package main

import (
	"context"
	"crypto/rand"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	slogmulti "github.com/samber/slog-multi"
)

type client chan string

var (
	clients   = make(map[client]bool)
	clientsMu sync.Mutex
	logCh     = make(chan string)
)

//go:embed index.html
var indexHTML []byte

type sseBroadcastHandler struct {
	*slog.JSONHandler
}

func (h *sseBroadcastHandler) Handle(ctx context.Context, r slog.Record) error {
	// Construct a map for JSON
	entry := make(map[string]any)
	entry["time"] = r.Time.Format(time.RFC3339)
	entry["level"] = r.Level.String()
	entry["msg"] = r.Message

	// Add structured attributes
	r.Attrs(func(a slog.Attr) bool {
		entry[a.Key] = a.Value.String()
		return true
	})

	// Marshal to JSON and send to logCh
	if b, err := json.Marshal(entry); err == nil {
		logCh <- string(b)
	}

	return nil
}

func main() {
	// Use our multiHandler as the global default
	logger := slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(io.MultiWriter(os.Stderr), &slog.HandlerOptions{
				AddSource: true,
			}),
			&sseBroadcastHandler{JSONHandler: slog.NewJSONHandler(io.Discard, nil)},
		),
	)
	slog.SetDefault(logger)

	// Example structured log
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			slog.Info("periodic log", "random", rand.Text())
		}
	}()

	// SSE endpoint
	http.HandleFunc("/logs/stream", sseHandler)

	// Serve embedded index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(indexHTML)
	})

	// Broadcast loop
	go func() {
		for msg := range logCh {
			clientsMu.Lock()
			for c := range clients {
				select {
				case c <- msg:
				default:
				}
			}
			clientsMu.Unlock()
		}
	}()

	log.Println("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	c := make(client)
	clientsMu.Lock()
	clients[c] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, c)
		clientsMu.Unlock()
		close(c)
	}()

	for {
		select {
		case msg := <-c:
			_, _ = fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
