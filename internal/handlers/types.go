package handlers

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	OriginalURL string
	ShortCode   string
}

type playLoad struct {
	Url string `json:"url"`
}
