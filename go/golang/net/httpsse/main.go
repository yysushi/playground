package main

import (
	"crypto/rand"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// embed index.html
//
//go:embed index.html
var indexHTML []byte

type client chan string

var (
	clients   = make(map[client]bool)
	clientsMu sync.Mutex
	logCh     = make(chan string)
)

func main() {
	// Set up structured logging with JSONHandler
	handler := slog.NewJSONHandler(log.Writer(), nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// Example structured logging from the app
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			data := map[string]any{
				"time":  time.Now().Format(time.RFC3339),
				"level": "INFO",
				"msg":   rand.Text(),
				"meta":  "example", // arbitrary extra fields
			}
			b, _ := json.Marshal(data)
			logCh <- string(b)
		}
	}()

	// SSE endpoint
	http.HandleFunc("/logs/stream", sseHandler)

	// Serve only the embedded index.html at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(indexHTML)
	})

	// Broadcast loop: thread safe sending
	go func() {
		for logLine := range logCh {
			clientsMu.Lock()
			for c := range clients {
				select {
				case c <- logLine:
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
