package repository

import (
	"fmt"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeagueRepository struct {
	db *gorm.DB
}

func NewLeagueRepository(db *gorm.DB) *LeagueRepository {
	return &LeagueRepository{db: db}
}

func (r *LeagueRepository) CreateLeague(league *models.League) (createdLeague *models.League, err error) {
	if err := r.db.Create(league).Error; err != nil {
		return nil, err
	}
	return league, nil
}

func (r *LeagueRepository) GetLeagueBy(query constants.QueryFilter, values string) (leagues *models.League, err error) {
	var league *models.League
	queryString := fmt.Sprintf("%s =?", query)
	if err := r.db.Where(queryString, values).First(&league).Error; err != nil {
		return nil, err
	}
	return league, nil
}

func (r *LeagueRepository) UpdateLeague(league *models.League) (updatedLeague *models.League, err error) {
	if err := r.db.Model(&models.League{}).Where("id =?", league.ID).Updates(league).Error; err != nil {
		return nil, err
	}
	return league, nil
}

func (r *LeagueRepository) DeleteLeague(id uuid.UUID) error {
	if err := r.db.Where("id =?", id).Delete(&models.League{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *LeagueRepository) AddPlayerToLeague(playerId uuid.UUID, leagueId uuid.UUID) (league *models.League, err error) {
	if err := r.db.Model(&models.League{}).Where("id = ?", leagueId).
		Association("Players").Append(&models.Player{Base: models.Base{ID: playerId}}); err != nil {
		return nil, err
	}

	league, err = r.GetLeagueBy(constants.ID, leagueId.String())
	if err != nil {
		return nil, err
	}

	return league, nil
}
