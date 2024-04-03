package auth

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/api/auth"
	"github.com/llmos/llmos-dashboard/pkg/generated/ent"
)

func RegisterAuthRoute(r *gin.Engine, c *ent.Client, ctx context.Context) error {

	auth := auth.NewAuthHandler(c, ctx)
	apiv1 := r.Group("/api/v1/auths")
	apiv1.Use(auth.AuthMiddleware)
	{
		apiv1.GET("/", auth.GetAuthorizedUser)
		apiv1.POST("/signin", auth.SignIn)
		apiv1.POST("/signup", auth.SignUp)
	}
	return nil
}
