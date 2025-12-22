package main

import (
	"bytes"
	"context"
	"crypto/rand"
	_ "embed"
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

type client chan []byte

var (
	clients   = make(map[client]bool)
	clientsMu sync.Mutex
	logCh     = make(chan []byte, 3)
)

//go:embed index.html
var indexHTML []byte

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

func newSSEBroadCaster() *sseBroadcastHandler {
	// SSE endpoint
	http.HandleFunc("/logs/stream", sseHandler)

	// Serve embedded index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(indexHTML)
	})

	return &sseBroadcastHandler{JSONHandler: slog.NewJSONHandler(io.Discard, nil)}
}

type sseBroadcastHandler struct {
	*slog.JSONHandler
}

func (h *sseBroadcastHandler) Handle(ctx context.Context, r slog.Record) error {
	// slog.JSONHanndler, an inner handler, won't provide an easy way of marshaling record
	// so, prepare additional json handler internally and get the value..
	// maybe, performance issue..
	buf := new(bytes.Buffer)
	jsonHandler := slog.NewJSONHandler(buf, nil)
	if err := jsonHandler.Handle(ctx, r); err != nil {
		return err
	}
	logCh <- buf.Bytes()
	return nil
}

func (h *sseBroadcastHandler) ListenAndServe(addr string, handler http.Handler) error {
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
	return http.ListenAndServe(addr, handler)
}

// func (h *sseBroadcastHandler) Wait()   {}

func main() {
	sse := newSSEBroadCaster()

	// Use our multiHandler as the global default
	logger := slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(io.MultiWriter(os.Stderr), &slog.HandlerOptions{
				AddSource: true,
			}),
			sse,
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

	log.Println("Listening on :8081")
	log.Fatal(sse.ListenAndServe(":8080", nil))
}
