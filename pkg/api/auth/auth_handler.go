package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos/llmos-dashboard/pkg/utils"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SingUp struct {
	Name string `json:"name" binding:"required"`
	Login
}

type Handler struct {
	client *ent.Client
	ctx    context.Context
}

func NewAuthHandler(c *ent.Client, ctx context.Context) Handler {
	return Handler{
		client: c,
		ctx:    ctx,
	}
}

func (h *Handler) SignIn(c *gin.Context) {
	var l Login
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Info("login info", l.Email)

	user, err := h.GetUserByEmail(l.Email)
	if err != nil || !utils.CheckPasswordHash(l.Password, user.Password) {
		slog.Error("failed to get user", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user, either email or password is invalid"})
		return
	}

	expiredTime := time.Now().Add(7 * 24 * time.Hour)
	tokenStr, err := utils.GenerateToken(user.Name, expiredTime)
	if err != nil {
		slog.Error("failed to generate token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":             tokenStr,
		"token_type":        "Bearer",
		"id":                user.ID,
		"email":             user.Email,
		"name":              user.Name,
		"role":              user.Role,
		"profile_image_url": user.ProfileImageURL,
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	var s SingUp
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPw, err := utils.HashPassword(s.Password)
	if err != nil {
		slog.Error("failed to hash password", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := User{
		Name:         s.Name,
		Email:        s.Email,
		Password:     hashPw,
		Role:         "user",
		ProfileImage: "",
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}

	slog.Debug("signup info", s)
	_, err = h.CreateUser(user)
	if err != nil {
		slog.Error("failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func (h *Handler) GetSessionUser(c *gin.Context) {
	userObj, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty user"})
		return
	}

	user, ok := userObj.(User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user obj"})
		return
	}

	c.JSON(200, gin.H{
		"id":                user.ID,
		"email":             user.Email,
		"name":              user.Name,
		"role":              user.Role,
		"profile_image_url": user.ProfileImage,
	})
}

func (h *Handler) AuthMiddleware(c *gin.Context) {
	slog.Debug("path", c.Request.URL.Path)
	if strings.Contains(c.Request.URL.Path, "signin") || strings.Contains(c.Request.URL.Path, "signup") {
		return
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
	user, err := h.GetUserByUsername(claims.Username)
	if err != nil {
		slog.Error("failed to get user", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get user: %s", err.Error())})
		return
	}
	c.Set("user", user)
	c.Next()
}
