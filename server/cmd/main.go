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

	// Get allowed origins from environment variables
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		// Build origins from individual environment variables
		var origins []string
		
		// Always allow localhost for development
		origins = append(origins, "http://localhost:3000")
		
		// Add Vercel frontend URL if provided
		if vercelUrl := os.Getenv("VERCEL_FRONTEND_URL"); vercelUrl != "" {
			origins = append(origins, vercelUrl)
		}
		
		// Add custom domain if provided
		if customDomain := os.Getenv("CUSTOM_DOMAIN"); customDomain != "" {
			origins = append(origins, customDomain)
		}
		
		corsConfig.AllowOrigins = origins
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
