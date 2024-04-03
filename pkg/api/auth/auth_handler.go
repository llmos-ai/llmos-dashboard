package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/constant"
	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	entv1User "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
	"github.com/llmos-ai/llmos-dashboard/pkg/utils"
)

const tokenType = "Bearer"

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SingUp struct {
	Name string `json:"name" binding:"required"`
	Login
}

type Handler struct {
	client *entv1.Client
	ctx    context.Context
}

func NewAuthHandler(c *entv1.Client, ctx context.Context) Handler {
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

	slog.Debug("login info", l.Email)

	user, err := h.GetUserByEmail(l.Email)
	if err != nil || !utils.CheckPasswordHash(l.Password, user.Password) {
		slog.Error("failed to get user", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": constant.MessageErrorLogin})
		return
	}

	expiredTime := time.Now().Add(7 * 24 * time.Hour)
	tokenStr, err := utils.GenerateToken(user.Name, expiredTime)
	if err != nil {
		slog.Error("failed to generate token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": constant.MessageErrorGenerateToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":             tokenStr,
		"token_type":        tokenType,
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

	var role = entv1User.RolePending
	users, err := h.ListUsers()
	if err != nil {
		slog.Error("failed to list users", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	if len(users) == 0 {
		role = entv1User.RoleAdmin
	}

	user := entv1.User{
		Name:            s.Name,
		Email:           s.Email,
		Password:        hashPw,
		Role:            role,
		ProfileImageURL: "/user.png",
	}

	slog.Debug("signup info", s)
	_, err = h.CreateUser(user)
	if err != nil {
		slog.Error("failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expiredTime := time.Now().Add(7 * 24 * time.Hour)
	tokenStr, err := utils.GenerateToken(user.Name, expiredTime)
	if err != nil {
		slog.Error("failed to generate token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": constant.MessageErrorGenerateToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":             tokenStr,
		"token_type":        tokenType,
		"id":                user.ID,
		"email":             user.Email,
		"name":              user.Name,
		"role":              user.Role,
		"profile_image_url": user.ProfileImageURL,
	})
}

func (h *Handler) GetAuthorizedUser(c *gin.Context) {
	user, err := utils.GetSessionUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user obj"})
		return
	}

	c.JSON(200, gin.H{
		"id":                user.ID,
		"email":             user.Email,
		"name":              user.Name,
		"role":              user.Role,
		"profile_image_url": user.ProfileImageURL,
	})
}

func (h *Handler) AuthMiddleware(c *gin.Context) {
	slog.Debug("path", c.Request.URL.Path)
	if strings.Contains(c.Request.URL.Path, "signin") || strings.Contains(c.Request.URL.Path, "signup") {
		slog.Debug("skipping auth middleware")
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
