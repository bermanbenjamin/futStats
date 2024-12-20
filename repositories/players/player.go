package repositories

import (
	"log"

	"github.com/bermanbenjamin/futStats/db"
	models "github.com/bermanbenjamin/futStats/models/players"
)

func GetPlayers(id uint64) (*models.Player, error) {
	var player models.Player
	if err := db.DB.Find(&player, id).Error; err != nil {
		log.Println("Failed to find player", err)
		return nil, err
	}
	return &player, nil
}

func GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	if err := db.DB.Find(&players).Error; err != nil {
		log.Println("Failed to get all players", err)
		return nil, err
	}
	return players, nil
}

func AddPlayer(player models.Player) error {
	if err := db.DB.Create(&player).Error; err != nil {
		log.Println("Failed to add player", err)
		return err
	}
	return nil
}

func UpdatePlayer(player models.Player) (models.Player, error) {
	if err := db.DB.Save(&player).Error; err != nil {
		log.Println("Failed to update player", err)
		return player, err
	}

	return player, nil
}

func DeletePlayer(id uint64) error {
	var player models.Player
	if err := db.DB.Where("id = ?", id).Delete(&player).Error; err != nil {
		log.Println("Failed to delete player", err)
		return err
	}
	return nil
}
