package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/router/auth"
)

type RegisterRouter func(*gin.Engine, *sql.DB)

var registeredRouters = []RegisterRouter{
	auth.RegisterAuthRoute,
}

func RegisterRouters(r *gin.Engine, db *sql.DB) {

	// enable CORS for all origins
	r.Use(CORSMiddleware())

	// Serve frontend static files
	r.StaticFS("static", gin.Dir("static", true))

	r.GET("api/config", GetAPIConfig)

	for _, router := range registeredRouters {
		router(r, db)
	}
}
