package main

import (
	"context"
	"google.golang.org/grpc"
	pb "private/Golang-/gRPC/interceptor/unary-interceptor/pb"
)

func main() {
	userConn, _ := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer userConn.Close()

	userService := pb.NewUserServiceClient(userConn)
	user, _ := userService.GetUser(context.Background(), &pb.GetUserRequest{UserId: 1})
	println(user.Name)
}
