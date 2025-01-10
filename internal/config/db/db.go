package db

import (
	"errors"
	"log"

	model "github.com/bermanbenjamin/futStats/internal/models"
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

	err = DB.AutoMigrate(&model.Player{}, &model.Season{}, &model.Match{}, &model.Event{}, &model.League{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return errors.New("Error migrating model.from database: " + err.Error())
	}

	return nil
}
