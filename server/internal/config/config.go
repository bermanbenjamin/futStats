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

// Define custom error types for better debugging
var (
	ErrMissingEnvFile    = errors.New("failed to load the .env file")
	ErrEmptyDatabaseUrl  = errors.New("DATABASE_URL is empty or invalid")
	ErrMissingPort       = errors.New("PORT not set in environment variables")
	ErrInvalidPortFormat = errors.New("PORT must be a valid number")
)

// LoadConfig loads environment variables and returns a populated Config struct
func LoadConfig() (*Config, error) {
	// Attempt to load the .env file, but do not fail if it doesn't exist
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: .env file not found, relying on environment variables: %v", err)
	}

	// Retrieve DATABASE_URL from environment
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Printf("Error: DATABASE_URL is empty or invalid")
		return nil, ErrEmptyDatabaseUrl
	}

	// Retrieve PORT from environment
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("Error: PORT not set in environment variables")
		return nil, ErrMissingPort
	}

	// Populate and return the config
	config := &Config{
		DatabaseUrl:   url,
		ServerAddress: port, // Prepend ":" to match the Go server address format
	}

	log.Printf("Configuration loaded successfully: %+v", config)
	return config, nil
}
