package http

import (
	"github.com/bermanbenjamin/futStats/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, dependencies *config.Dependencies) {
	api := router.Group("/api/v1")
	{
		api.GET("/players/:id", dependencies.PlayerHandler.GetPlayer)
		api.GET("/players", dependencies.PlayerHandler.GetAllPlayers)
		api.POST("/players", dependencies.PlayerHandler.CreatePlayer)
		api.PUT("/players", dependencies.PlayerHandler.UpdatePlayer)
		api.DELETE("/players/:id", dependencies.PlayerHandler.DeletePlayerById)
	}
}
