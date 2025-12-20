package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
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

// FanoutHandler duplicates log records to two destinations:
// 1) A wrapped handler (console/file/whatever)
// 2) SSE broadcaster (via logCh)
type FanoutHandler struct {
	dest slog.Handler
}

func (h *FanoutHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.dest.Enabled(ctx, level)
}

func (h *FanoutHandler) Handle(ctx context.Context, r slog.Record) error {
	// First let the wrapped handler (e.g. JSONHandler) output normally
	err := h.dest.Handle(ctx, r)
	if err != nil {
		return err
	}

	// Prepare a structured map of the record
	obj := map[string]any{
		"time":  r.Time.Format(time.RFC3339),
		"level": r.Level.String(),
		"msg":   r.Message,
	}

	// Collect attributes
	r.Attrs(func(a slog.Attr) bool {
		obj[a.Key] = a.Value.String()
		return true
	})

	// Marshal to JSON and send to SSE channel
	if b, e := json.Marshal(obj); e == nil {
		logCh <- string(b)
	}

	return nil
}

func (h *FanoutHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &FanoutHandler{dest: h.dest.WithAttrs(attrs)}
}

func (h *FanoutHandler) WithGroup(name string) slog.Handler {
	return &FanoutHandler{dest: h.dest.WithGroup(name)}
}

func main() {
	// Use JSONHandler as the base (console logging in JSON format)
	baseHandler := slog.NewJSONHandler(log.Writer(), &slog.HandlerOptions{AddSource: true})

	// Wrap it in our custom handler
	logger := slog.New(&FanoutHandler{dest: baseHandler})
	slog.SetDefault(logger)

	// Now all slog.<Level> calls will volume out to both console and logCh

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(indexHTML)
	})

	http.HandleFunc("/logs/stream", sseHandler)

	// Example real app logs:
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			slog.Info("periodic log", "random", rand.Text())
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
