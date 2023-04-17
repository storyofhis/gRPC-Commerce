package service

import (
	"context"
	"log"
	"net/http"

	"github.com/storyofhis/orders-service/models"
	"github.com/storyofhis/orders-service/pb"
	"github.com/storyofhis/orders-service/repositories"
	"github.com/storyofhis/orders-service/service/client"
)

type OrderServerSvc struct {
	Repo   repositories.OrderServerRepositories
	Client client.ProductServiceClient
	pb.OrderServiceServer
}

func NewProductServerSvc(repo repositories.OrderServerRepositories, client client.ProductServiceClient) *OrderServerSvc {
	return &OrderServerSvc{Repo: repo, Client: client}
}

func (svc *OrderServerSvc) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := svc.Client.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	} else if product.Data.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "Stock too less",
		}, nil
	}

	order := models.Order{
		Price:     product.Data.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,
	}

	err = svc.Repo.CreateOrder(ctx, &order)
	if err != nil {
		log.Fatalln("Cannot Create Order, ", err)
	}

	res, err := svc.Client.DecreaseStock(req.ProductId, order.Id)
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if res.Status == http.StatusConflict {
		svc.Repo.DB.Delete(&models.Order{}, order.Id)
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
		// CreatedAt: order.CreatedAt,
	}, nil
}
