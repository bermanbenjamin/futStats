package main

import (
	"log"
	"time"

	routers "github.com/bermanbenjamin/futStats/cmd/api"
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
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true, // Permite cookies e cabeçalhos de autenticação
		MaxAge:           12 * time.Hour,
	}
	g.Use(cors.New(config))

	routers.SetupRouter(g, dep)

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := g.Run(":" + cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
