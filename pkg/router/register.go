package router

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/auth"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/litellm"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/ollama"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/openai"
	"github.com/llmos-ai/llmos-dashboard/pkg/router/webapi"
)

type RegisterRouter func(*gin.Engine, *ent.Client, context.Context) error

var registeredRouters = []RegisterRouter{
	auth.RegisterAuthRoute,
	litellm.RegisterLiteLLM,
	openai.RegisterLiteLLM,
	webapi.RegisterWebApi,
	ollama.RegisterOllama,
}

func RegisterRouters(r *gin.Engine, c *ent.Client, ctx context.Context) {

	// enable CORS for all origins
	r.Use(CORSMiddleware())

	// Serve frontend static files
	r.StaticFS("static", gin.Dir("static", true))

	r.GET("api/config", GetAPIConfig)

	for _, router := range registeredRouters {
		router(r, c, ctx)
	}
}
