package routers

import (
	"github.com/bermanbenjamin/futStats/cmd/api/middlewares"
	"github.com/bermanbenjamin/futStats/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, dependencies *config.Dependencies) {
	v1 := router.Group("/api/v1")
	{

		basePath := v1.Group("")
		basePath.Use(middlewares.AuthMiddleware)
		{
			basePath.GET("/leagues/:id", dependencies.LeagueHandler.GetLeagueById)
			basePath.POST("/leagues", dependencies.LeagueHandler.CreateLeague)
			basePath.PUT("/leagues", dependencies.LeagueHandler.UpdateLeague)
			basePath.DELETE("/leagues/:id", dependencies.LeagueHandler.DeleteLeague)
		}

		organization := v1.Group("/:slug")
		organization.Use(middlewares.AuthMiddleware)
		{

			organization.GET("/players/:id", dependencies.PlayerHandler.GetPlayer)
			organization.GET("/players", dependencies.PlayerHandler.GetAllPlayers)
			organization.POST("/players", dependencies.PlayerHandler.CreatePlayer)
			organization.PUT("/players", dependencies.PlayerHandler.UpdatePlayer)
			organization.DELETE("/players/:id", dependencies.PlayerHandler.DeletePlayerById)

			organization.GET("/events/:id", dependencies.EventHandler.GetEventByPlayerId)
			organization.POST("/events", dependencies.EventHandler.CreateEvent)
			organization.PUT("/events", dependencies.EventHandler.UpdateEvent)
			organization.DELETE("/events/:id", dependencies.EventHandler.DeleteEvent)

		}
		auth := v1.Group("/auth")
		{
			auth.POST("/login", dependencies.AuthHandler.Login)
			auth.POST("/logout", dependencies.AuthHandler.Logout)
			auth.POST("/signup", dependencies.AuthHandler.SignUp)
		}

	}

}
