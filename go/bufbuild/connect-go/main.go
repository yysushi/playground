package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	petstorev1 "github.com/yysushi/playground/go/bufbuild/connect-go/internal/gen/petstore/v1"
	"github.com/yysushi/playground/go/bufbuild/connect-go/internal/gen/petstore/v1/petstorev1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type petstoreServer struct {
}

func NewPetstoreServer() petstorev1connect.PetServiceHandler {
	return &petstoreServer{}
}

func (e *petstoreServer) Create(
	ctx context.Context,
	req *connect.Request[petstorev1.CreateRequest],
) (*connect.Response[petstorev1.CreateResponse], error) {
	return connect.NewResponse(&petstorev1.CreateResponse{
		NickName: req.Msg.Name[:1],
	}), nil
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(petstorev1connect.NewPetServiceHandler(
		NewPetstoreServer(),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(petstorev1connect.PetServiceName),
	))
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(petstorev1connect.PetServiceName),
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(petstorev1connect.PetServiceName),
	))
	srv := &http.Server{
		Addr: "localhost:8080",
		Handler: h2c.NewHandler(
			mux,
			&http2.Server{},
		),
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err) //nolint:gocritic
	}
}
