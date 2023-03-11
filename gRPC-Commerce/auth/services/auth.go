package services

import (
	"context"
	"net/http"

	"github.com/storyofhis/gRPC-Commerce/auth/models"
	"github.com/storyofhis/gRPC-Commerce/auth/pb"
	"github.com/storyofhis/gRPC-Commerce/auth/repositories"
	"github.com/storyofhis/gRPC-Commerce/auth/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthSvc struct {
	repo repositories.AuthRepositories
	Jwt  utils.JwtWrapper
}

func NewAuthSvc(repo repositories.AuthRepositories) *AuthSvc {
	return &AuthSvc{
		repo: repo,
	}
}

func (svc *AuthSvc) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := svc.repo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-MAIL already exist",
		}, nil
	}

	_, err = svc.repo.FindUserByUsername(ctx, req.Username)
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

	err = svc.repo.CreateUser(ctx, user)
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

func (svc *AuthSvc) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := svc.repo.FindUserByEmail(ctx, req.Email)
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
	err = bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Status Bad Request",
		}, nil
	}

	// response
	token, _ := svc.Jwt.GenerateToken(*user)
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (svc *AuthSvc) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := svc.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user := new(models.Users)

	_, err = svc.repo.FindUserByEmail(ctx, claims.Email)
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
