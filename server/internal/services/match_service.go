package services

import (
	"errors"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type MatchService struct {
	repo       *repository.MatchRepository
	leagueRepo *repository.LeagueRepository
}

func NewMatchService(repo *repository.MatchRepository, leagueRepo *repository.LeagueRepository) *MatchService {
	return &MatchService{repo: repo, leagueRepo: leagueRepo}
}

func (s *MatchService) CreateMatch(match *models.Match) (*models.Match, error) {
	_, err := s.leagueRepo.GetLeagueBy(constants.ID, match.LeagueId.String())
	if err != nil {
		return nil, errors.New("league not found")
	}
	return s.repo.CreateMatch(match)
}

func (s *MatchService) GetMatchById(id uuid.UUID) (*models.Match, error) {
	return s.repo.GetMatchById(id)
}

func (s *MatchService) GetMatchesByLeagueSlug(slug string) ([]models.Match, error) {
	league, err := s.leagueRepo.GetLeagueBy(constants.SLUG, slug)
	if err != nil {
		return nil, errors.New("league not found")
	}
	return s.repo.GetMatchesByLeagueId(league.ID)
}

func (s *MatchService) DeleteMatch(id uuid.UUID) error {
	return s.repo.DeleteMatch(id)
}
