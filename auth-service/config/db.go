package config

import (
	"fmt"
	"log"
	"os"

	"github.com/storyofhis/auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() (*gorm.DB, error) {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASS")
	name := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Can not open connection")
		return nil, err
	}

	// automigrate
	db.Debug().AutoMigrate(&models.Users{})

	return db, nil
}
