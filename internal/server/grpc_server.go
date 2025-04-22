package server

import (
	"log"
	"net"

	"github.com/qu1etboy/go-grpc-starter/config"
	"github.com/qu1etboy/go-grpc-starter/gen/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/qu1etboy/go-grpc-starter/internal/entity"
	grpc_srv "github.com/qu1etboy/go-grpc-starter/internal/handlers/grpc"
	"github.com/qu1etboy/go-grpc-starter/internal/repository"
	userrepository "github.com/qu1etboy/go-grpc-starter/internal/repository/user_repository"
)

func StartGrpcServer(database *gorm.DB, config config.Config) {
	lis, err := net.Listen("tcp", ":8980")
	if err != nil {
		log.Fatalln("failed to listen", err)
		return
	}

	s := grpc.NewServer()

	registerGrpcService(s, database)

	log.Println("gRPC server listen at port 8980")
	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
		return
	}
}

func registerGrpcService(s *grpc.Server, database *gorm.DB) {
	repo := userrepository.UserRepository{BaseRepository: repository.BaseRepository[entity.User]{DB: database}}
	proto.RegisterUserServiceServer(s, &grpc_srv.UserGrpcService{
		UserRepository: repo,
	})
}
