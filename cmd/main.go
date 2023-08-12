package main

import (
	"log"
	"net"

	"github.com/guilhermealvess/oauth-api/internal/http/grpc"
	"github.com/guilhermealvess/oauth-api/internal/http/rest"
)

func main() {
	app := rest.NewApp()
	grpcServer := grpc.NewServer()

	go app.Start(":3000")

	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
