package main

import (
	"log"
	"os"
	"strings"

	routers "github.com/bermanbenjamin/futStats/cmd/api"
	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/config"
	"github.com/bermanbenjamin/futStats/internal/db"
	"github.com/bermanbenjamin/futStats/internal/logger"
	"github.com/bermanbenjamin/futStats/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	loggerConfig := logger.Config{
		Level:  cfg.LogLevel,
		Format: cfg.LogFormat,
	}

	err = logger.InitGlobal(loggerConfig)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Get global logger instance
	appLogger := logger.GetGlobal()
	defer appLogger.Sync()

	appLogger.Info("Starting FutStats server",
		zap.String("environment", cfg.Environment),
		zap.String("log_level", cfg.LogLevel),
		zap.String("log_format", cfg.LogFormat),
	)

	err = db.InitDB(cfg.DatabaseUrl)
	if err != nil {
		appLogger.Fatal("Failed to initialize database", zap.Error(err))
	}

	dep := config.InitializeDependencies(db.DB)

	g := gin.Default()

	// Add logging middleware
	g.Use(middlewares.LoggingMiddleware())

	// Custom CORS middleware
	g.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if origin is allowed
		allowed := false

		// Allow localhost for development
		if origin == "http://localhost:3000" {
			allowed = true
		}

		// Get allowed origins from environment variables
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
		if allowedOrigins != "" {
			origins := strings.Split(allowedOrigins, ",")
			for _, allowedOrigin := range origins {
				allowedOrigin = strings.TrimSpace(allowedOrigin)
				if origin == allowedOrigin {
					allowed = true
					break
				}
			}
		}

		// Allow any Vercel deployment from your project (fallback)
		if !allowed && origin != "" &&
			(strings.HasPrefix(origin, "https://client-") &&
				strings.HasSuffix(origin, "-bermanbenjamins-projects.vercel.app")) {
			allowed = true
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

	appLogger.Info("Custom CORS middleware configured",
		zap.String("localhost_allowed", "http://localhost:3000"),
		zap.String("vercel_pattern", "https://client-*-bermanbenjamins-projects.vercel.app"),
	)

	if allowedOrigins := os.Getenv("ALLOWED_ORIGINS"); allowedOrigins != "" {
		appLogger.Info("Additional CORS origins configured", zap.String("origins", allowedOrigins))
	}

	routers.SetupRouter(g, dep)

	appLogger.Info("Starting server", zap.String("address", ":"+cfg.ServerAddress))
	if err := g.Run(":" + cfg.ServerAddress); err != nil {
		appLogger.Fatal("Failed to start server", zap.Error(err))
	}
}
