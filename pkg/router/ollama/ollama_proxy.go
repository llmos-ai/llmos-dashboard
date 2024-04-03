package ollama

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/api/auth"
	"github.com/llmos/llmos-dashboard/pkg/generated/ent"
)

var defaultTarget = "http://localhost:11434"

func RegisterOllama(r *gin.Engine, client *ent.Client, ctx context.Context) error {
	auth := auth.NewAuthHandler(client, ctx)
	ollamaApi := r.Group("ollama")
	ollamaApi.Use(auth.AuthMiddleware)

	url, err := url.Parse(defaultTarget)
	if err != nil {
		return fmt.Errorf("failed to parse ollama target url: %s", err)
	}

	{
		ollamaApi.GET("/api/tags", ReverseProxy(url))
		ollamaApi.GET("/api/version", ReverseProxy(url))
		ollamaApi.POST("/api/chat", ReverseProxy(url))
		ollamaApi.GET("/cancel/:id", CancelRequest)
	}
	return nil
}

func CancelRequest(c *gin.Context) {
	c.JSON(200, gin.H{"status": true})
}

func ReverseProxy(url *url.URL) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func trimOllamaPrefix(path string) string {
	return strings.TrimPrefix(path, "/ollama")
}
