package Models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string `json:"name"`
	CategoryID uint `json:"category_id"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImageURL string `json:"image_url"`
}

