package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type client chan string

var (
	clients   = make(map[client]bool)
	clientsMu sync.Mutex
	logCh     = make(chan string)
)

func main() {
	http.HandleFunc("/logs/stream", sseHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Broadcast loop
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

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			logCh <- rand.Text()
		}
	}()

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Register client
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

	// Subscribe until disconnect
	for {
		select {
		case msg := <-c:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
