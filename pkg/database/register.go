package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"

	"github.com/llmos/llmos-dashboard/pkg/database/auth"
)

const dbFileName = "llmos-ui.db"

type registerDB func(*sql.DB) error

var registerDBs = []registerDB{
	auth.RegisterUserDB,
}

func RegisterSQLiteDB() (*sql.DB, error) {
	var err error
	sql, err := sql.Open("sqlite3", dbFileName)
	if err != nil || sql.Ping() != nil {
		return sql, err
	}

	for _, register := range registerDBs {
		err = register(sql)
		if err != nil {
			slog.Error("Failed to register auth", err)
			return nil, err
		}
	}

	return sql, nil
}
