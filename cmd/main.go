package main

import (
	"github.com/qu1etboy/go-grpc-starter/config"
	"github.com/qu1etboy/go-grpc-starter/internal/server"
)

func main() {
	config.LoadConfig()
	server.StartGrpcServer()
}
