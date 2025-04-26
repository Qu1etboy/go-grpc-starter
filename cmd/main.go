package main

import (
	"log"

	"github.com/qu1etboy/go-grpc-starter/config"
	"github.com/qu1etboy/go-grpc-starter/internal/db"
	"github.com/qu1etboy/go-grpc-starter/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	database, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatalln("failed to connect to database", err)
		return
	}

	server.StartGrpcServer(database, cfg)
}
