package handler

import (
	"context"
	pb "grpc-proto-server/proto"
)

type Handler struct {
	pb.UnimplementedGreeterServer
}

func (h *Handler) Hello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Greeting: request.GetName(),
	}, nil
}
