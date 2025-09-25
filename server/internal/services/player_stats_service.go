package services

import (
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

// PlayerStatsService handles player statistics using normalized queries
type PlayerStatsService struct {
	eventRepo *repository.EventsRepository
}

func NewPlayerStatsService(eventRepo *repository.EventsRepository) *PlayerStatsService {
	return &PlayerStatsService{
		eventRepo: eventRepo,
	}
}

// GetPlayerStats populates player statistics using COUNT queries
func (s *PlayerStatsService) GetPlayerStats(player *models.Player) error {
	playerID := player.ID

	// Count goals
	goals, err := s.eventRepo.CountEventsByPlayerAndType(playerID, "Goal")
	if err != nil {
		return err
	}
	player.Goals = goals

	// Count assists
	assists, err := s.eventRepo.CountEventsByPlayerAndType(playerID, "Assist")
	if err != nil {
		return err
	}
	player.Assists = assists

	// Count disarms
	disarms, err := s.eventRepo.CountEventsByPlayerAndType(playerID, "Disarm")
	if err != nil {
		return err
	}
	player.Disarms = disarms

	// Count dribbles
	dribbles, err := s.eventRepo.CountEventsByPlayerAndType(playerID, "Dribble")
	if err != nil {
		return err
	}
	player.Dribbles = dribbles

	// Count matches
	matches, err := s.eventRepo.CountEventsByPlayerAndType(playerID, "Match")
	if err != nil {
		return err
	}
	player.Matches = matches

	// Count red cards
	redCards, err := s.eventRepo.CountEventsByPlayerAndType(playerID, "RedCard")
	if err != nil {
		return err
	}
	player.RedCards = redCards

	// Count yellow cards
	yellowCards, err := s.eventRepo.CountEventsByPlayerAndType(playerID, "YellowCard")
	if err != nil {
		return err
	}
	player.YellowCards = yellowCards

	return nil
}

// GetPlayerStatsByType gets a specific stat count for a player
func (s *PlayerStatsService) GetPlayerStatsByType(playerID uuid.UUID, eventType string) (int, error) {
	return s.eventRepo.CountEventsByPlayerAndType(playerID, eventType)
}

// GetPlayersWithStats gets multiple players with their statistics populated
func (s *PlayerStatsService) GetPlayersWithStats(players []*models.Player) error {
	for _, player := range players {
		if err := s.GetPlayerStats(player); err != nil {
			return err
		}
	}
	return nil
}
