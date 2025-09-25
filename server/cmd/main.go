package main

import (
	"log"

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
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://client-j1b0qb677-bermanbenjamins-projects.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", constants.QUERY_FILTER},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	}

	log.Printf("CORS Configuration:")
	log.Printf("- Allowed Origins: %v", corsConfig.AllowOrigins)
	log.Printf("- Allowed Methods: %v", corsConfig.AllowMethods)
	log.Printf("- Allow Credentials: %v", corsConfig.AllowCredentials)

	g.Use(cors.New(corsConfig))

	routers.SetupRouter(g, dep)

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := g.Run(":" + cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
