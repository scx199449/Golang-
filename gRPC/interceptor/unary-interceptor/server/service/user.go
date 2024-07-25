package service

import (
	"context"
	"fmt"
	pb "private/Golang-/gRPC/interceptor/unary-interceptor/pb"
)

type UserServer struct{}

func (u *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	role := ctx.Value("role")

	return &pb.GetUserResponse{
		Name: fmt.Sprintf(`user-%s-%d`, role, req.UserId),
	}, nil
}
