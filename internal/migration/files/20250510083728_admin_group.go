package migrations

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAdminGroup, downAdminGroup)
}

const groupAdminGroup string = "admin"

func upAdminGroup(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
	INSERT INTO "group" (id, name, description, role, created_at, updated_at)
	VALUES (?, ?, ?, ?, datetime('now'), datetime('now'));
	`, utils.UUIDv4(), groupAdminGroup, "admin group", "admin")
	return err
}

func downAdminGroup(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
	DELETE FROM "group" WHERE name = ?;
	`, groupAdminGroup)
	return err
}
