package main

import (
	"net"

	"github.com/guilhermealvess/oauth-api/internal/interface/grpc"
	"github.com/guilhermealvess/oauth-api/internal/interface/rest"
)

func main() {
	app := rest.NewApp()
	grpcServer := grpc.NewServer()

	go func() {
		app.Start(":3000")
	}()

	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
