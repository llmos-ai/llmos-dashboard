package modelfile

import (
	"context"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
)

type Handler struct {
	client *entv1.Client
	ctx    context.Context
}

func NewHandler(c *entv1.Client, ctx context.Context) Handler {
	return Handler{
		client: c,
		ctx:    ctx,
	}
}

func (h *Handler) GetModelFiles() (entv1.Modelfiles, error) {
	modelfiles, err := h.client.Modelfile.Query().All(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying modelfiles: %w", err)
	}
	slog.Debug("get model files: ", modelfiles)
	return modelfiles, nil
}
