package common

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/storyofhis/auth-service/models"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

var (
	ErrTokenInvalid  = errors.New("token invalid")
	ErrTokenInactive = errors.New("token inactive")
)

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64  `json:"id"`
	Email string `json:"email"`
	// Role  string `json:"role"`
}

func (w *JwtWrapper) GenerateToken(user models.Users) (string, error) {
	claims := &jwtClaims{
		Id:    int64(user.Id),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}
