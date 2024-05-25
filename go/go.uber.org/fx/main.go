package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", 0)
	logger.Print("Executing NewLogger.")
	return logger
}

func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

type Config struct {
	Addr string
}

func NewConfig() (*Config, error) {
	configFileName := flag.String("config", "config.sample.json", "configuration file the server reads")
	flag.Parse()
	b, err := ioutil.ReadFile(*configFileName)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func NewMux(lc fx.Lifecycle, logger *log.Logger, config *Config) *http.ServeMux {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    config.Addr,
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})
	return mux
}

func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

func RegisterFailure(f *Failure) {
}

type Failure struct{}

func NewFailure(lc fx.Lifecycle) *Failure {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return errors.New("failed at start")
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return &Failure{}
}

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewHandler,
			NewConfig,
			NewMux,
			NewFailure,
		),
		fx.Invoke(
			Register,
			RegisterFailure,
		),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
	<-app.Done()
	stopCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
