package routers

import (
	"github.com/bermanbenjamin/futStats/internal/config"
	"github.com/bermanbenjamin/futStats/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, dependencies *config.Dependencies) {
	// Health check endpoint for Railway
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "futStats API",
			"version": "1.0.0",
		})
	})

	v1 := router.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			// Handle OPTIONS requests for CORS preflight
			auth.OPTIONS("/login", func(c *gin.Context) {
				c.Status(200)
			})
			auth.OPTIONS("/logout", func(c *gin.Context) {
				c.Status(200)
			})
			auth.OPTIONS("/signup", func(c *gin.Context) {
				c.Status(200)
			})

			auth.POST("/login", dependencies.AuthHandler.Login)
			auth.POST("/logout", dependencies.AuthHandler.Logout)
			auth.POST("/signup", dependencies.AuthHandler.SignUp)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middlewares.AuthMiddleware)
		{
			// League owner routes
			leagues := protected.Group("/leagues")
			leagues.Use(func(ctx *gin.Context) {
				middlewares.OwnerMiddleware(ctx, dependencies)
			})
			{
				leagues.PUT("", dependencies.LeagueHandler.UpdateLeague)
				leagues.DELETE("/:leagueSlug", dependencies.LeagueHandler.DeleteLeague)
				leagues.POST("/:leagueSlug/players", dependencies.LeagueHandler.AddPlayerToLeague)
			}

			// General league routes (unprotected)
			leaguesUnprotected := v1.Group("/leagues")
			{
				leaguesUnprotected.GET("/:leagueSlug", dependencies.LeagueHandler.GetLeagueBySlug)
				leaguesUnprotected.POST("", dependencies.LeagueHandler.CreateLeague)
			}

			// Player routes
			players := protected.Group("/players")
			{
				players.GET("/:id", dependencies.PlayerHandler.GetPlayerBy)
				players.GET("", dependencies.PlayerHandler.GetAllPlayers)
				players.POST("", dependencies.PlayerHandler.CreatePlayer)
				players.PUT("", dependencies.PlayerHandler.UpdatePlayer)
				players.DELETE("/:id", dependencies.PlayerHandler.DeletePlayerById)
			}

			// Event routes
			events := protected.Group("/events")
			{
				events.GET("/:id", dependencies.EventHandler.GetEventByPlayerId)
				events.POST("", dependencies.EventHandler.CreateEvent)
				events.PUT("", dependencies.EventHandler.UpdateEvent)
				events.DELETE("/:id", dependencies.EventHandler.DeleteEvent)
			}
		}
	}
}
