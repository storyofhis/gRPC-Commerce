package repositories

import (
	"context"
	"time"

	"github.com/storyofhis/product-service/models"
	"gorm.io/gorm"
)

type ProductServerRepositories struct {
	DB *gorm.DB
}

func NewProductServerRepo(db *gorm.DB) *ProductServerRepositories {
	return &ProductServerRepositories{DB: db}
}

func (repo *ProductServerRepositories) CreateProduct(ctx context.Context, product *models.Product) error {
	product.CreatedAt = time.Now()
	if err := repo.DB.WithContext(ctx).Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductServerRepositories) FindOne(ctx context.Context, product *models.Product, id int64) error {
	err := repo.DB.WithContext(ctx).First(&product, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductServerRepositories) DecreaseStockById(ctx context.Context, stock models.StockDecreaseLog, id int64) error {
	err := repo.DB.WithContext(ctx).Where(
		&models.StockDecreaseLog{
			OrderId: id,
		},
	).First(&stock).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductServerRepositories) UpdateProductById(ctx context.Context, product *models.Product) {
	_ = repo.DB.WithContext(ctx).Save(&product)
}

func (repo *ProductServerRepositories) CreateLog(ctx context.Context, log *models.StockDecreaseLog) {
	_ = repo.DB.WithContext(ctx).Create(&log)
}
