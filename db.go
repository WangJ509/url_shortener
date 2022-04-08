package main

import (
	"fmt"
	"os"
	"time"

	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/clause"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrRecordExpired  = errors.New("record expired")
)

var (
	db  *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("db_user"), os.Getenv("db_pass"), os.Getenv("db_url"), os.Getenv("db_name"))
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to open db connection")
	}
}

type ShortenURL struct {
	URL         string    `gorm:"column:url"`
	OriginalURL string    `gorm:"column:original_url"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ExpiredAt   time.Time `gorm:"column:expired_at"`
}

func CreateShortenURL(url string, originalURL string, expireTime time.Time) error {
	shortenURL := ShortenURL{
		URL:         url,
		OriginalURL: originalURL,
		ExpiredAt:   expireTime,
	}

	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "url"}},
		DoUpdates: clause.AssignmentColumns([]string{"expired_at"}),
	}).Create(&shortenURL)

	return result.Error
}

func GetOriginalURL(url string) (originalURL string, err error) {
	shortenURL := ShortenURL{}
	result := db.Where("url = ?", url).Where("").Take(&shortenURL)
	if result.Error == gorm.ErrRecordNotFound {
		return "", ErrRecordNotFound
	}
	if shortenURL.ExpiredAt.Before(time.Now()) {
		return "", ErrRecordExpired
	}

	return shortenURL.OriginalURL, result.Error
}
