package service

import (
	"fmt"
	pb "private/Golang-/gRPC/service-stream/pb"
	"strconv"
	"time"
)

type UserServer struct {
}

func (u *UserServer) GetUser(req *pb.GetUserRequest, stream pb.User_GetUserServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&pb.GetUserResponse{
			Name:   fmt.Sprintf(`%s-%s`, req.Search, strconv.Itoa(i+1)),
			UserId: strconv.Itoa(i + 1),
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}

	return nil
}
