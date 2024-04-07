package modelfile

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/utils"
)

type ModelFileRequest struct {
	TagName   string    `json:"tagName" binding:"required"`
	UserID    string    `json:"userId" binding:"required"`
	Modelfile Modelfile `json:"modelfile" binding:"required"`
}

type ModelFileUpdate struct {
	Id        uuid.UUID `json:"id" binding:"required"`
	TagName   string    `json:"tagName" binding:"required"`
	Modelfile Modelfile `json:"modelfile" binding:"required"`
}

type Modelfile struct {
	Title             string    `json:"title" binding:"required"`
	TagName           string    `json:"tagName" binding:"required"`
	Content           string    `json:"content" binding:"required"`
	Desc              string    `json:"desc" binding:"required"`
	ImageURL          string    `json:"imageUrl,omitempty"`
	Categories        []string  `json:"categories" binding:"required"`
	SuggestionPrompts []Content `json:"suggestionPrompts"`
}

type Content struct {
	Content string `json:"content"`
}

type ModelfileResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	TagName   string    `json:"tagName"`
	Modelfile Modelfile `json:"modelfile"`
}

type FindByTagRequest struct {
	TagName string `json:"tagName"`
}

func (h *Handler) CreateModelFile(c *gin.Context) {
	user, err := utils.GetSessionUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var req ModelFileRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		slog.Error("failed to parse chat obj", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mf, err := json.Marshal(req.Modelfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	newMF, err := h.Create(user, req, string(mf))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSONP(http.StatusOK, newMF)
}

func (h *Handler) ListModelFile(c *gin.Context) {
	modelfiles, err := h.GetModelFiles()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	mfs := make([]ModelfileResponse, 0)

	for _, mf := range modelfiles {
		var modelfile Modelfile
		err = json.Unmarshal([]byte(mf.Modelfile), &modelfile)
		if err != nil {
			slog.Error("failed to parse modelfile obj", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		mfs = append(mfs, ModelfileResponse{
			ID:        mf.ID,
			UserID:    mf.UserId,
			TagName:   mf.TagName,
			Modelfile: modelfile,
			CreatedAt: mf.CreatedAt,
		})
	}

	c.JSONP(http.StatusOK, mfs)
}

func (h *Handler) DeleteModelFile(c *gin.Context) {
	id := c.Param("tagName")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.DeleteByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSONP(http.StatusOK, gin.H{"status": true})
}

func (h *Handler) GetModelFileByTagName(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	req := FindByTagRequest{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mfs, err := h.client.Modelfile.Query().
		Where(modelfile.TagName(req.TagName)).Only(h.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSONP(http.StatusOK, mfs)
}

func (h *Handler) UpdateModelFile(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var req ModelFileUpdate
	err = json.Unmarshal(data, &req)
	if err != nil {
		slog.Error("failed to parse chat obj", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mf, err := json.Marshal(req.Modelfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	updatedMF, err := h.Update(req, string(mf))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSONP(http.StatusOK, updatedMF)
}
