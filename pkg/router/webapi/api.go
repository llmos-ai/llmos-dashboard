package webapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/api/auth"
	"github.com/llmos-ai/llmos-dashboard/pkg/api/chat"
	"github.com/llmos-ai/llmos-dashboard/pkg/api/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
)

func RegisterWebApi(r *gin.Engine, client *ent.Client, ctx context.Context) error {
	auth := auth.NewAuthHandler(client, ctx)
	api := r.Group("/api/v1")
	api.Use(auth.AuthMiddleware)

	modelHandler := modelfile.NewHandler(client, ctx)
	chatHandler := chat.NewHandler(client, ctx)
	{
		api.GET("/documents/", ListDocuments)
		api.GET("/prompts/", ListPrompts)

		// Chat API
		api.GET("/chats/tags/all", ListChatTags)
		api.GET("/chats/", chatHandler.GetUserChats)
		api.POST("/chats/new", chatHandler.CreateChat)
		api.GET("/chats/:id", chatHandler.GetChatByID)
		api.POST("/chats/:id", chatHandler.UpdateChatByID)
		api.DELETE("/chats/:id", chatHandler.DeleteChatByID)
		api.GET("/chats/:id/tags", chatHandler.GetChatTagsByID)

		// User API
		api.GET("/users/", auth.ListAllUser)
		api.POST("/users/:id/update", auth.UpdateUser)
		api.POST("/users/update/role", auth.UpdateUserRole)

		// Modefile API
		api.GET("/modelfiles/", modelHandler.ListModelFile)
		api.POST("/modelfiles/", modelHandler.GetModelFileByTagName)
		api.POST("/modelfiles/create", modelHandler.CreateModelFile)
		api.POST("/modelfiles/update", modelHandler.UpdateModelFile)
		api.DELETE("/modelfiles/:tagName", modelHandler.DeleteModelFile)
	}
	return nil
}

func ListDocuments(c *gin.Context) {
	c.JSONP(http.StatusOK, []string{})
}

func ListPrompts(c *gin.Context) {
	c.JSONP(http.StatusOK, []string{})
}

func ListChatTags(c *gin.Context) {
	c.JSONP(http.StatusOK, []string{})
}
