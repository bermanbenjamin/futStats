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
	
	// For debugging, let's make CORS more permissive
	corsConfig.AllowAllOrigins = true
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
	corsConfig.MaxAge = 12 * 60 * 60 // 12 hours

	log.Printf("CORS Configuration (Debug Mode - Allow All Origins):")
	log.Printf("- Allow All Origins: %v", corsConfig.AllowAllOrigins)
	log.Printf("- Allowed Methods: %v", corsConfig.AllowMethods)
	log.Printf("- Allowed Headers: %v", corsConfig.AllowHeaders)

	g.Use(cors.New(corsConfig))

	routers.SetupRouter(g, dep)

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := g.Run(":" + cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
