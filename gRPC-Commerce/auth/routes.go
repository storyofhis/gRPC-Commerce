package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/gRPC-Commerce/auth/controllers"
	"github.com/storyofhis/gRPC-Commerce/config"
)

func RegisterRouters(r *gin.Engine, conf *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: NewServiceClient(conf),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	controllers.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	controllers.Login(ctx, svc.Client)
}
