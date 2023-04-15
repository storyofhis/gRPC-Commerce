package main

import (
	"fmt"
	"log"
	"net"

	"github.com/storyofhis/auth-service/common"
	"github.com/storyofhis/auth-service/config"
	"github.com/storyofhis/auth-service/pb"
	"github.com/storyofhis/auth-service/repositories"
	"github.com/storyofhis/auth-service/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	// err := godotenv.Load("./config/envs/dev.env")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db, err := config.ConnectionDB(c.DBUrl)
	if err != nil {
		log.Fatalln("Can not Connection to DB", err)
	}

	jwt := common.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "auth-service",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing: ", err)
	}

	fmt.Println("Auth Service on PORT: ", c.Port)

	repo := repositories.NewAuthServerRepo(db)
	svc := services.NewServiceServer(repo, jwt)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, svc)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
