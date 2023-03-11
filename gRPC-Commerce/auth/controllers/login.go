package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/auth/controllers/params"
	"github.com/storyofhis/gRPC-Commerce/auth/pb"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	var body params.Login

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), res)
}
