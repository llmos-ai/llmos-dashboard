package setting

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/api/auth"
	"github.com/llmos-ai/llmos-dashboard/pkg/api/setting"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/settings"
)

func Register(r *gin.Engine, client *ent.Client, ctx context.Context) error {
	auth := auth.NewAuthHandler(client, ctx)
	api := r.Group("/api/v1")
	api.Use(auth.AuthMiddleware, auth.AdminMiddleware)

	handler := setting.NewHandler(client, ctx)
	{
		api.GET("/settings/", handler.GetAllSettings)
		api.POST("/settings/", handler.UpdateSettingByName)
	}

	return settings.SetProvider(&handler)
}
