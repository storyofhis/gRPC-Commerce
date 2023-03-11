package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/auth/controllers/params"
	"github.com/storyofhis/gRPC-Commerce/auth/pb"
)

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	var body params.Register

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Name:     body.Name,
		Username: body.Username,
		Age:      body.Age,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), res)
}
