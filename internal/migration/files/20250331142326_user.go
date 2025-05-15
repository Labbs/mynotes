package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labbs/mynotion/pkg/config"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUser, downUser)
}

func upUser(ctx context.Context, tx *sql.Tx) error {
	// Get current dialect from goose
	var query string
	switch config.Database.Dialect {
	case "sqlite":
		query = `
		CREATE TABLE IF NOT EXISTS user (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			avatar_url TEXT,
			preferences JSONB,
			active bool NOT NULL,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL
		);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_user_name ON user (name);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON user (email);
		CREATE INDEX IF NOT EXISTS idx_user_active ON user (active);
		`
	case "postgres":
		query = `
		CREATE TABLE IF NOT EXISTS 'user' (
			id uuid PRIMARY KEY,
			name varchar NOT NULL,
			email varchar NOT NULL,
			password varchar NOT NULL,
			avatar_url varchar,
			preferences jsonb,
			active bool NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL
		);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_user_name ON user (name);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON user (email);
		CREATE INDEX IF NOT EXISTS idx_user_active ON user (active);
		`
	case "mysql":
		return fmt.Errorf("mysql dialect is not supported yet")
	default:
		return fmt.Errorf("unsupported dialect: %s", config.Database.Dialect)
	}
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downUser(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS user;")
	return err
}
