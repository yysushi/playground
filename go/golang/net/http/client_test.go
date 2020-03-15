package main

import (
	"io"
	"io/ioutil"
	"net/http"

	"net/http/httptest"
	"testing"
)

func TestHTTPClient(t *testing.T) {
	// setup
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintln(w, "Hello, client")
		io.WriteString(w, "Hello, client")
	}))
	defer ts.Close()

	t.Run("simple", func(t *testing.T) {
		// send request
		res, err := http.Get(ts.URL)
		if err != nil {
			t.Fatalf("no error but %s", err)
		}
		// analyze response
		value, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatalf("no error but %s", err)
		}
		expected := "Hello, client"
		if string(value) != expected {
			t.Fatalf("expected `%s` but `%s`", expected, value)
		}
	})

	t.Run("request", func(t *testing.T) {
		// create request
		req, err := http.NewRequest("GET", ts.URL, nil)
		if err != nil {
			t.Fatalf("no error but %s", err)
		}
		// create client
		var client http.Client
		// send request
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("no error but %s", err)
		}
		// analyze response
		value, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatalf("no error but %s", err)
		}
		expected := "Hello, client"
		if string(value) != expected {
			t.Fatalf("expected `%s` but `%s`", expected, value)
		}
	})
}
