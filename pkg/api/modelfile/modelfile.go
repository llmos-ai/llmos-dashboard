package modelfile

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
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
		return nil, err
	}
	slog.Debug("get model files: ", modelfiles)
	return modelfiles, nil
}

func (h *Handler) Create(user *entv1.User, req ModelFileRequest, mf string) (*entv1.Modelfile, error) {
	modelfile, err := h.client.Modelfile.
		Create().
		SetOwner(user).
		SetTagName(req.TagName).
		SetModelfile(mf).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	slog.Debug("modelfile created successfully", mf)
	return modelfile, nil
}

func (h *Handler) Update(update ModelFileUpdate, modelfile string) (*entv1.Modelfile, error) {
	mf, err := h.client.Modelfile.
		UpdateOneID(update.Id).
		SetTagName(update.TagName).
		SetModelfile(modelfile).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	slog.Debug("modelfile updated successfully", mf)
	return mf, nil
}

func (h *Handler) DeleteByUser(id uuid.UUID, userId uuid.UUID) error {
	return h.client.Modelfile.
		DeleteOneID(id).
		Where(modelfile.UserId(userId)).
		Exec(h.ctx)
}

func (h *Handler) DeleteByID(id uuid.UUID) error {
	return h.client.Modelfile.
		DeleteOneID(id).
		Exec(h.ctx)
}
