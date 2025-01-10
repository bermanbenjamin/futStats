package service

import (
	model "github.com/bermanbenjamin/futStats/internal/models"
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
