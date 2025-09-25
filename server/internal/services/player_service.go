package services

import (
	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/logger"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PlayerService struct {
	repo               repository.PlayerRepositoryInterface
	playerStatsService *PlayerStatsService
	logger             *logger.Logger
}

func NewPlayerService(repo repository.PlayerRepositoryInterface, eventRepo *repository.EventsRepository) *PlayerService {
	return &PlayerService{
		repo:               repo,
		playerStatsService: NewPlayerStatsService(eventRepo),
		logger:             logger.GetGlobal().WithComponent("player_service"),
	}
}

func (s *PlayerService) GetAllPlayers(filterQuery constants.QueryFilter, filterValue string) ([]*models.Player, error) {
	s.logger.Info("Getting all players",
		zap.String("filter", string(filterQuery)),
		zap.String("value", filterValue))

	players, err := s.repo.GetAllPlayersBy(filterQuery, filterValue)
	if err != nil {
		s.logger.Error("Failed to get players from repository",
			zap.String("filter", string(filterQuery)),
			zap.String("value", filterValue),
			zap.Error(err))
		return nil, err
	}

	// Populate stats for all players
	if err := s.playerStatsService.GetPlayersWithStats(players); err != nil {
		s.logger.Error("Failed to populate player stats",
			zap.Int("player_count", len(players)),
			zap.Error(err))
		return nil, err
	}

	s.logger.Info("Successfully retrieved players",
		zap.Int("count", len(players)),
		zap.String("filter", string(filterQuery)))

	return players, nil
}

func (s *PlayerService) GetPlayerBy(filter constants.QueryFilter, value string) (*models.Player, error) {
	s.logger.Info("Getting player by filter",
		zap.String("filter", string(filter)),
		zap.String("value", value))

	player, err := s.repo.GetPlayerBy(filter, value)
	if err != nil {
		s.logger.Error("Failed to get player from repository",
			zap.String("filter", string(filter)),
			zap.String("value", value),
			zap.Error(err))
		return nil, err
	}
	if player == nil {
		s.logger.Warn("Player not found",
			zap.String("filter", string(filter)),
			zap.String("value", value))
		return nil, nil
	}

	// Populate stats for the player
	if err := s.playerStatsService.GetPlayerStats(player); err != nil {
		s.logger.Error("Failed to populate player stats",
			zap.String("player_id", player.ID.String()),
			zap.Error(err))
		return nil, err
	}

	s.logger.Info("Successfully retrieved player",
		zap.String("player_id", player.ID.String()),
		zap.String("name", player.Name),
		zap.String("email", player.Email))

	return player, nil
}

func (s *PlayerService) CreatePlayer(player *models.Player) (*models.Player, error) {
	s.logger.Info("Creating new player",
		zap.String("name", player.Name),
		zap.String("email", player.Email),
		zap.Int("age", player.Age))

	createdPlayer, err := s.repo.CreatePlayer(player)
	if err != nil {
		s.logger.Error("Failed to create player",
			zap.String("name", player.Name),
			zap.String("email", player.Email),
			zap.Error(err))
		return nil, err
	}

	s.logger.LogBusinessEvent("player_created", "player", createdPlayer.ID.String(), map[string]interface{}{
		"name":  createdPlayer.Name,
		"email": createdPlayer.Email,
		"age":   createdPlayer.Age,
	})

	return createdPlayer, nil
}

func (s *PlayerService) UpdatePlayer(player *models.Player) (*models.Player, error) {
	s.logger.Info("Updating player",
		zap.String("player_id", player.ID.String()),
		zap.String("name", player.Name),
		zap.String("email", player.Email))

	updatedPlayer, err := s.repo.UpdatePlayer(player)
	if err != nil {
		s.logger.Error("Failed to update player",
			zap.String("player_id", player.ID.String()),
			zap.Error(err))
		return nil, err
	}

	s.logger.LogBusinessEvent("player_updated", "player", updatedPlayer.ID.String(), map[string]interface{}{
		"name":  updatedPlayer.Name,
		"email": updatedPlayer.Email,
		"age":   updatedPlayer.Age,
	})

	return updatedPlayer, nil
}

func (s *PlayerService) DeletePlayer(id uuid.UUID) error {
	s.logger.Info("Deleting player", zap.String("player_id", id.String()))

	err := s.repo.DeletePlayer(id)
	if err != nil {
		s.logger.Error("Failed to delete player",
			zap.String("player_id", id.String()),
			zap.Error(err))
		return err
	}

	s.logger.LogBusinessEvent("player_deleted", "player", id.String(), nil)
	return nil
}
