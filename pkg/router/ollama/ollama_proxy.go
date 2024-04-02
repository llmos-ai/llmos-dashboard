package ollama

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/api/auth"
	"github.com/llmos/llmos-dashboard/pkg/generated/ent"
)

var defaultTarget = "http://localhost:11434"

func RegisterOllama(r *gin.Engine, client *ent.Client, ctx context.Context) {
	auth := auth.NewAuthHandler(client, ctx)
	ollamaApi := r.Group("ollama")
	ollamaApi.Use(auth.AuthMiddleware)
	{
		ollamaApi.GET("/api/tags", ReverseProxy(defaultTarget))
	}
}

func ReverseProxy(target string) gin.HandlerFunc {
	url, err := url.Parse(target)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(500, gin.H{"failed to parse target": err.Error()})
		}
	}
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			//r := c.Request
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			req.URL.Path = trimOllamaPrefix(req.URL.Path)
			//req.Header["Authorization"] = []string{r.Header.Get("Authorization")}
			//fmt.Printf("director req: %+v", req)
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func trimOllamaPrefix(path string) string {
	return strings.TrimPrefix(path, "/ollama")
}
