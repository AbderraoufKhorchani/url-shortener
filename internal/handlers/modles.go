package handlers

import (
	"log"
	"time"

	"github.com/AbderraoufKhorchani/url-shortener/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func New(dbPool *gorm.DB) error {
	db = dbPool
	return db.AutoMigrate(&URL{})
}

func ConnectToDB(dsn string) (*gorm.DB, error) {
	var counts int64
	for {
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection, err
		}

		if counts > 10 {
			return nil, err
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}

func saveURL(originalURL string) (*URL, error) {
	shortcode, err := utils.GenerateShortCode(6) // Specify the desired length of the shortcode
	if err != nil {
		print(1)
		return nil, err

	}

	newURL := &URL{
		OriginalURL: originalURL,
		ShortCode:   shortcode,
	}

	result := db.Create(newURL)
	if result.Error != nil {
		print(2)
		return nil, result.Error
	}

	return newURL, nil
}

func getURL(shortcode string) (*URL, error) {

	var url URL
	result := db.Where("short_code = ?", shortcode).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}
