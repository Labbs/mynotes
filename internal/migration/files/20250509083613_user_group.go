package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labbs/mynotes/pkg/config"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUserGroup, downUserGroup)
}

func upUserGroup(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	var query string
	switch config.Database.Dialect {
	case "sqlite":
		query = `
		CREATE TABLE IF NOT EXISTS user_group (
			user_id TEXT NOT NULL,
			group_id TEXT NOT NULL,
			created_at datetime NOT NULL,
			PRIMARY KEY (user_id, group_id)
		);
		CREATE INDEX IF NOT EXISTS idx_user_group_user_id ON user_group (user_id);
		CREATE INDEX IF NOT EXISTS idx_user_group_group_id ON user_group (group_id);
		`
	case "postgres":
		query = `
		CREATE TABLE IF NOT EXISTS user_group (
			user_id uuid NOT NULL,
			group_id uuid NOT NULL,
			created_at timestamp NOT NULL,
			PRIMARY KEY (user_id, group_id)
		);
		CREATE INDEX IF NOT EXISTS idx_user_group_user_id ON user_group (user_id);
		CREATE INDEX IF NOT EXISTS idx_user_group_group_id ON user_group (group_id);
		`
	case "mysql":
		return fmt.Errorf("mysql dialect is not supported yet")
	default:
		return fmt.Errorf("unsupported dialect: %s", config.Database.Dialect)
	}
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downUserGroup(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS user_group`)
	return err
}
