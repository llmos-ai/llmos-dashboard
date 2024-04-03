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

func (h *Handler) CreateUser(user entv1.User) (*entv1.User, error) {
	slog.Debug("Inserting user record ...", user.Email)
	createdUser, err := h.client.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRole(user.Role).
		SetProfileImageURL(user.ProfileImageURL).
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

func (h *Handler) ListUsers() ([]*entv1.User, error) {
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
