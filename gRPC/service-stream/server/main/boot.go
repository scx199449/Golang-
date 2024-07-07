package main

import (
	"google.golang.org/grpc"
	"net"
	pb "private/Golang-/gRPC/service-stream/pb"
	"private/Golang-/gRPC/service-stream/server/service"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	pb.RegisterUserServer(server, &service.UserServer{})

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
