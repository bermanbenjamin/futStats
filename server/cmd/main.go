package main

import (
	"log"
	"os"
	"strings"

	routers "github.com/bermanbenjamin/futStats/cmd/api"
	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/config"
	"github.com/bermanbenjamin/futStats/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	err = db.InitDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	dep := config.InitializeDependencies(db.DB)

	g := gin.Default()

	// CORS configuration for production
	corsConfig := cors.DefaultConfig()

	// Get allowed origins from environment or use default
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		// Default origins for development and production
		corsConfig.AllowOrigins = []string{
			"http://localhost:3000", // Local development
			"https://client-9veaycdr8-bermanbenjamins-projects.vercel.app", // Current Vercel deployment
			"https://futstats-frontend.vercel.app",                         // Vercel deployment
			"https://futstats.vercel.app",                                  // Custom domain
		}
	} else {
		// Split comma-separated origins
		corsConfig.AllowOrigins = strings.Split(allowedOrigins, ",")
	}

	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
		"X-Requested-With",
		constants.QUERY_FILTER,
	}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}

	g.Use(cors.New(corsConfig))

	routers.SetupRouter(g, dep)

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := g.Run(":" + cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
