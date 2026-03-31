package repository

import (
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *SeasonRepository {
	return &SeasonRepository{db: db}
}

func (r *SeasonRepository) CreateSeason(season *models.Season, leagueId uuid.UUID) (*models.Season, error) {
	if err := r.db.Create(season).Error; err != nil {
		log.Printf("Error creating season: %v", err)
		return nil, err
	}

	// Associate season with league via join table
	var league models.League
	if err := r.db.First(&league, leagueId).Error; err != nil {
		log.Printf("Error finding league %v: %v", leagueId, err)
		return nil, err
	}
	if err := r.db.Model(&league).Association("Seasons").Append(season); err != nil {
		log.Printf("Error associating season with league: %v", err)
		return nil, err
	}

	return season, nil
}

func (r *SeasonRepository) GetSeasonsByLeagueId(leagueId uuid.UUID) ([]models.Season, error) {
	var league models.League
	if err := r.db.Preload("Seasons").First(&league, leagueId).Error; err != nil {
		log.Printf("Error getting seasons for league %v: %v", leagueId, err)
		return nil, err
	}
	return league.Seasons, nil
}

func (r *SeasonRepository) GetSeasonById(id uuid.UUID) (*models.Season, error) {
	var season models.Season
	if err := r.db.First(&season, id).Error; err != nil {
		log.Printf("Error getting season %v: %v", id, err)
		return nil, err
	}
	return &season, nil
}
