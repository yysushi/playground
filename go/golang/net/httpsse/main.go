package main

import (
	"context"
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

	"crypto/rand"
)

type client chan string

var (
	clients   = make(map[client]bool)
	clientsMu sync.Mutex
	logCh     = make(chan string)
)

//go:embed index.html
var indexHTML []byte

// multiHandler writes log records to two destinations:
// - A console output handler (JSON over os.Stderr)
// - The SSE broadcast channel
type multiHandler struct {
	consoleHandler slog.Handler
}

func newMultiHandler() slog.Handler {
	// Create a handler that writes JSON to STDERR (separate writer).
	console := slog.NewJSONHandler(io.MultiWriter(os.Stderr), &slog.HandlerOptions{
		AddSource: true,
	})
	return &multiHandler{consoleHandler: console}
}

func (h *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.consoleHandler.Enabled(ctx, level)
}

func (h *multiHandler) Handle(ctx context.Context, r slog.Record) error {
	// First write to the console handler
	if err := h.consoleHandler.Handle(ctx, r); err != nil {
		return err
	}

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

func (h *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &multiHandler{
		consoleHandler: h.consoleHandler.WithAttrs(attrs),
	}
}

func (h *multiHandler) WithGroup(name string) slog.Handler {
	return &multiHandler{
		consoleHandler: h.consoleHandler.WithGroup(name),
	}
}

func main() {
	// Use our multiHandler as the global default
	logger := slog.New(newMultiHandler())
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

	log.Println("Listening on :8080")
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
