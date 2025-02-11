package db

import (
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

	// Drop existing tables to reset relationships
	err = DB.Migrator().DropTable(&models.League{}, &models.LeagueMember{}, &models.Player{}, &models.Season{}, &models.Match{}, &models.Event{})
	if err != nil {
		log.Printf("Failed to drop tables: %v", err)
	}

	// Migrate in correct order - Player first, then League
	err = DB.AutoMigrate(
		&models.Player{}, // Player must be migrated first since League depends on it
		&models.League{},
		&models.LeagueMember{}, // Add the join table model
		&models.Season{},
		&models.Match{},
		&models.Event{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return err
	}

	return nil
}
