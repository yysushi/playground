package serer_example

import (
	"testing"

	"fmt"
	"net/http"
)

// func TestMain(m *testing.M) {
// 	m.Run()
// }

func TestServer1(t *testing.T) {
	http.HandleFunc("/handler1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler1\n")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

type MyHandler struct {
	Name string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my name is %s\n", h.Name)
}

func TestServer2(t *testing.T) {
	var h *MyHandler = &MyHandler{
		Name: "handler2",
	}
	if err := http.ListenAndServe(":8080", h); err != nil {
		panic(err)
	}
}

func TestServer3(t *testing.T) {
	var h *MyHandler = &MyHandler{
		Name: "handler3",
	}
	http.Handle("/handler3", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func TestServer4(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/handler4", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler4\n")
	})
	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}

func TestServer5(t *testing.T) {
	h := &MyHandler{
		Name: "handler5",
	}
	m := http.NewServeMux()
	m.Handle("/handler5", h)
	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}

func TestServer6(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/handler4", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler4\n")
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestServer7(t *testing.T) {
	h := &MyHandler{
		Name: "handler7",
	}
	m := http.NewServeMux()
	m.Handle("/handler7", h)
	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

// func TestServer8(t *testing.T) {
// 	h := &MyHandler{
// 		Name: "handler8",
// 	}
// 	m := http.NewServeMux()
// 	m.Handle("/handler8", h)
// 	m2 := http.NewServeMux()
// 	m2.Handle("/handler8", m)
// 	if err := http.ListenAndServe(":8080", m2); err != nil {
// 		panic(err)
// 	}
// }
