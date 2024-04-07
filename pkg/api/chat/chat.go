package chat

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
)

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
	slog.Debug("get user model files: ", user.Name, chats)
	return chats, nil
}

func (h *Handler) Create(user *entv1.User, req NewChatRequest) (*entv1.Chat, error) {
	chat, err := h.client.Chat.
		Create().
		SetTitle(req.Title).
		SetHistory(req.History).
		SetMessages(req.Messages).
		SetModels(req.Models).
		SetTags(req.Tags).
		SetOwner(user).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (h *Handler) Update(id uuid.UUID, req UpdateChatRequest) (*entv1.Chat, error) {
	client := h.client.Chat.UpdateOneID(id).
		SetNillableHistory(req.History).
		SetNillableTitle(req.Title)

	if req.Messages != nil || len(req.Messages) > 0 {
		client.SetMessages(req.Messages)
	}

	chat, err := client.Save(h.ctx)
	if err != nil {
		return nil, err
	}
	slog.Debug("updated chat: ", chat)
	return chat, nil
}
