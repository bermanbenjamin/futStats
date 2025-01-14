package routers

import (
	"github.com/bermanbenjamin/futStats/internal/config"
	middlewares "github.com/bermanbenjamin/futStats/internal/transport/http/middlerware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, dependencies *config.Dependencies) {
	v1 := router.Group("/api/v1")
	{
		protected := v1.Group("/:slug")
		protected.Use(middlewares.AuthMiddleware)
		{

			protected.GET("/players/:id", dependencies.PlayerHandler.GetPlayer)
			protected.GET("/players", dependencies.PlayerHandler.GetAllPlayers)
			protected.POST("/players", dependencies.PlayerHandler.CreatePlayer)
			protected.PUT("/players", dependencies.PlayerHandler.UpdatePlayer)
			protected.DELETE("/players/:id", dependencies.PlayerHandler.DeletePlayerById)

			protected.GET("/events/:id", dependencies.EventHandler.GetEventByPlayerId)
			protected.POST("/events", dependencies.EventHandler.CreateEvent)
			protected.PUT("/events", dependencies.EventHandler.UpdateEvent)
			protected.DELETE("/events/:id", dependencies.EventHandler.DeleteEvent)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/login", dependencies.AuthHandler.Login)
			auth.POST("/logout", dependencies.AuthHandler.Logout)
			auth.POST("/signup", dependencies.AuthHandler.SignUp)
		}

	}

}
