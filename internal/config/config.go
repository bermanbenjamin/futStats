package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl   string
	ServerAddress string
}

// Define custom error types
var ErrMissingEnvFile = errors.New("failed to load the .env file")
var ErrEmptyDatabaseUrl = errors.New("DATABASE_URL is empty or invalid")
var ErrMissingPort = errors.New("PORT not set in environment variables")

func LoadConfig() (cfg *Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
		return nil, ErrMissingEnvFile
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Printf("DATABASE_URL is empty or invalid")
		return nil, ErrEmptyDatabaseUrl
	}

	address := os.Getenv("PORT")
	if address == "" {
		return nil, ErrMissingPort
	}

	cfg = &Config{
		DatabaseUrl:   url,
		ServerAddress: address,
	}

	return cfg, nil
}
