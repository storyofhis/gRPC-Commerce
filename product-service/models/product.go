package models

import (
	"time"
)

type Product struct {
	Id               int64            `json:"id" gorm:"primaryKey"`
	Name             string           `json:"name"`
	Stock            int64            `json:"stock"`
	Price            int64            `json:"price"`
	StockDecreaseLog StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
	DeletedAt        time.Time        `json:"deleted_at"`
}
