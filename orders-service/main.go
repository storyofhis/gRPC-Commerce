package main

import (
	"fmt"
	"log"
	"net"

	"github.com/storyofhis/orders-service/config"
	"github.com/storyofhis/orders-service/pb"
	"github.com/storyofhis/orders-service/repositories"
	"github.com/storyofhis/orders-service/service"
	"github.com/storyofhis/orders-service/service/client"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db, err := config.Init(c.DBUrl)
	if err != nil {
		log.Fatalln("Cannot Initializa DB: ", err)
	}

	list, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Cannot listing: ", err)
	}

	client := client.InitProductServiceClient(c.ProductSvcUrl)
	repo := repositories.NewOrderServerRepositories(db)
	svc := service.NewProductServerSvc(*repo, client)

	if err != nil {
		log.Fatalln("Cannot listing Client: ", err)
	}

	fmt.Println("Order Svc On Port", c.Port)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, svc)

	if err = grpcServer.Serve(list); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
