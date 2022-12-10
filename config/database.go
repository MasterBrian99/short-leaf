package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseSetup() {
	database, err := gorm.Open(sqlite.Open(os.Getenv("DB_FILE")), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Failed")
	}

	DB = database
}

func GetDB() *gorm.DB {
	return DB
}
