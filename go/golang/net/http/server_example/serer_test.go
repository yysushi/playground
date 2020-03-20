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

// not good example
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
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler4\n")
	})
	http.Handle("/handler4", h)
	if err := http.ListenAndServe(":8080", h); err != nil {
		panic(err)
	}
}

func TestServer5(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/handler5", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler5\n")
	})
	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}

func TestServer6(t *testing.T) {
	h := &MyHandler{
		Name: "handler6",
	}
	m := http.NewServeMux()
	m.Handle("/handler6", h)
	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}

func TestServer7(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler7\n")
	})
	m := http.NewServeMux()
	m.Handle("/handler7", h)
	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}

func TestServer8(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/handler8", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler8\n")
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestServer9(t *testing.T) {
	h := &MyHandler{
		Name: "handler9",
	}
	m := http.NewServeMux()
	m.Handle("/handler9", h)
	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestServer10(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is handler10\n")
	})
	m := http.NewServeMux()
	m.Handle("/handler10", h)
	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

// func TestServer11(t *testing.T) {
// 	h := &MyHandler{
// 		Name: "handler11",
// 	}
// 	m := http.NewServeMux()
// 	m.Handle("/handler11", h)
// 	m2 := http.NewServeMux()
// 	m2.Handle("/handler11", m)
// 	if err := http.ListenAndServe(":8080", m2); err != nil {
// 		panic(err)
// 	}
// }
