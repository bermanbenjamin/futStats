package main

import (
	controllers "github.com/bermanbenjamin/futStats/controllers/players"
	"github.com/bermanbenjamin/futStats/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	g := gin.Default()

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

		g.Run(":8080")
	}
}
