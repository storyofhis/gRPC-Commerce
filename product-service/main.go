package main

import (
	"fmt"
	"log"
	"net"

	"github.com/storyofhis/product-service/config"
	"github.com/storyofhis/product-service/pb"
	"github.com/storyofhis/product-service/repositories"
	"github.com/storyofhis/product-service/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db, err := config.Init(c.DBUrl)
	if err != nil {
		log.Fatalln("Cannot Connection to DB: ", err)
	}

	list, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to Listing: ", err)
	}

	fmt.Println("Product svc on PORT: ", c.Port)

	repo := repositories.NewProductServerRepo(db)
	svc := service.NewServiceServer(repo)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, svc)

	if err = grpcServer.Serve(list); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
