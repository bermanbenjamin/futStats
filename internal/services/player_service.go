package services

import (
	"errors"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/models/enums"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type PlayerService struct {
	repo *repository.PlayerRepository
}

func NewPlayerService(repo *repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) GetAllPlayers(filterQuery constants.QueryFilter, filterValue string) ([]*models.Player, error) {
	return s.repo.GetAllPlayersBy(filterQuery, filterValue)
}

func (s *PlayerService) GetPlayerBy(filter constants.QueryFilter, value string) (*models.Player, error) {
	return s.repo.GetPlayerBy(filter, value)
}

func (s *PlayerService) CreatePlayer(player *models.Player) (*models.Player, error) {
	return s.repo.AddPlayer(*player)
}

func (s *PlayerService) UpdatePlayer(player *models.Player) (*models.Player, error) {
	return s.repo.UpdatePlayer(*player)
}

func (s *PlayerService) DeletePlayer(id uuid.UUID) error {
	return s.repo.DeletePlayer(id)
}

func (s *PlayerService) UpdatePlayerByEvent(event models.Event, isCreateEvent bool) (*models.Player, error) {
	player := event.Player

	switch event.Type {
	case enums.Assist:
		if isCreateEvent {
			player.Assists++
		} else {
			player.Assists--
		}
	case enums.Disarm:
		if isCreateEvent {
			player.Disarms++
		} else {
			player.Disarms--
		}
	case enums.Dribble:
		if isCreateEvent {
			player.Dribbles++
		} else {
			player.Dribbles--
		}
	case enums.Goal:
		if isCreateEvent {
			player.Matches++
		} else {
			player.Matches--
		}
	default:
		return nil, errors.New("Unknown event type for event " + event.Type)

	}

	return s.repo.UpdatePlayer(player)
}
