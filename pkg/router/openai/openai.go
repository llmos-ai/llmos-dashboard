package openai

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/api/auth"
	"github.com/llmos/llmos-dashboard/pkg/generated/ent"
)

func RegisterLiteLLM(r *gin.Engine, client *ent.Client, ctx context.Context) error {
	auth := auth.NewAuthHandler(client, ctx)
	api := r.Group("/openai")
	api.Use(auth.AuthMiddleware)
	{
		api.GET("/api/models", getModelFiles)
	}

	return nil
}

func getModelFiles(c *gin.Context) {
	c.JSON(200, gin.H{"status": true, "auth": "true"})
}
