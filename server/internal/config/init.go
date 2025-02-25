package config

import (
	"github.com/bermanbenjamin/futStats/cmd/api/handlers"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/bermanbenjamin/futStats/internal/services"
	"gorm.io/gorm"
)

type Dependencies struct {
	PlayerRepository *repository.PlayerRepository
	PlayerService    *services.PlayerService
	PlayerHandler    *handlers.PlayerHandler
	EventRepository  *repository.EventsRepository
	EventService     *services.EventService
	EventHandler     *handlers.EventsHandler
	AuthHandler      *handlers.AuthHandler
	AuthService      *services.AuthService
	LeagueRepository *repository.LeagueRepository
	LeagueService    *services.LeagueService
	LeagueHandler    *handlers.LeagueHandler
}

func InitializeDependencies(db *gorm.DB) *Dependencies {
	playerRepo := repository.NewPlayerRepository(db)
	eventRepo := repository.NewEventsRepository(db)
	leagueRepo := repository.NewLeagueRepository(db)

	playerService := services.NewPlayerService(playerRepo)
	eventService := services.NewEventService(eventRepo, playerService)
	authService := services.NewAuthService(playerService)
	leagueService := services.NewLeagueService(leagueRepo, playerService)

	playerHandler := handlers.NewPlayerHandler(playerService)
	eventHandler := handlers.NewEventsHandler(eventService)
	authHandler := handlers.NewAuthHandler(authService)
	leagueHandler := handlers.NewLeagueHandler(leagueService, playerService)

	return &Dependencies{
		PlayerRepository: playerRepo,
		PlayerService:    playerService,
		PlayerHandler:    playerHandler,
		EventRepository:  eventRepo,
		EventService:     eventService,
		EventHandler:     eventHandler,
		AuthHandler:      authHandler,
		AuthService:      authService,
		LeagueRepository: leagueRepo,
		LeagueService:    leagueService,
		LeagueHandler:    leagueHandler,
	}
}
