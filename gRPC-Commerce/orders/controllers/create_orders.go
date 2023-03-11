package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/orders/controllers/params"
	"github.com/storyofhis/gRPC-Commerce/orders/pb"
)

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	var body params.CreateOrder

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
