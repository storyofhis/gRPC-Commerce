package repositories

import (
	"context"
	"time"

	"github.com/storyofhis/orders-service/models"
	"gorm.io/gorm"
)

type OrderServerRepositories struct {
	DB *gorm.DB
}

func NewOrderServerRepositories(db *gorm.DB) *OrderServerRepositories {
	return &OrderServerRepositories{
		DB: db,
	}
}

func (repo *OrderServerRepositories) CreateOrder(ctx context.Context, order *models.Order) error {
	order.CreatedAt = time.Now()
	err := repo.DB.WithContext(ctx).Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}
