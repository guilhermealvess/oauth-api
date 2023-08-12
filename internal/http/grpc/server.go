package grpc

import (
	"github.com/guilhermealvess/oauth-api/internal/http/grpc/handler"
	"github.com/guilhermealvess/oauth-api/pkg/grpc/pb"
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
	pb.RegisterApplicationServer(server, handler.NewApiServerHandler())
}
