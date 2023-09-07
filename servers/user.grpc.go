package servers

import (
	"context"

	"github.com/Kazai0/grpc-client/pb"
)

type Users struct {
	pb.UnimplementedUserServiceServer
}

//func to create user
func (u *Users) CreateUser(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {

	userResponse := &pb.User{}
	userResponse.Id = req.User.Id
	userResponse.Name = req.User.Name
	userResponse.Email = req.User.Email

	return &pb.UserResponse{
		User: userResponse,
	}, nil
}

func (u *Users) mustEmbedUnimplementedUserServiceServer() {}
