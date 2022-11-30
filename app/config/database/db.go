package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username  string
	Password  string
	ShortUrls []ShortUrls
}

type ShortUrls struct {
	gorm.Model
	shortUrl    string `gorm:"unique"`
	originalURL string `gorm:"not null"`
	UserID      uint
}

func DatabaseSetup() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{}, &ShortUrls{})
	if err != nil {
		fmt.Println(err)
	}
}
