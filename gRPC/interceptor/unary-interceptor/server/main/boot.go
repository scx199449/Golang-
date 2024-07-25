package main

import (
	"google.golang.org/grpc"
	"net"
	pb "private/Golang-/gRPC/interceptor/unary-interceptor/pb"
	"private/Golang-/gRPC/interceptor/unary-interceptor/server/interceptor"
	"private/Golang-/gRPC/interceptor/unary-interceptor/server/service"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.AuthInterceptor))
	defer server.Stop()

	pb.RegisterUserServiceServer(server, &service.UserServer{})

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
