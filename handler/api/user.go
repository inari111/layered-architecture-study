package api

import (
	"context"

	"github.com/inari111/layered-architecture-study/application/user"

	"github.com/inari111/layered-architecture-study/handler/rpc"
)

func NewUserService(app user.Application) pb.UserService {
	return &userService{
		app: app,
	}
}

type userService struct {
	app user.Application
}

func (s *userService) Register(ctx context.Context, req *pb.UserServiceRegisterRequest) (*pb.UserServiceRegisterResponse, error) {
	if err := s.app.Register(ctx); err != nil {
		return nil, err
	}
	return &pb.UserServiceRegisterResponse{}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, req *pb.UserServiceUpdateProfileRequest) (*pb.UserServiceUpdateProfileResponse, error) {
	if err := s.app.UpdateProfile(
		ctx,
		req.GetName(),
		int(req.GetAge()),
		req.GetDescription(),
	); err != nil {
		return nil, err
	}
	return &pb.UserServiceUpdateProfileResponse{}, nil
}
