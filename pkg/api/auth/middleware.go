package auth

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	entuser "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
	"github.com/llmos-ai/llmos-dashboard/pkg/settings"
	"github.com/llmos-ai/llmos-dashboard/pkg/utils"
)

var authWhiteListPaths = []string{
	"/api/v1/auths/signin",
	"/api/v1/auths/signup",
}

func (h *Handler) AuthMiddleware(c *gin.Context) {
	slog.Debug("path", c.Request.URL.Path)
	for _, path := range authWhiteListPaths {
		if strings.TrimSuffix(strings.ToLower(c.Request.URL.Path), "/") == path {
			slog.Debug("skipping auth middleware", c.Request.URL.Path)
			c.Next()
			return
		}
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}

	tokenString = tokenString[len("Bearer "):]
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// if auth exist, find user by token
	user, err := h.GetUserByID(claims.UUID)
	if err != nil {
		slog.Error("failed to get user", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get user: %s", err.Error())})
		return
	}
	c.Set("user", user)
	c.Next()
}

func (h *Handler) AdminMiddleware(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	u, ok := user.(*entv1.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}
	if u.Role != entuser.RoleAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	c.Next()
}

func (h *Handler) SingUpMiddleware(c *gin.Context) {
	if settings.Signup.Get() != "true" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Signup is disabled, please contact admin to enable it."})
		return
	}
}
