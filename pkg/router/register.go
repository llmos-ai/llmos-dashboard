package router

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/auth"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/localllm"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/openai"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/setting"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/webapi"
)

type RegisterRouter func(*gin.Engine, *ent.Client, context.Context) error

var registeredRouters = []RegisterRouter{
	auth.RegisterAuthRoute,
	openai.RegisterLiteLLM,
	webapi.RegisterWebApi,
	localllm.Register,
	setting.Register,
}

func RegisterRouters(r *gin.Engine, c *ent.Client, ctx context.Context) error {
	// enable CORS for all origins
	r.Use(CORSMiddleware())
	pprof.Register(r)

	// serve static files
	r.Use(static.Serve("/", static.LocalFile("ui/build", true)))
	// fallback to index.html
	r.NoRoute(defaultNoRouter)

	r.StaticFS("/static", gin.Dir("static", true))
	r.GET("/api/config", GetAPIConfig)
	r.GET("/api/changelog", GetChangelog)

	for _, router := range registeredRouters {
		err := router(r, c, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func defaultNoRouter(c *gin.Context) {
	if !strings.HasPrefix(c.Request.URL.Path, "/api") {
		c.File("./ui/build/index.html")
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "API not found"})
}
