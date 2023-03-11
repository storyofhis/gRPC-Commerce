package products

import (
	"fmt"

	"github.com/storyofhis/gRPC-Commerce/config"
	"github.com/storyofhis/gRPC-Commerce/products/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductsServiceClient
}

func NewServiceClient(c *config.Config) pb.ProductsServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Couldn't connect: ", err)
	}

	return pb.NewProductsServiceClient(cc)
}
