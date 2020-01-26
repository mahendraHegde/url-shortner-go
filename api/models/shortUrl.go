package models

import "github.com/jinzhu/gorm"

type ShortUrl struct {
	Model
	Url    string `gorm:"unique;not null" json:"url" binding:"required"`
	Suffix string `json:"suffix,omitempty"`
	Short  string `gorm:"unique;not null" json:"short"`
}

func (shortUrl *ShortUrl) Insert(db *gorm.DB) error {
	return db.Create(&shortUrl).Error
}
func (shortUrl *ShortUrl) GetLastUrlSeq(db *gorm.DB) (ShortUrl, error) {
	res := ShortUrl{}
	err := db.Last(&res).Where(&ShortUrl{Suffix: ""}).Error
	return res, err
}
func (shortUrl *ShortUrl) GetByUrl(db *gorm.DB, url string) (ShortUrl, error) {
	res := ShortUrl{}
	err := db.Where(&ShortUrl{Url: url}).First(&res).Error
	return res, err
}
