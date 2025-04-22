package grpc

import (
	"context"

	"github.com/qu1etboy/go-grpc-starter/gen/proto"
)

type UserRepository struct{}

type UserGrpcService struct {
	proto.UnimplementedUserServiceServer
	UserRepository UserRepository
}

func (r UserRepository) GetUserById(id uint64) (*proto.User, error) {
	return &proto.User{
		Id:    id,
		Name:  "Weerawong Vonggatunyu",
		Email: "weerawong.v@gmail.com",
	}, nil
}

func (s UserGrpcService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	return s.UserRepository.GetUserById(req.GetId())
}
