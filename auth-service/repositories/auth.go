package repositories

import (
	"context"
	"strings"
	"time"

	"github.com/storyofhis/auth-service/models"
	"gorm.io/gorm"
)

type AuthServerRepo struct {
	DB *gorm.DB
}

func NewAuthServerRepo(db *gorm.DB) *AuthServerRepo {
	return &AuthServerRepo{
		DB: db,
	}
}

func (repo *AuthServerRepo) CreateUser(ctx context.Context, user *models.Users) error {
	user.CreatedAt = time.Now()
	if err := repo.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AuthServerRepo) FindUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where("LOWER(email) = ?", strings.ToLower(email)).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthServerRepo) FindUserByUsername(ctx context.Context, username string) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where("LOWER(username) = ?", strings.ToLower(username)).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthServerRepo) FindUserByUId(ctx context.Context, id int64) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where("id = ?", id).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
