package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labbs/mynotion/pkg/config"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSession, downSession)
}

func upSession(ctx context.Context, tx *sql.Tx) error {
	// Get current dialect from goose
	var query string
	switch config.Database.Dialect {
	case "sqlite":
		query = `
		CREATE TABLE IF NOT EXISTS session (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			user_agent TEXT NOT NULL,
			ip_address TEXT NOT NULL,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_session_user_id ON session (user_id);
		`
	case "postgres":
		query = `
		CREATE TABLE IF NOT EXISTS session (
			id uuid PRIMARY KEY,
			user_id uuid NOT NULL,
			user_agent varchar NOT NULL,
			ip_address varchar NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_session_user_id ON session (user_id);
		`
	case "mysql":
		query = `
		CREATE TABLE IF NOT EXISTS session (
			id varchar(36) PRIMARY KEY,
			user_id varchar(36) NOT NULL,
			user_agent varchar(255) NOT NULL,
			ip_address varchar(255) NOT NULL,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_session_user_id ON session (user_id);
		`
	default:
		return fmt.Errorf("unsupported dialect: %s", config.Database.Dialect)
	}
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downSession(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS session")
	return err
}
