package orders

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/auth"
	"github.com/storyofhis/gRPC-Commerce/config"
	"github.com/storyofhis/gRPC-Commerce/orders/controllers"
)

func RegisterRouters(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.NewAuthMiddlewareConfig(authSvc)

	svc := &ServiceClient{
		Client: NewServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	controllers.CreateOrder(ctx, svc.Client)
}
