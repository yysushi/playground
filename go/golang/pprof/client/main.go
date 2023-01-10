package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "net/http/pprof"
)

func requestLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		resp, err := http.Get("http://localhost:8080/hoge")
		if err != nil {
			log.Printf("request error: %s\n", err)
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
	}
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	requestLoop(ctx)
}
