package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint `gorm:"primary_key" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
}

func Migrate(db *gorm.DB) {

	db.AutoMigrate(&ShortUrl{})
}
