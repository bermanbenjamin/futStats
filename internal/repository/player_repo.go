package repository

import (
	"fmt"
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/transport/http/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	DB *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{DB: db}
}
func (r *PlayerRepository) GetPlayerBy(filter constants.QueryFilter, value string) (*models.Player, error) {
	var player models.Player
	query := fmt.Sprintf("%s = ?", filter)
	if err := r.DB.Where(query, value).First(&player).Error; err != nil {
		log.Printf("Error retrieving player with filter %s='%s': %v", filter, value, err)
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) GetAllPlayersBy(filterQuery constants.QueryFilter, filterValue string) ([]*models.Player, error) {
	var players []*models.Player
	if filterQuery != "" && filterValue != "" {
		if err := r.DB.Where(filterQuery+" LIKE?", "%"+filterValue+"%").Find(&players).Error; err != nil {
			log.Printf("Error retrieving players with filter %s='%s': %v", filterQuery, filterValue, err)
			return nil, err
		}
	} else {
		if err := r.DB.Find(&players).Error; err != nil {
			log.Printf("Error retrieving all players: %v", err)
			return nil, err
		}
	}
	return players, nil
}

func (r *PlayerRepository) AddPlayer(player models.Player) (*models.Player, error) {
	if err := r.DB.Create(&player).Error; err != nil {
		log.Printf("Error adding player: %v", err)
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) UpdatePlayer(player models.Player) (*models.Player, error) {
	if err := r.DB.Save(&player).Error; err != nil {
		log.Printf("Error updating player with ID %s: %v", player.ID, err)
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) DeletePlayer(id uuid.UUID) error {
	if err := r.DB.Where("id = ?", id.String()).Delete(&models.Player{}).Error; err != nil {
		log.Printf("Error deleting player with ID %d: %v", id, err)
		return err
	}
	return nil
}
