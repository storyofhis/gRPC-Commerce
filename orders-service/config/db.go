package config

import (
	"log"

	"github.com/storyofhis/orders-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// automigrate
	db.Debug().AutoMigrate(models.Order{})
	return db, err
}
