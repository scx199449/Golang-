package service

import (
	"context"
	"fmt"
	pb "private/Golang-/gRPC/simple-rpc/pb"
)

type UserServiceServer struct {
}

func (u *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		UserId: req.UserId,
		Name:   fmt.Sprintf(`user-%s`, req.UserId),
	}, nil
}
