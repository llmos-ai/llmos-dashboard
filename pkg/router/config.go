package router

import (
	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/config"
	"github.com/llmos/llmos-dashboard/pkg/constant"
	"github.com/llmos/llmos-dashboard/pkg/version"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func GetAPIConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":                     true,
		"name":                       constant.AppName,
		"version":                    version.GetFriendlyVersion(),
		"images":                     false,
		"default_models":             nil,
		"default_prompt_suggestions": config.GetDefaultPromptSuggestions(),
	})
}
