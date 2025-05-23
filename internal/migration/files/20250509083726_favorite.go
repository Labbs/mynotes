package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labbs/mynotes/pkg/config"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upFavorite, downFavorite)
}

func upFavorite(ctx context.Context, tx *sql.Tx) error {
	var query string
	switch config.Database.Dialect {
	case "sqlite":
		query = `
		CREATE TABLE IF NOT EXISTS favorite (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			document_id TEXT NOT NULL,
			database_id TEXT NOT NULL,
			position TEXT NOT NULL,
			created_at datetime NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_favorite_user ON favorite (user_id);
		`
	case "postgres":
		query = `
		CREATE TABLE IF NOT EXISTS favorite (
			id uuid PRIMARY KEY,
			user_id uuid NOT NULL,
			document_id uuid NOT NULL,
			database_id uuid NOT NULL,
			position varchar NOT NULL,
			created_at timestamp NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_favorite_user ON favorite (user_id);
		`
	case "mysql":
		return fmt.Errorf("mysql dialect is not supported yet")
	default:
		return fmt.Errorf("unsupported dialect: %s", config.Database.Dialect)
	}
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downFavorite(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS favorite`)
	return err
}
