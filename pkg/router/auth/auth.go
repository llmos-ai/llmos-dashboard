package auth

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/api/auth"
)

func RegisterAuthRoute(r *gin.Engine, db *sql.DB) {

	auth := auth.NewAuthHandler(db)
	apiv1 := r.Group("api/v1/auths")
	apiv1.Use(auth.AuthMiddleware)
	{
		apiv1.GET("/", auth.GetSessionUser)
		apiv1.POST("/signin", auth.SignIn)
		apiv1.POST("/signup", auth.SignUp)
	}
}
