package services

import (
	"context"
	"net/http"

	"github.com/storyofhis/auth-service/common"
	"github.com/storyofhis/auth-service/models"
	"github.com/storyofhis/auth-service/pb"
	"github.com/storyofhis/auth-service/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServerSvc struct {
	Repo *repositories.AuthServerRepo
	Jwt  common.JwtWrapper
}

func NewServiceServer(repo *repositories.AuthServerRepo, Jwt common.JwtWrapper) *AuthServerSvc {
	return &AuthServerSvc{
		Repo: repo,
		Jwt:  Jwt,
	}
}

func (svc *AuthServerSvc) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := svc.Repo.FindUserByEmail(ctx, req)
	// fmt.Println("teesty", user)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-MAIL already exist",
		}, nil
	}
	_, err = svc.Repo.FindUserByUsername(ctx, req)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Username already exist",
		}, nil
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  "Internal Server Error",
		}, nil
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	// request from user
	user.Email = req.Email
	user.Name = req.Name
	user.Username = req.Username
	user.Age = req.Age
	user.Password = string(hashedPass)

	err = svc.Repo.CreateUser(ctx, user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  "Internal Server Error",
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Id:     int64(user.Id),
		Data: &pb.RegisterRequest{
			Email:    req.Username,
			Name:     req.Name,
			Username: req.Username,
			Age:      req.Age,
			Password: string(hashedPass),
		},
	}, nil
}

func (svc *AuthServerSvc) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := svc.Repo.FindUserByEmailLogin(ctx, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.LoginResponse{
				Status: http.StatusBadRequest,
				Error:  "Invalid Credentials",
			}, err
		}
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Invalid Server Error",
		}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Status Bad Request" + err.Error(),
		}, nil
	}

	// response
	token, _ := svc.Jwt.GenerateToken(user)
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (svc *AuthServerSvc) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := svc.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user := new(models.Users)

	err = svc.Repo.DB.Where(
		&models.Users{
			Email: claims.Email,
		},
	).First(user).Error
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: int64(user.Id),
	}, nil
}
