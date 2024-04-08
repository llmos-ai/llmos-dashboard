package auth

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

type User struct {
	ID           uint32 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	ProfileImage string `json:"profileImageUrl"`
	CreatedAt    string `json:"createdAt"`
}

func (h *Handler) CreateUser(user *entv1.User) (*entv1.User, error) {
	slog.Debug("Inserting user record ...", user.Email)
	createdUser, err := h.client.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRole(user.Role).
		SetProfileImageUrl(user.ProfileImageUrl).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	slog.Debug("User created successfully", createdUser)
	return createdUser, nil
}

func (h *Handler) GetUserByEmail(email string) (*entv1.User, error) {
	user, err := h.client.User.
		Query().
		Where(user.Email(email)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	slog.Debug("user returned: ", user)
	return user, nil
}

func (h *Handler) GetUserByID(id uuid.UUID) (*entv1.User, error) {
	slog.Debug("Fetching user record by username ...")
	user, err := h.client.User.Get(h.ctx, id)
	if err != nil {
		return nil, err
	}
	slog.Debug("user returned: ", user)
	return user, nil
}

func (h *Handler) GetAllUser() (entv1.Users, error) {
	users, err := h.client.User.Query().All(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}
	slog.Debug("users returned: ", users)
	return users, nil
}

func (h *Handler) DeleteUser(email string) error {
	slog.Debug("Deleting user record ...", email)
	id, err := h.client.User.Delete().
		Where(user.Email(email)).
		Exec(h.ctx)
	if err != nil {
		return err
	}
	slog.Debug("User deleted successfully", id)
	return nil
}

func (h *Handler) UpdateUserRoleByID(id uuid.UUID, role user.Role) (*entv1.User, error) {
	user, err := h.client.User.UpdateOneID(id).
		SetRole(role).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *Handler) UpdateUserByID(id uuid.UUID, u UpdateUser) (*entv1.User, error) {
	user, err := h.client.User.UpdateOneID(id).
		SetEmail(u.Email).
		SetName(u.Name).
		SetNillablePassword(u.Password).
		SetProfileImageUrl(u.ProfileImageUrl).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
