package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	pb "private/Golang-/gRPC/service-stream/pb"
)

func main() {
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer userConn.Close()

	userService := pb.NewUserClient(userConn)
	userStream, err := userService.GetUser(context.Background(), &pb.GetUserRequest{Search: "user"})
	for {
		user, err := userStream.Recv()
		if err == io.EOF {
			break
		}
		// 处理可能出现的错误
		fmt.Println("Search Result : ", user)
	}

	if err != nil {
		panic(err)
	}
}
