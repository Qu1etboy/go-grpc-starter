package db

import (
	"fmt"

	"github.com/qu1etboy/go-grpc-starter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config config.Config) (*gorm.DB, error) {
	// https://github.com/go-gorm/postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.DB.Host, config.DB.User, config.DB.Password, config.DB.DBName, config.DB.Port, config.DB.SSLMode, config.DB.TimeZone,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	return db, err
}
