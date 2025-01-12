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
	EventRepository *repository.EventsRepository
	EventService    *service.EventService
	EventHandler    *handler.EventsHandler
}

// InitializeDependencies sets up all dependencies like repositories, services, and handlers
func InitializeDependencies(db *gorm.DB) *Dependencies {
	// Initialize repositories
	playerRepo := repository.NewPlayerRepository(db)
	eventRepo := repository.NewEventsRepository(db)

	// Initialize services
	playerService := service.NewPlayerService(playerRepo)
	eventService := service.NewEventService(eventRepo, playerService)

	// Initialize HTTP handlers
	playerHandler := handler.NewPlayerHandler(playerService)
	eventHandler := handler.NewEventsHandler(eventService)

	// Return a struct containing all dependencies
	return &Dependencies{
		PlayerRepository: playerRepo,
		PlayerService:    playerService,
		PlayerHandler:    playerHandler,
		EventRepository: eventRepo,
        EventService:    eventService,
        EventHandler:    eventHandler,
	}
}
