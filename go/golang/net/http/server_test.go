package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"

	"testing"
)

// type Handler interface {
// 	ServeHTTP(http.ResponseWriter, *http.Request)
// }

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintln(w, "hello1")
}

func TestHTTPServer(t *testing.T) {
	mux := http.NewServeMux()
	// 1. register handler to default server mux
	mux.Handle("/count1", new(countHandler))
	// 2. register handler func to default server mux
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "hello2")
	}
	mux.HandleFunc("/count2", h2)
	// server
	s := http.Server{
		// Addr:    ":8080",
		Handler: mux,
	}
	start := make(chan error)
	done := make(chan error)
	go func() {
		l, err := net.Listen("tcp", ":8080")
		start <- err
		if err != nil {
			return
		}
		err = s.Serve(l)
		done <- err
	}()
	// client
	err := <-start
	if err != nil {
		t.Fatalf("no error but %s", err)
	}
	resp, err := http.Get("http://localhost:8080/count1")
	if err != nil {
		t.Fatalf("no error but %s", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("no error but %s", err)
	}
	if strings.Compare(string(b), "hello1\n") != 0 {
		t.Fatalf("expected hello1\\n, but %s", string(b))
	}
	if err := s.Shutdown(context.Background()); err != nil {
		t.Fatalf("no error but %s", err)
	}
	if err := <-done; err != http.ErrServerClosed {
		t.Fatalf("no error but %s", err)
	}
}
