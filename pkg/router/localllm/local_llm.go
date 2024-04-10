package localllm

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/api/auth"
	"github.com/llmos-ai/llmos-dashboard/pkg/api/localllm"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/settings"
)

func Register(r *gin.Engine, client *ent.Client, ctx context.Context) error {
	auth := auth.NewAuthHandler(client, ctx)
	api := r.Group("/localllm")
	api.Use(auth.AuthMiddleware)
	h := localllm.NewHandler(client, ctx)
	{
		// reverse proxy for ollama apis
		api.GET("/ollama/api/version", ReverseProxy)                         // Get ollama version
		api.GET("/ollama/api/tags", ReverseProxy)                            // List Local Models
		api.POST("/ollama/api/generate", ReverseProxy)                       // Generate a completion
		api.POST("/ollama/api/chat", ReverseProxy)                           // Generate a chat completion
		api.POST("/ollama/api/create", auth.AdminMiddleware, ReverseProxy)   // Create a Model
		api.POST("/ollama/api/pull", auth.AdminMiddleware, ReverseProxy)     // Pull a Model
		api.DELETE("/ollama/api/delete", auth.AdminMiddleware, ReverseProxy) // Delete a Model

		// custom localllm api
		api.GET("/url", h.GetLocalLLMUrl)
		api.POST("/url/update", auth.AdminMiddleware, h.UpdateLocalLLMUrl)
		api.GET("/cancel/:id", h.CancelRequest)
	}

	return nil
}

func ReverseProxy(c *gin.Context) {
	localLLMUrl := settings.LocalLLMServerURL.Get()
	url, err := url.Parse(localLLMUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = url.Host
		req.Header.Set("Origin", "")
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = trimOllamaPrefix(req.URL.Path)
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func trimOllamaPrefix(path string) string {
	return strings.TrimPrefix(path, "/localllm/ollama")
}
