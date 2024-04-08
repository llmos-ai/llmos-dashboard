package auth

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	entv1User "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
	"github.com/llmos-ai/llmos-dashboard/pkg/settings"
	v1 "github.com/llmos-ai/llmos-dashboard/pkg/types/v1"

	"github.com/llmos-ai/llmos-dashboard/pkg/constant"
	"github.com/llmos-ai/llmos-dashboard/pkg/utils"
)

const tokenType = "Bearer"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SingUp struct {
	Name string `json:"name" binding:"required"`
	LoginRequest
}

type UpdateRole struct {
	ID   uuid.UUID      `json:"id" binding:"required"`
	Role entv1User.Role `json:"role" binding:"required"`
}

type UpdateUser struct {
	Email           string  `json:"email" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	Password        *string `json:"password"`
	ProfileImageUrl string  `json:"profileImageUrl"`
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
	var l LoginRequest
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

	tokenStr, err := utils.GenerateToken(user.ID)
	if err != nil {
		slog.Error("failed to generate token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": constant.MessageErrorGenerateToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":           tokenStr,
		"tokenType":       tokenType,
		"id":              user.ID,
		"email":           user.Email,
		"name":            user.Name,
		"role":            user.Role,
		"profileImageUrl": user.ProfileImageUrl,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": constant.MessageErrorHashPassword})
		return
	}

	role := v1.GetUserRole(settings.DefaultUserRole.Get())
	// always set the first user as the admin user
	users, err := h.GetAllUser()
	if err != nil {
		slog.Error("failed to list users", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		role = entv1User.RoleAdmin
	}

	user := &entv1.User{
		Name:            s.Name,
		Email:           s.Email,
		Password:        hashPw,
		Role:            role,
		ProfileImageUrl: "/user.png",
	}

	slog.Debug("signup info", s)
	user, err = h.CreateUser(user)
	if err != nil {
		slog.Error("failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenStr, err := utils.GenerateToken(user.ID)
	if err != nil {
		slog.Error("failed to generate token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": constant.MessageErrorGenerateToken})
		return
	}

	c.Set("user", user)
	c.JSON(http.StatusOK, gin.H{
		"token":           tokenStr,
		"tokenType":       tokenType,
		"id":              user.ID,
		"email":           user.Email,
		"name":            user.Name,
		"role":            user.Role,
		"profileImageUrl": user.ProfileImageUrl,
	})
}

func (h *Handler) GetAuthorizedUser(c *gin.Context) {
	user, err := utils.GetSessionUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":              user.ID,
		"email":           user.Email,
		"name":            user.Name,
		"role":            user.Role,
		"profileImageUrl": user.ProfileImageUrl,
	})
}

func (h *Handler) ListAllUser(c *gin.Context) {
	users, err := h.GetAllUser()
	if err != nil {
		slog.Error("failed to list users", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, users)
}

func (h *Handler) UpdateUserRole(c *gin.Context) {
	update := UpdateRole{}
	// using BindJson method to serialize body with struct
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UpdateUserRoleByID(update.ID, update.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	uuid, err := uuid.Parse(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	update := UpdateUser{}
	// using BindJson method to serialize body with struct
	if err = c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if update.Password != nil {
		passwd, err := utils.HashPassword(*update.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		update.Password = &passwd
	}

	user, err := h.UpdateUserByID(uuid, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
