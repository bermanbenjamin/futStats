package repository

import (
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) CreateMatch(match *models.Match) (*models.Match, error) {
	if err := r.db.Create(match).Error; err != nil {
		log.Printf("Error creating match: %v", err)
		return nil, err
	}
	if err := r.db.Preload("League").Preload("Events").First(match, match.ID).Error; err != nil {
		log.Printf("Error loading match after create: %v", err)
		return nil, err
	}
	return match, nil
}

func (r *MatchRepository) GetMatchById(id uuid.UUID) (*models.Match, error) {
	var match models.Match
	if err := r.db.Preload("League").Preload("Events").Where("id = ?", id).First(&match).Error; err != nil {
		log.Printf("Error getting match with ID %v: %v", id, err)
		return nil, err
	}
	return &match, nil
}

func (r *MatchRepository) GetMatchesByLeagueId(leagueId uuid.UUID) ([]models.Match, error) {
	var matches []models.Match
	if err := r.db.Preload("Events").Where("league_id = ?", leagueId).Find(&matches).Error; err != nil {
		log.Printf("Error getting matches for league %v: %v", leagueId, err)
		return nil, err
	}
	return matches, nil
}

func (r *MatchRepository) DeleteMatch(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Match{}).Error; err != nil {
		log.Printf("Error deleting match with ID %v: %v", id, err)
		return err
	}
	return nil
}
