package router

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/auth"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/ollama"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/openai"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/setting"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/webapi"
)

type RegisterRouter func(*gin.Engine, *ent.Client, context.Context) error

var registeredRouters = []RegisterRouter{
	auth.RegisterAuthRoute,
	openai.RegisterLiteLLM,
	webapi.RegisterWebApi,
	ollama.RegisterOllama,
	setting.Register,
}

func RegisterRouters(r *gin.Engine, c *ent.Client, ctx context.Context) error {

	// enable CORS for all origins
	r.Use(CORSMiddleware())

	// Serve frontend static files
	r.StaticFS("static", gin.Dir("static", true))

	r.GET("api/config", GetAPIConfig)
	r.GET("api/changelog", GetChangelog)

	for _, router := range registeredRouters {
		err := router(r, c, ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
