package service

import (
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type PlayerService struct {
	repo *repository.PlayerRepository
}

func NewPlayerService(repo *repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) GetAllPlayers() ([]models.Player, error) {
	return s.repo.GetAllPlayers()
}

func (s *PlayerService) GetPlayer(id uuid.UUID) (*models.Player, error) {
	return s.repo.GetPlayerById(id)
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
