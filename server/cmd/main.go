package main

import (
	"log"

	routers "github.com/bermanbenjamin/futStats/cmd/api"
	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/config"
	"github.com/bermanbenjamin/futStats/internal/db"
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

	// Custom CORS middleware
	g.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		allowedOrigins := []string{
			"http://localhost:3000",
			"https://client-jlcuuxswn-bermanbenjamins-projects.vercel.app",
		}

		// Check if origin is allowed
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With, "+constants.QUERY_FILTER)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "43200") // 12 hours

		// Handle preflight OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	log.Printf("Custom CORS middleware configured")
	log.Printf("- Allowed Origins: http://localhost:3000, https://client-j1b0qb677-bermanbenjamins-projects.vercel.app")

	routers.SetupRouter(g, dep)

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := g.Run(":" + cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
