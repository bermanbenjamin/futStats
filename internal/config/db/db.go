package db

import (
	"errors"
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbURL string) (err error) {

	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	err = DB.AutoMigrate(&models.Player{}, &models.Season{}, &models.Match{}, &models.Event{}, &models.League{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return errors.New("Error migrating models from database: " + err.Error())
	}

	return nil
}
