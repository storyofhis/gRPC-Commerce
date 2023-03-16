package config

import (
	"log"

	"github.com/storyofhis/auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(url string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("Can not open connection")
		return nil, err
	}

	// automigrate
	db.Debug().AutoMigrate(&models.Users{})

	return db, nil
}
