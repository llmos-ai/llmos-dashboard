package auth

import (
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"

	entv1 "github.com/llmos/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos/llmos-dashboard/pkg/generated/ent/user"
)

type User struct {
	ID           uint32 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	ProfileImage string `json:"profile_image_url"`
	CreatedAt    string `json:"created_at"`
}

func (h *Handler) CreateUser(user User) (*entv1.User, error) {
	slog.Info("Inserting user record ...", user.Email)
	createdUser, err := h.client.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRole(user.Role).
		SetProfileImageURL(user.ProfileImage).
		Save(h.ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("User created successfully", h)
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

func (h *Handler) GetUserByUsername(username string) (*entv1.User, error) {
	slog.Debug("Fetching user record by username ...")
	user, err := h.client.User.
		Query().
		Where(user.Name(username)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	slog.Debug("user returned: ", h)
	return user, nil
}

func (h *Handler) DeleteUser(email string) error {
	slog.Info("Deleting user record ...")
	id, err := h.client.User.Delete().
		Where(user.Email(email)).
		Exec(h.ctx)
	if err != nil {
		return err
	}
	slog.Info("User deleted successfully", id)
	return nil
}
