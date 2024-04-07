package chat

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/chat"
	v1 "github.com/llmos-ai/llmos-dashboard/pkg/types/v1"
	"github.com/llmos-ai/llmos-dashboard/pkg/utils"
)

type Handler struct {
	client *entv1.Client
	ctx    context.Context
}

type NewChatRequest struct {
	History  v1.Histroy   `json:"history" binding:"required"`
	Messages []v1.Message `json:"messages" binding:"required"`
	Models   []string     `json:"models" binding:"required"`
	Title    string       `json:"title" binding:"required"`
	Tags     []string     `json:"tags"`
}

type UpdateChatRequest struct {
	Title    *string      `json:"title,omitempty"`
	History  *v1.Histroy  `json:"history" binding:"required"`
	Messages []v1.Message `json:"messages" binding:"required"`
}

func NewHandler(c *entv1.Client, ctx context.Context) Handler {
	return Handler{
		client: c,
		ctx:    ctx,
	}
}

func (h *Handler) GetUserChats(c *gin.Context) {
	user, err := utils.GetSessionUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "error": err.Error()})
		return
	}
	chats, err := h.ListByUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, chats)
}
func (h *Handler) CreateChat(c *gin.Context) {
	// get session user
	user, err := utils.GetSessionUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "error": err.Error()})
		return
	}

	chatData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat not found"})
		return
	}

	var req NewChatRequest
	err = json.Unmarshal(chatData, &req)
	if err != nil {
		slog.Error("failed to parse chat obj", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	chat, err := h.Create(user, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, chat)
}

func (h *Handler) UpdateChatByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat id is empty"})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "invalid chat id"})
		return
	}

	chatData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var req UpdateChatRequest
	err = json.Unmarshal(chatData, &req)
	if err != nil {
		slog.Error("failed to parse chat obj", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	savedChat, err := h.Update(uuid, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, savedChat)
}

func (h *Handler) GetChatByID(c *gin.Context) {
	// get session user
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat id empty"})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "invalid chat id"})
		return
	}

	// get chat by id
	chat, err := h.client.Chat.Query().
		Where(chat.ID(uuid)).
		Only(h.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chat":      chat,
		"id":        chat.ID,
		"titile":    chat.Title,
		"userId":    chat.UserId,
		"createdAt": chat.CreatedAt,
	})

}

func (h *Handler) GetChatTagsByID(c *gin.Context) {
	c.JSONP(http.StatusOK, []string{})
}

func (h *Handler) DeleteChatByID(c *gin.Context) {
	// get chat id
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat id not found"})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.client.Chat.Delete().Where(chat.ID(uid)).Exec(h.ctx); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSONP(http.StatusOK, gin.H{"status": true})
}
