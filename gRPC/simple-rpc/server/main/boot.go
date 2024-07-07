package main

import (
	"google.golang.org/grpc"
	"net"
	pb "private/Golang-/gRPC/simple-rpc/pb"
	"private/Golang-/gRPC/simple-rpc/server/service"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	pb.RegisterUserServiceServer(server, &service.UserServiceServer{})

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
