package config

import (
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/bermanbenjamin/futStats/internal/service"
	handler "github.com/bermanbenjamin/futStats/internal/transport/http/handlers"
	"gorm.io/gorm"
)

type Dependencies struct {
	PlayerRepository *repository.PlayerRepository
	PlayerService    *service.PlayerService
	PlayerHandler    *handler.PlayerHandler
}

// InitializeDependencies sets up all dependencies like repositories, services, and handlers
func InitializeDependencies(db *gorm.DB) *Dependencies {
	// Initialize repositories
	playerRepo := repository.NewPlayerRepository(db)

	// Initialize services
	playerService := service.NewPlayerService(playerRepo)

	// Initialize HTTP handlers
	playerHandler := handler.NewPlayerHandler(playerService)

	// Return a struct containing all dependencies
	return &Dependencies{
		PlayerRepository: playerRepo,
		PlayerService:    playerService,
		PlayerHandler:    playerHandler,
	}
}
