package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/chat"
	"github.com/llmos-ai/llmos-dashboard/pkg/utils"
)

type Handler struct {
	client *entv1.Client
	ctx    context.Context
}

type Chat struct {
	Chat map[string]interface{} "json:chat"
}

func NewHandler(c *entv1.Client, ctx context.Context) Handler {
	return Handler{
		client: c,
		ctx:    ctx,
	}
}

func (h *Handler) ListAll() (entv1.Chats, error) {
	chats, err := h.client.Chat.Query().All(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying chats: %w", err)
	}
	slog.Debug("get model files: ", chats)
	return chats, nil
}

func (h *Handler) ListByUser(user *entv1.User) ([]*entv1.Chat, error) {
	chats, err := user.QueryChats().All(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying chats: %w", err)
	}
	slog.Debug("get model files: ", chats)
	return chats, nil
}

func (h *Handler) Create(chatStr string, user *entv1.User) (*entv1.Chat, error) {
	chat, err := h.client.Chat.
		Create().
		SetTitle("New Chat").
		SetChat(chatStr).
		SetOwner(user).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (h *Handler) Update(chatStr string, id uuid.UUID) (*entv1.Chat, error) {
	chat, err := h.client.Chat.
		UpdateOneID(id).
		SetChat(chatStr).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	slog.Debug("updated chat: ", chat)
	return chat, nil
}

func (h *Handler) ListUserChats(c *gin.Context) {
	user, err := utils.GetSessionUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "error": err.Error()})
		return
	}
	chats, err := h.ListByUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, chats)
}

func (h *Handler) UpdateChatByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat id not found"})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "invalid chat id"})
		return
	}

	jsonChat, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat not found"})
		return
	}

	var chatObj Chat
	err = json.Unmarshal(jsonChat, &chatObj)
	if err != nil {
		slog.Error("failed to parse chat obj", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	chatStr, err := json.Marshal(chatObj.Chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	savedChat, err := h.Update(string(chatStr), uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"id":         savedChat.ID,
		"title":      savedChat.Title,
		"chat":       chatObj.Chat,
		"user_id":    savedChat.UserID,
		"created_at": savedChat.CreatedAt,
	})
}

func (h *Handler) CreateChat(c *gin.Context) {
	// get session user
	user, err := utils.GetSessionUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "error": err.Error()})
		return
	}

	jsonChat, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat not found"})
		return
	}

	var chatObj Chat
	err = json.Unmarshal(jsonChat, &chatObj)
	if err != nil {
		slog.Error("failed to parse chat obj", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	chatStr, err := json.Marshal(chatObj.Chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	chat, err := h.Create(string(chatStr), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, chat)
}

func (h *Handler) GetChatByID(c *gin.Context) {
	// get session user
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "chat id not found"})
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

	var chatObj interface{}
	if err = json.Unmarshal([]byte(chat.Chat), &chatObj); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"id":         chat.ID,
		"title":      chat.Title,
		"chat":       chatObj,
		"user_id":    chat.UserID,
		"created_at": chat.CreatedAt,
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

	h.client.Chat.Delete().
		Where(chat.ID(uuid.MustParse(id))).
		Exec(h.ctx)

	c.JSONP(http.StatusOK, gin.H{"status": true})
}
