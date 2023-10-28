package Models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string  `json:"name" yaml:"name"`
	CategoryID  uint    `json:"category_id" yaml:"category_id"`
	Description string  `json:"description" yaml:"description"`
	Price       float64 `json:"price" yaml:"price"`
	ImageURL    string  `json:"image_url"`
}
