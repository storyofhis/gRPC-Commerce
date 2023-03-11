package auth

import (
	"fmt"

	"github.com/storyofhis/gRPC-Commerce/auth/pb"
	"github.com/storyofhis/gRPC-Commerce/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func NewServiceClient(conf *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(conf.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not Connect: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
