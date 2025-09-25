package db

import (
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dbURL string) (err error) {
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: false,
	})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	// Enable UUID extension for PostgreSQL
	err = DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Printf("Warning: Failed to create UUID extension (might already exist): %v", err)
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
