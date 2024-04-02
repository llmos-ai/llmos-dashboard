package auth

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

const schema string = `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  role TEXT,
  profile_image_url TEXT,
  created_at TEXT NOT NULL
  );`

type UserDB struct {
	db *sql.DB
}

type User struct {
	ID           uint32 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	ProfileImage string `json:"profile_image_url"`
	CreatedAt    string `json:"created_at"`
}

func NewUserDB(db *sql.DB) UserDB {
	return UserDB{
		db: db,
	}
}

func RegisterUserDB(db *sql.DB) error {
	_, err := db.Exec(schema) // Execute SQL Statements
	if err != nil {
		return err
	}
	slog.Info("user table created successfully")
	return nil
}

func (u *UserDB) CreateUser(user User) error {
	slog.Info("Inserting user record ...", user.Email)
	userSql := `INSERT INTO users(name, email, password, role, profile_image_url, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := u.db.Prepare(userSql) // Prepare statement.
	if err != nil {
		return err
	}
	result, err := statement.Exec(user.Name, user.Email, user.Password, user.Role, user.ProfileImage, user.CreatedAt)
	if err != nil {
		return err
	}
	slog.Info("User created successfully", result)
	return nil
}

func (u *UserDB) GetUserByEmail(email string) (User, error) {
	slog.Debug("Fetching user record ...")
	userSql := `SELECT * FROM users WHERE email = ?`
	row := u.db.QueryRow(userSql, email)
	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.ProfileImage, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	slog.Debug("User fetched successfully", user)
	return user, nil
}

func (u *UserDB) GetUserByUsername(username string) (User, error) {
	slog.Debug("Fetching user record by username ...")
	userSql := `SELECT * FROM users WHERE name = ?`
	row := u.db.QueryRow(userSql, username)
	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.ProfileImage, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	slog.Debug("User fetched successfully", user)
	return user, nil
}

func (u *UserDB) UpdateUser(user User) error {
	slog.Info("Updating user record ...")
	userSql := `UPDATE users SET name = ?, password = ?, role = ?, profile_image_url = ? WHERE email = ?`
	statement, err := u.db.Prepare(userSql) // Prepare statement.
	if err != nil {
		return err
	}
	result, err := statement.Exec(user.Name, user.Password, user.Role, user.ProfileImage, user.Email)
	if err != nil {
		return err
	}
	slog.Info("User updated successfully", result)
	return nil
}

func (u *UserDB) DeleteUser(email string) error {
	slog.Info("Deleting user record ...")
	userSql := `DELETE FROM users WHERE email = ?`
	statement, err := u.db.Prepare(userSql) // Prepare statement.
	if err != nil {
		return err
	}
	user, err := statement.Exec(email)
	if err != nil {
		return err
	}
	slog.Info("User deleted successfully", user)
	return nil
}
