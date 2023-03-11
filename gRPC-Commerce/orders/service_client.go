package orders

import (
	"fmt"

	"github.com/storyofhis/gRPC-Commerce/config"
	"github.com/storyofhis/gRPC-Commerce/orders/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func NewServiceClient(conf *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(conf.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Couldn't connect: ", err)
	}
	return pb.NewOrderServiceClient(cc)
}
