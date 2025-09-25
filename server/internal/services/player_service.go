package services

import (
	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type PlayerService struct {
	repo               repository.PlayerRepositoryInterface
	playerStatsService *PlayerStatsService
}

func NewPlayerService(repo repository.PlayerRepositoryInterface, eventRepo *repository.EventsRepository) *PlayerService {
	return &PlayerService{
		repo:               repo,
		playerStatsService: NewPlayerStatsService(eventRepo),
	}
}

func (s *PlayerService) GetAllPlayers(filterQuery constants.QueryFilter, filterValue string) ([]*models.Player, error) {
	players, err := s.repo.GetAllPlayersBy(filterQuery, filterValue)
	if err != nil {
		return nil, err
	}

	// Populate stats for all players
	if err := s.playerStatsService.GetPlayersWithStats(players); err != nil {
		return nil, err
	}

	return players, nil
}

func (s *PlayerService) GetPlayerBy(filter constants.QueryFilter, value string) (*models.Player, error) {
	player, err := s.repo.GetPlayerBy(filter, value)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, nil
	}

	// Populate stats for the player
	if err := s.playerStatsService.GetPlayerStats(player); err != nil {
		return nil, err
	}

	return player, nil
}

func (s *PlayerService) CreatePlayer(player *models.Player) (*models.Player, error) {
	return s.repo.CreatePlayer(player)
}

func (s *PlayerService) UpdatePlayer(player *models.Player) (*models.Player, error) {
	return s.repo.UpdatePlayer(player)
}

func (s *PlayerService) DeletePlayer(id uuid.UUID) error {
	return s.repo.DeletePlayer(id)
}
