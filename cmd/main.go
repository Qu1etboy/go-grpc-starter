package main

import (
	"log"

	"github.com/qu1etboy/go-grpc-starter/config"
	"github.com/qu1etboy/go-grpc-starter/db"
	"github.com/qu1etboy/go-grpc-starter/internal/entity"
	"github.com/qu1etboy/go-grpc-starter/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	database, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatalln("failed to connect to database", err)
		return
	}

	database.AutoMigrate(&entity.User{})

	server.StartGrpcServer(database, cfg)
}
