package userrepository

import (
	"github.com/qu1etboy/go-grpc-starter/internal/entity"
	"github.com/qu1etboy/go-grpc-starter/internal/repository"
)

type UserRepository struct {
	repository.BaseRepository[entity.User]
}
