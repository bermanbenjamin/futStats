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
	MatchRepository  *repository.MatchRepository
	MatchService     *services.MatchService
	MatchHandler     *handlers.MatchHandler
	SeasonRepository *repository.SeasonRepository
	SeasonService    *services.SeasonService
	SeasonHandler    *handlers.SeasonHandler
}

func InitializeDependencies(db *gorm.DB) *Dependencies {
	playerRepo := repository.NewPlayerRepository(db)
	eventRepo := repository.NewEventsRepository(db)
	leagueRepo := repository.NewLeagueRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	seasonRepo := repository.NewSeasonRepository(db)

	// Inject dependencies - normalized approach
	playerService := services.NewPlayerService(playerRepo, eventRepo)
	eventService := services.NewEventService(eventRepo, matchRepo)
	authService := services.NewAuthService(playerService)
	leagueService := services.NewLeagueService(leagueRepo, playerService)
	matchService := services.NewMatchService(matchRepo, leagueRepo)
	seasonService := services.NewSeasonService(seasonRepo, leagueRepo)

	playerHandler := handlers.NewPlayerHandler(playerService)
	eventHandler := handlers.NewEventsHandler(eventService)
	authHandler := handlers.NewAuthHandler(authService)
	leagueHandler := handlers.NewLeagueHandler(leagueService, playerService)
	matchHandler := handlers.NewMatchHandler(matchService)
	seasonHandler := handlers.NewSeasonHandler(seasonService)

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
		MatchRepository:  matchRepo,
		MatchService:     matchService,
		MatchHandler:     matchHandler,
		SeasonRepository: seasonRepo,
		SeasonService:    seasonService,
		SeasonHandler:    seasonHandler,
	}
}
