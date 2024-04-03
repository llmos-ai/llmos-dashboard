package modelfile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListModelFile(c *gin.Context) {
	modelfiles, err := h.GetModelFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}
	//jsonResult := json.Marshal(modelfiles)
	//c.JSON(200, gin.H{jsonResult})
	c.JSONP(http.StatusOK, modelfiles)
}
