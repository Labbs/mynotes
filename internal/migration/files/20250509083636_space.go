package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labbs/mynotion/pkg/config"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSpace, downSpace)
}

func upSpace(ctx context.Context, tx *sql.Tx) error {
	var query string
	switch config.Database.Dialect {
	case "sqlite":
		query = `
		CREATE TABLE IF NOT EXISTS space (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			slug TEXT NOT NULL,
			icon TEXT,
			icon_color TEXT,
			description TEXT,
			type TEXT,
			members TEXT,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			deleted_at datetime
		);
		CREATE INDEX IF NOT EXISTS idx_space_name ON space (name);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_space_slug ON space (slug);
		CREATE INDEX IF NOT EXISTS idx_space_deleted_at ON space (deleted_at);
		`
	case "postgres":
		query = `
		CREATE TABLE IF NOT EXISTS space (
			id uuid PRIMARY KEY,
			name varchar NOT NULL,
			slug varchar NOT NULL,
			icon varchar,
			icon_color varchar,
			description varchar,
			type varchar,
			members jsonb,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			deleted_at timestamp
		);
		CREATE INDEX IF NOT EXISTS idx_space_name ON space (name);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_space_slug ON space (slug);
		CREATE INDEX IF NOT EXISTS idx_space_deleted_at ON space (deleted_at);
		`
	case "mysql":
		return fmt.Errorf("mysql dialect is not supported yet")
	default:
		return fmt.Errorf("unsupported dialect: %s", config.Database.Dialect)
	}
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downSpace(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS space")
	return err
}
