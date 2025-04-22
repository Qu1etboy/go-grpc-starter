package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST string `env:"DB_HOST"`
}

// Load config from .env
// If not provide or has any error
// will return default config
func LoadConfig() Config {
	config := Config{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return config
	}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("Error reading the environment variables: %v", err)
		return config
	}

	log.Println("config", config)
	return config
}
