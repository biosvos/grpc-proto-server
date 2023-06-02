package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
)

type Gateway struct {
	mux    *runtime.ServeMux
	server *http.Server
}

func NewGateway(address string) *Gateway {
	mux := runtime.NewServeMux()
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	return &Gateway{
		mux:    mux,
		server: server,
	}
}

func (g *Gateway) Register(ctx context.Context, fn func(ctx context.Context, mux *runtime.ServeMux) error) error {
	return fn(ctx, g.mux)
}

func (g *Gateway) Run() error {
	return g.server.ListenAndServe()
}
