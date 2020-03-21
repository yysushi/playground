package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"

	"go.uber.org/fx"
)

// Config ...
type Config struct {
	Addr string
}

// NewConfig ...
func NewConfig() *Config {
	configFileName := flag.String("config", "config.sample.json", "")
	flag.Parse()
	b, err := ioutil.ReadFile(*configFileName)
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return &config
}

// NewServeMux ...
func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return mux
}

// NewServer ...
func NewServer(mux *http.ServeMux, config *Config) *http.Server {
	return &http.Server{
		Addr:    config.Addr,
		Handler: mux,
	}
}

// Register ...
func Register(lifecycle fx.Lifecycle, server *http.Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)
}

func main() {
	fx.New(
		fx.Provide(NewConfig),
		fx.Provide(NewServer),
		fx.Provide(NewServeMux),
		fx.Invoke(Register),
	).Run()
}
