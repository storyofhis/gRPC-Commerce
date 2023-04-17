package config

import (
	"log"

	"github.com/storyofhis/product-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// automigrate
	db.Debug().AutoMigrate(models.Product{}, models.StockDecreaseLog{})
	return db, err
}
