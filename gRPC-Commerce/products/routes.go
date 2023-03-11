package products

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/auth"
	"github.com/storyofhis/gRPC-Commerce/config"
	"github.com/storyofhis/gRPC-Commerce/products/controllers"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.NewAuthMiddlewareConfig(authSvc)

	svc := &ServiceClient{
		Client: NewServiceClient(c),
	}

	routes := r.Group("/products")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.POST("/:id", svc.FindOne)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	controllers.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	controllers.FindOne(ctx, svc.Client)
}
