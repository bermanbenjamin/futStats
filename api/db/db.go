package db

import (
	"log"
	"os"

	"github.com/bermanbenjamin/futStats/api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Ensure environment variables are set.")
	}
	dsn := os.Getenv("DB_STRING_PATH")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.Season{})
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(models.League{})
	db.AutoMigrate(&models.Match{})

	DB = db
}
