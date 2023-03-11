package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func NewAuthMiddlewareConfig(svc *ServiceClient) *AuthMiddlewareConfig {
	return &AuthMiddlewareConfig{
		svc: svc,
	}
}

func (middleware *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := middleware.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}
