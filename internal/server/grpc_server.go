package server

import (
	"log"
	"net"

	"github.com/qu1etboy/go-grpc-starter/gen/proto"
	"google.golang.org/grpc"

	grpc_srv "github.com/qu1etboy/go-grpc-starter/internal/handlers/grpc"
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":8980")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	registerGrpcService(s)

	log.Println("gRPC server listen at port 8980")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerGrpcService(s *grpc.Server) {
	repo := grpc_srv.UserRepository{}
	proto.RegisterUserServiceServer(s, &grpc_srv.UserGrpcService{
		UserRepository: repo,
	})
}
