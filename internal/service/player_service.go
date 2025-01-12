package service

import (
	"errors"

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

func (s *PlayerService) UpdatePlayerByEvent(event model.Event, isCreateEvent bool) (*model.Player, error) {
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
		default: return nil, errors.New("Unknown event type for event " + event.Type)
        
	}

    return s.repo.UpdatePlayer(player)
}
