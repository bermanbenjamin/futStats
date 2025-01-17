package services

import (
	"errors"
	"log"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/commons"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type LeagueService struct {
	repo          *repository.LeagueRepository
	playerService *PlayerService
}

func NewLeagueService(repo *repository.LeagueRepository, playerService *PlayerService) *LeagueService {
	return &LeagueService{repo: repo, playerService: playerService}
}

func (s *LeagueService) CreateLeague(league *models.League) (*models.League, error) {
	player, err := s.playerService.GetPlayerBy(constants.ID, league.OwnerId.String())

	if err != nil {
		return nil, errors.New("could not find player with id")
	}

	log.Printf("player with id %s already exists", league.OwnerId.String())

	if league.Slug == "" {
		league.Slug = commons.Slugify(league.Name)
	}

	league, err = s.repo.CreateLeague(league)

	if err != nil {
		return nil, err
	}

	return s.AddPlayerToLeague(player.ID, league.ID)
}

func (s *LeagueService) GetLeagueBy(query constants.QueryFilter, values string) (*models.League, error) {
	return s.repo.GetLeagueBy(query, values)
}

func (s *LeagueService) UpdateLeague(league *models.League) (updated *models.League, err error) {
	return s.repo.UpdateLeague(league)
}

func (s *LeagueService) DeleteLeague(league *models.League) error {
	return s.repo.DeleteLeague(league.ID)
}

func (s *LeagueService) AddPlayerToLeague(playerId uuid.UUID, leagueId uuid.UUID) (league *models.League, err error) {
	_, err = s.repo.GetLeagueBy(constants.ID, leagueId.String())

	if err != nil {
		return nil, errors.New("could not find league with id")
	}

	_, err = s.playerService.GetPlayerBy(constants.ID, playerId.String())

	if err != nil {
		return nil, errors.New("could not find player with id")
	}

	return s.repo.AddPlayerToLeague(playerId, leagueId)
}
