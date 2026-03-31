package services

import (
	"errors"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type SeasonService struct {
	repo       *repository.SeasonRepository
	leagueRepo *repository.LeagueRepository
}

func NewSeasonService(repo *repository.SeasonRepository, leagueRepo *repository.LeagueRepository) *SeasonService {
	return &SeasonService{repo: repo, leagueRepo: leagueRepo}
}

func (s *SeasonService) CreateSeason(season *models.Season, leagueSlug string) (*models.Season, error) {
	league, err := s.leagueRepo.GetLeagueBy(constants.SLUG, leagueSlug)
	if err != nil {
		return nil, errors.New("league not found")
	}
	return s.repo.CreateSeason(season, league.ID)
}

func (s *SeasonService) GetSeasonsByLeagueSlug(slug string) ([]models.Season, error) {
	league, err := s.leagueRepo.GetLeagueBy(constants.SLUG, slug)
	if err != nil {
		return nil, errors.New("league not found")
	}
	return s.repo.GetSeasonsByLeagueId(league.ID)
}

func (s *SeasonService) GetSeasonById(id uuid.UUID) (*models.Season, error) {
	return s.repo.GetSeasonById(id)
}
