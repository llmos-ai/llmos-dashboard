package setting

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/settings"
)

type settingRequest struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value"`
}

func (h *handler) GetAllSettings(c *gin.Context) {
	settings, err := h.client.Setting.Query().All(h.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, settings)
}

func (h *handler) UpdateSettingByName(c *gin.Context) {
	var setting settingRequest
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if setting.Name == settings.TokenExpireTimeSettingName {
		if err := validateSettingTokenExpireTime(setting.Value); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if setting.Name == settings.WebhookURLSettingName {
		if err := validateSettingWebhookURL(setting.Value); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	err := h.Set(setting.Name, setting.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	settings, err := h.client.Setting.Query().All(h.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, settings)
}

func validateSettingTokenExpireTime(value string) error {
	_, err := time.ParseDuration(value)
	if err != nil {
		return fmt.Errorf("invalid duration: %s", value)
	}
	return nil
}

func validateSettingWebhookURL(value string) error {
	// allow to reset empty value
	if value == "" {
		return nil
	}
	_, err := url.ParseRequestURI(value)
	if err != nil {
		return fmt.Errorf("invalid webhook url: %s", value)
	}
	return nil
}
