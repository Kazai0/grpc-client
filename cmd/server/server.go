package main

import (
	"log"
	"net"

	"github.com/Kazai0/grpc-client/pb"
	"github.com/Kazai0/grpc-client/servers"
	"google.golang.org/grpc"
)

//create server gRPC
func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &servers.Users{})

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("canot start grpc %v", err)
	}

	grpcServer.Serve(listener)
}
