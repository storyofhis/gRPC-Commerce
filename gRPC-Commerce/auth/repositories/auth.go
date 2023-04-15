package repositories

import (
	"context"
	"time"

	"github.com/storyofhis/gRPC-Commerce/auth/models"
	"github.com/storyofhis/gRPC-Commerce/auth/pb"
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

func (repo *AuthRepositories) FindUserByEmail(ctx context.Context, req *pb.RegisterRequest) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where(
		&models.Users{
			Email: req.Email,
		},
	).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthRepositories) FindUserByUsername(ctx context.Context, req *pb.RegisterRequest) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where(
		&models.Users{
			Username: req.Username,
		},
	).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthRepositories) FindUserByUId(ctx context.Context, id int64) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where(
		&models.Users{
			Id: uint(id),
		},
	).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthRepositories) FindUserByEmailLogin(ctx context.Context, req *pb.LoginRequest) (*models.Users, error) {
	user := new(models.Users)
	err := repo.DB.WithContext(ctx).Where(
		&models.Users{
			Email: req.Email,
		},
	).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
