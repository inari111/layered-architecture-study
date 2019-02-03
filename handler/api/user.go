package api

import (
	"context"

	"github.com/inari111/layered-architecture-study/handler/rpc"
)

func NewUserService() pb.UserService {
	return &userService{}
}

type userService struct{}

func (*userService) Register(context.Context, *pb.UserServiceRegisterRequest) (*pb.UserServiceRegisterResponse, error) {
	panic("implement me")
}

func (*userService) UpdateProfile(context.Context, *pb.UserServiceUpdateProfileRequest) (*pb.UserServiceUpdateProfileResponse, error) {
	panic("implement me")
}
