package repository

import (
	"fmt"
	"log"

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

func (r *LeagueRepository) CreateLeague(league *models.League, owner *models.Player) (*models.League, error) {
	league.Members = append(league.Members, *owner)

	if err := r.db.Create(league).Error; err != nil {
		log.Printf("Error creating league record: %v", err)
		return nil, err
	}

	if err := r.db.Preload("Members").First(league, league.ID).Error; err != nil {
		log.Printf("Error loading league data: %v", err)
		return nil, err
	}

	return league, nil
}

func (r *LeagueRepository) GetLeagueBy(query constants.QueryFilter, value string) (leagues *models.League, err error) {
	var league *models.League
	queryString := fmt.Sprintf("%s =?", query)

	if err := r.db.Preload("Owner").Preload("Members").Where(queryString, value).First(&league).Error; err != nil {
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

func (r *LeagueRepository) AddPlayerToLeague(player *models.Player, league *models.League) (*models.League, error) {

	if err := r.db.Model(league).Association("Members").Append(player); err != nil {
		return nil, fmt.Errorf("failed to add member: %w", err)
	}

	if err := r.db.Preload("Members").First(league, league.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to load updated league data: %w", err)
	}

	return league, nil
}
