package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

type HogeHandler struct {
	Name string
}

func (h *HogeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "my name is %s\n", h.Name)
}

func main() {
	var h *HogeHandler = &HogeHandler{
		Name: "hoge",
	}
	http.Handle("/hoge", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("listen and serve error: %s\n", err)
	}
}
