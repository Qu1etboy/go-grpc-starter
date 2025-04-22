package grpc

import (
	"context"

	"github.com/qu1etboy/go-grpc-starter/gen/proto"
	userrepository "github.com/qu1etboy/go-grpc-starter/internal/repository/user_repository"
)

type UserGrpcService struct {
	proto.UnimplementedUserServiceServer
	UserRepository userrepository.UserRepository
}

func (s UserGrpcService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	user, err := s.UserRepository.FindById(req.GetId())

	if err != nil {
		return nil, err
	}

	return &proto.User{
		Id:    uint64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
