package repository

import (
	"fmt"
	"log"

	model "github.com/bermanbenjamin/futStats/internal/models"
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
func (r *PlayerRepository) GetPlayerBy(filter constants.QueryFilter, value string) (*model.Player, error) {
	var player model.Player
	query := fmt.Sprintf("%s = ?", filter)
	if err := r.DB.Where(query, value).First(&player).Error; err != nil {
		log.Printf("Error retrieving player with filter %s='%s': %v", filter, value, err)
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) GetAllPlayersBy(filterQuery constants.QueryFilter, filterValue string) ([]*model.Player, error) {
	var players []*model.Player
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

func (r *PlayerRepository) AddPlayer(player model.Player) (*model.Player, error) {
	if err := r.DB.Create(&player).Error; err != nil {
		log.Printf("Error adding player: %v", err)
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) UpdatePlayer(player model.Player) (*model.Player, error) {
	if err := r.DB.Save(&player).Error; err != nil {
		log.Printf("Error updating player with ID %s: %v", player.ID, err)
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) DeletePlayer(id uuid.UUID) error {
	if err := r.DB.Where("id = ?", id.String()).Delete(&model.Player{}).Error; err != nil {
		log.Printf("Error deleting player with ID %d: %v", id, err)
		return err
	}
	return nil
}
