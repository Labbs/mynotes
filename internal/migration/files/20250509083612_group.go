package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labbs/zotion/pkg/config"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upGroup, downGroup)
}

func upGroup(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	var query string
	switch config.Database.Dialect {
	case "sqlite":
		query = `
		CREATE TABLE IF NOT EXISTS "group" (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			role TEXT,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL
		);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_group_name ON "group" (name);
		`
	case "postgres":
		query = `
		CREATE TABLE IF NOT EXISTS "group" (
			id uuid PRIMARY KEY,
			name varchar NOT NULL,
			description varchar,
			role varchar,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL
		);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_group_name ON "group" (name);
		`
	case "mysql":
		return fmt.Errorf("mysql dialect is not supported yet")
	default:
		return fmt.Errorf("unsupported dialect: %s", config.Database.Dialect)
	}
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downGroup(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS "group"`)
	return err
}
