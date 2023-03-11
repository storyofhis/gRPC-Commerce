package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/auth"
	"github.com/storyofhis/gRPC-Commerce/config"
	"github.com/storyofhis/gRPC-Commerce/orders"
	"github.com/storyofhis/gRPC-Commerce/products"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	route := gin.Default()

	if err != nil {
		log.Println(err)
	}

	authSvc := auth.RegisterRouters(route, &c)
	orders.RegisterRouters(route, &c, authSvc)
	products.RegisterRoutes(route, &c, authSvc)

	route.Run(c.Port)
}
