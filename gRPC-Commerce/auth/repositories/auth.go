package repositories

import (
	"context"
	"strings"
	"time"

	"github.com/storyofhis/gRPC-Commerce/auth/models"
	"gorm.io/gorm"
)

type AuthRepositories struct {
	DB *gorm.DB
}

func NewAuthRepositories(db *gorm.DB) *AuthRepositories {
	return &AuthRepositories{
		DB: db,
	}
}

func (repo *AuthRepositories) CreateUser(ctx context.Context, user *models.Users) error {
	user.CreatedAt = time.Now()
	if err := repo.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AuthRepositories) FindUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where("LOWER(email) = ?", strings.ToLower(email)).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthRepositories) FindUserByUsername(ctx context.Context, username string) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where("LOWER(username) = ?", strings.ToLower(username)).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthRepositories) FindUserByUId(ctx context.Context, id int64) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where("id = ?", id).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
