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
	LogLevel      string
	LogFormat     string
	Environment   string
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
	// Only load .env file in development (when ENVIRONMENT is not set to production)
	env := os.Getenv("ENVIRONMENT")
	if env != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("Warning: .env file not found, relying on environment variables: %v", err)
		}
	} else {
		log.Printf("Production environment detected, using system environment variables")
	}

	// Retrieve DATABASE_URL from environment
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Printf("Error: DATABASE_URL is empty or invalid")
		log.Printf("Please ensure PostgreSQL database is added to your Railway project")
		log.Printf("Run: railway add --database postgresql")
		return nil, ErrEmptyDatabaseUrl
	}

	// Retrieve PORT from environment, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Warning: PORT not set, using default: 8080")
	}

	// Retrieve log level from environment, default to info
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	// Retrieve log format from environment, default to json
	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "" {
		logFormat = "json"
	}

	// Get environment
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	// Populate and return the config
	config := &Config{
		DatabaseUrl:   url,
		ServerAddress: port,
		LogLevel:      logLevel,
		LogFormat:     logFormat,
		Environment:   environment,
	}

	log.Printf("Configuration loaded successfully:")
	log.Printf("- Database URL: %s", maskDatabaseUrl(url))
	log.Printf("- Server Address: %s", port)
	log.Printf("- Environment: %s", environment)
	log.Printf("- Log Level: %s", logLevel)
	log.Printf("- Log Format: %s", logFormat)

	return config, nil
}

// maskDatabaseUrl masks sensitive parts of the database URL for logging
func maskDatabaseUrl(url string) string {
	if len(url) < 20 {
		return "***"
	}
	return url[:10] + "***" + url[len(url)-10:]
}
