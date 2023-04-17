package service

import (
	"context"
	"net/http"

	"github.com/storyofhis/product-service/models"
	"github.com/storyofhis/product-service/pb"
	"github.com/storyofhis/product-service/repositories"
)

type ProductServerSvc struct {
	Repo *repositories.ProductServerRepositories
	pb.ProductServiceServer
}

func NewServiceServer(repo *repositories.ProductServerRepositories) *ProductServerSvc {
	return &ProductServerSvc{
		Repo: repo,
	}
}

func (svc *ProductServerSvc) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := models.Product{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	err := svc.Repo.CreateProduct(ctx, &product)
	if err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, err
	}
	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
		Data: &pb.CreateProductRequest{
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		},
	}, nil
}

func (svc *ProductServerSvc) FindOne(ctx context.Context, req *pb.FindOneDataRequest) (*pb.FindOneDataResponse, error) {
	product := models.Product{}

	err := svc.Repo.FindOne(ctx, &product, req.Id)
	if err != nil {
		return &pb.FindOneDataResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id: product.Id,
		Data: &pb.CreateProductRequest{
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		},
	}

	return &pb.FindOneDataResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (svc *ProductServerSvc) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	err := svc.Repo.FindOne(ctx, &product, req.Id)
	if err != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock to low",
		}, nil
	}

	var log models.StockDecreaseLog
	err = svc.Repo.DecreaseStockById(ctx, log, req.OrderId)
	if err != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product = models.Product{
		Stock: product.Stock - 1,
	}
	svc.Repo.UpdateProductById(ctx, &product)

	log = models.StockDecreaseLog{
		OrderId:      req.OrderId,
		ProductRefer: product.Id,
	}

	svc.Repo.CreateLog(ctx, &log)
	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
