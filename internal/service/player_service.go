package service

import (
	model "github.com/bermanbenjamin/futStats/internal/models"
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

func (s *PlayerService) GetAllPlayers(filterQuery string, filterValue string) ([]*model.Player, error) {
	return s.repo.GetAllPlayersBy(filterQuery, filterValue)
}

func (s *PlayerService) GetPlayer(id uuid.UUID) (*model.Player, error) {
	return s.repo.GetPlayerById(id)
}

func (s *PlayerService) CreatePlayer(player *model.Player) (*model.Player, error) {
	return s.repo.AddPlayer(*player)
}

func (s *PlayerService) UpdatePlayer(player *model.Player) (*model.Player, error) {
	return s.repo.UpdatePlayer(*player)
}

func (s *PlayerService) DeletePlayer(id uuid.UUID) error {
	return s.repo.DeletePlayer(id)
}

func (s *PlayerService) UpdatePlayerByEvent(event model.Event) (*model.Player, error) {
	player := event.Player

	switch event.Type {
		case enums.Goal:
            player.Goals++
        case enums.Assist:
            player.Assists++
        case enums.Dribble:
            player.Dribbles++
        case enums.Disarm:
            player.Disarms++
        default:
            return nil, nil // Invalid event type, do nothing with the player's stats.
	}

    return s.repo.UpdatePlayer(player)
}
