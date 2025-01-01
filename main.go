package main

import (
	"os"

	controllers "github.com/bermanbenjamin/futStats/controllers/players"
	"github.com/bermanbenjamin/futStats/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	g := gin.Default()

	db.Init()                 // Use env vars for DB connection
	port := os.Getenv("PORT") // Default to 8080 if unset
	if port == "" {
		port = "8080"
	}

	{
		v1 := g.Group("/api/v1")
		var players = v1.Group("/players")
		{
			players.GET("/", controllers.GetPlayers)
			players.GET("/:id", controllers.GetPlayer)
			players.POST("/", controllers.CreatePlayer)
			players.PUT("/:id", controllers.UpdatePlayer)
			players.DELETE("/:id", controllers.DeletePlayer)
		}

	}

	// Routes setup
	g.Run(":" + port)
}
