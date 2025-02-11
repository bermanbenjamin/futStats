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
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Create the league record
		if err := tx.Create(league).Error; err != nil {
			log.Printf("Error creating league record: %v", err)
			return err
		}

		log.Printf("League created: %v", league)

		// Add the owner as a member of the league
		if err := tx.Model(league).Association("Members").Append(&owner); err != nil {
			log.Printf("Error adding owner as member: %v", err)
			return err
		}

		log.Printf("Owner added as member: %v", owner.ID)

		// Load the complete league data with associations
		if err := tx.First(league, league.ID).Error; err != nil {
			log.Printf("Error loading league data: %v", err)
			return err
		}

		log.Printf("League loaded: %v", league)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create league: %w", err)
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
	// First get the league
	league = &models.League{}
	if err := r.db.First(league, leagueId).Error; err != nil {
		return nil, fmt.Errorf("league not found: %w", err)
	}

	// Add the player as a member
	if err := r.db.Model(league).Association("Members").Append(&models.Player{Base: models.Base{ID: playerId}}); err != nil {
		return nil, fmt.Errorf("failed to add member: %w", err)
	}

	// Reload the league with all associations
	if err := r.db.Preload("Owner").Preload("Members").First(league, leagueId).Error; err != nil {
		return nil, fmt.Errorf("failed to load league data: %w", err)
	}

	return league, nil
}
