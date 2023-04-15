package repositories

import (
	"context"
	"time"

	"github.com/storyofhis/auth-service/models"
	"github.com/storyofhis/auth-service/pb"
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

func (repo *AuthServerRepo) CreateUser(ctx context.Context, user models.Users) error {
	user.CreatedAt = time.Now()
	if err := repo.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AuthServerRepo) FindUserByEmail(ctx context.Context, req *pb.RegisterRequest) (models.Users, error) {
	// user := new(models.Users)
	var user models.Users
	err := repo.DB.WithContext(ctx).Where(
		&models.Users{
			Email: req.Email,
		},
	).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *AuthServerRepo) FindUserByUsername(ctx context.Context, req *pb.RegisterRequest) (*models.Users, error) {
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

func (repo *AuthServerRepo) FindUserByUId(ctx context.Context, id int64) (*models.Users, error) {
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

func (repo *AuthServerRepo) FindUserByEmailLogin(ctx context.Context, req *pb.LoginRequest) (models.Users, error) {
	var user models.Users
	err := repo.DB.WithContext(ctx).Where(
		&models.Users{
			Email: req.Email,
		},
	).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
