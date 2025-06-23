package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labbs/zotion/pkg/config"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upDocument, downDocument)
}

func upDocument(ctx context.Context, tx *sql.Tx) error {
	var query string
	switch config.Database.Dialect {
	case "sqlite":
		query = `
		CREATE TABLE IF NOT EXISTS document (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			slug TEXT NOT NULL,
			type TEXT NOT NULL,
			config JSONB,
			metadata JSONB,
			parent_id TEXT,
			properties JSONB,
			space_id TEXT,
			content TEXT,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			deleted_at datetime
		);
		CREATE INDEX IF NOT EXISTS idx_document_name ON document (name);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_document_slug ON document (slug);
		CREATE INDEX IF NOT EXISTS idx_document_deleted_at ON document (deleted_at);
		CREATE INDEX IF NOT EXISTS idx_document_space_id ON document (space_id);
		`
	case "postgres":
		query = `
		CREATE TABLE IF NOT EXISTS document (
			id uuid PRIMARY KEY,
			name varchar NOT NULL,
			slug varchar NOT NULL,
			type varchar NOT NULL,
			config jsonb,
			metadata jsonb,
			parent_id varchar,
			properties jsonb,
			space_id varchar,
			content text,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			deleted_at timestamp
		);
		CREATE INDEX IF NOT EXISTS idx_document_name ON document (name);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_document_slug ON document (slug);
		CREATE INDEX IF NOT EXISTS idx_document_deleted_at ON document (deleted_at);
		CREATE INDEX IF NOT EXISTS idx_document_space_id ON document (space_id);
		`
	case "mysql":
		return fmt.Errorf("mysql dialect is not supported yet")
	default:
		return fmt.Errorf("unsupported dialect: %s", config.Database.Dialect)
	}
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downDocument(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS document;")
	return err
}
