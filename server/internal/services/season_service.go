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

func (s *SeasonService) GetSeasonStats(seasonId uuid.UUID) ([]*models.Player, error) {
	return s.repo.GetSeasonStats(seasonId)
}

func (s *SeasonService) FinishSeason(leagueSlug string, seasonId uuid.UUID) (*models.Season, error) {
	league, err := s.leagueRepo.GetLeagueBy(constants.SLUG, leagueSlug)
	if err != nil {
		return nil, errors.New("league not found")
	}

	seasons, err := s.repo.GetSeasonsByLeagueId(league.ID)
	if err != nil {
		return nil, errors.New("could not retrieve seasons for league")
	}

	found := false
	for _, season := range seasons {
		if season.ID == seasonId {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("season does not belong to the specified league")
	}

	return s.repo.FinishSeason(seasonId)
}
