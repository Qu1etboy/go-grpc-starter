package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	DBName   string `env:"DB_NAME,required"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
	TimeZone string `env:"DB_TIMEZONE" envDefault:"UTC"`
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
