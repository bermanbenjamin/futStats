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
	// Add the owner to the Members slice
	league.Members = append(league.Members, *owner)

	// Create the league record along with the owner as a member
	if err := r.db.Create(league).Error; err != nil {
		log.Printf("Error creating league record: %v", err)
		return nil, err
	}

	// Load the complete league data with associations
	if err := r.db.Preload("Members").First(league, league.ID).Error; err != nil {
		log.Printf("Error loading league data: %v", err)
		return nil, err
	}

	log.Printf("League created and owner added as member: %v", league)

	return league, nil
}

func (r *LeagueRepository) GetLeagueBy(query constants.QueryFilter, value string) (leagues *models.League, err error) {
	var league *models.League
	queryString := fmt.Sprintf("%s =?", query)

	// Use Preload to load all associations
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

func (r *LeagueRepository) AddPlayerToLeague(playerId uuid.UUID, leagueId uuid.UUID) (league *models.League, err error) {
	// First get the league
	league = &models.League{}
	if err := r.db.Joins("JOIN league_members ON league_members.league_id = leagues.id").First(league, leagueId).Error; err != nil {
		return nil, fmt.Errorf("league not found: %w", err)
	}

	// Add the player as a member using the existing player ID
	player := &models.Player{Base: models.Base{ID: playerId}} // Ensure we are using the existing player
	if err := r.db.Model(league).Association("Members").Append(player); err != nil {
		return nil, fmt.Errorf("failed to add member: %w", err)
	}

	// Reload the league with all associations using joins
	if err := r.db.Joins("JOIN owners ON owners.id = leagues.owner_id").Joins("JOIN league_members ON league_members.league_id = leagues.id").First(league, leagueId).Error; err != nil {
		return nil, fmt.Errorf("failed to load league data: %w", err)
	}

	return league, nil
}
