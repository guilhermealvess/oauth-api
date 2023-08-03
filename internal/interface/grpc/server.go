package grpc

import (
	"github.com/guilhermealvess/oauth-api/internal/interface/grpc/handler"
	"github.com/guilhermealvess/oauth-api/internal/interface/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer() *grpc.Server {
	server := grpc.NewServer()
	registerServices(server)

	reflection.Register(server)

	return server
}

func registerServices(server *grpc.Server) {
	pb.RegisterApiServerServer(server, handler.NewApiServerHandler())
}
