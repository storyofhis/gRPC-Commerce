package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/products/controllers/params"
	"github.com/storyofhis/gRPC-Commerce/products/pb"
)

func CreateProduct(ctx *gin.Context, c pb.ProductsServiceClient) {
	var body params.CreateProduct

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductsRequest{
		Name:        body.Name,
		Description: body.Description,
		Sku:         body.Sku,
		Stock:       body.Stock,
		Price:       body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
