package repository

import (
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	DB *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{DB: db}
}

func (r *PlayerRepository) GetPlayerById(id uuid.UUID) (*models.Player, error) {
	var player models.Player
	if err := r.DB.First(&player, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Player with ID %d not found", id)
			return nil, nil // Return nil without error for not found cases
		}
		log.Printf("Error retrieving player with ID %d: %v", id, err)
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	if err := r.DB.Find(&players).Error; err != nil {
		log.Printf("Error retrieving all players: %v", err)
		return nil, err
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
