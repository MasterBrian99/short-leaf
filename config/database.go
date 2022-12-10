package config

import (
	"log"
	"os"

	"github.com/MasterBrian99/short-leaf/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseSetup() {
	DB, err := gorm.Open(sqlite.Open(os.Getenv("DB_FILE")), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Failed")
	}
	err = DB.AutoMigrate(&models.ShortUrls{}, &models.User{})

	if err != nil {
		log.Fatal("Migrate Failed")
	}
}
