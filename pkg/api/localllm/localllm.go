package localllm

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/settings"
)

type handler struct {
	client *entv1.Client
	ctx    context.Context
}

type updateLocalLLMUrlRequest struct {
	URL string `json:"url"`
}

func NewHandler(client *entv1.Client, ctx context.Context) *handler {
	return &handler{client: client, ctx: ctx}
}

func (h *handler) GetLocalLLMUrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"url": settings.LocalLLMServerURL.Get(),
	})
}

func (h *handler) UpdateLocalLLMUrl(c *gin.Context) {
	var req updateLocalLLMUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := settings.LocalLLMServerURL.Set(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": req.URL})
}

func (h *handler) CancelRequest(c *gin.Context) {
	c.JSON(200, gin.H{"status": true})
}
