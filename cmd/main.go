package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/andikem/kraken/pkg/routes"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", home)

	routes.RegisterRoutes(grpcServer)

	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	mixedHandler := newHTTPandGRPCMux(httpMux, grpcServer)
	http2Server := &http2.Server{}
	http1Server := &http.Server{Handler: h2c.NewHandler(mixedHandler, http2Server)}
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	err = http1Server.Serve(lis)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		panic(err)
	}
}

func newHTTPandGRPCMux(httpHand http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
			return
		}
		httpHand.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from http handler!\n")
}
