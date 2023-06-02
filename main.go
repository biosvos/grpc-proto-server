package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-proto-server/gateway"
	"grpc-proto-server/handler"
	pb "grpc-proto-server/proto"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	defer server.GracefulStop()

	pb.RegisterGreeterServer(server, &handler.Handler{})
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		err = server.Serve(listen)
		if err != nil {
			log.Fatal(err)
		}
		cancelFunc()
	}()

	gw := gateway.NewGateway(":9090")
	conn, err := grpc.DialContext(
		ctx,
		listen.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = gw.Register(ctx, func(ctx context.Context, mux *runtime.ServeMux) error {
		return pb.RegisterGreeterHandler(ctx, mux, conn)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = gw.Run()
	if err != nil {
		log.Fatal(err)
	}
}
