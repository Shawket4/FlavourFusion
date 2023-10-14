package Models

import (

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	
	if err := DB.AutoMigrate(&Category{}, &Item{}); err != nil {
		panic("Couldn't Migrate Models")
	}
	DB.Session(&gorm.Session{FullSaveAssociations: true})
}
