package models

import "gorm.io/gorm"

type ShortUrls struct {
	gorm.Model
	shortUrl    string `gorm:"unique"`
	originalURL string `gorm:"not null"`
	UserID      uint
}
