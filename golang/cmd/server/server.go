package main

import (
	"log"
	"net"

	"github.com/marlonklc/poc-grpc/gopb"
	"github.com/marlonklc/poc-grpc/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	grpcServer := grpc.NewServer()
	gopb.RegisterMathServiceServer(grpcServer, &servers.Math{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Cannot start grpc server: %y", err)
	}

	grpcServer.Serve(listener)
}
