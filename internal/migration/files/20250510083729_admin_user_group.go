package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAdminUserGroup, downAdminUserGroup)
}

func upAdminUserGroup(ctx context.Context, tx *sql.Tx) error {
	var userId, groupId string

	err := tx.QueryRowContext(ctx, `
	SELECT id FROM user WHERE email = ?;
	`, emailAdminUser).Scan(&userId)
	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, `
	SELECT id FROM "group" WHERE name = ?;
	`, groupAdminGroup).Scan(&groupId)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
	INSERT INTO user_group (user_id, group_id, created_at)
	VALUES (?, ?, datetime('now'));
	`, userId, groupId)

	return err
}

func downAdminUserGroup(ctx context.Context, tx *sql.Tx) error {
	var userId, groupId string

	err := tx.QueryRowContext(ctx, `
	SELECT id FROM user WHERE email = ?;
	`, emailAdminUser).Scan(&userId)
	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, `
	SELECT id FROM "group" WHERE name = ?;
	`, groupAdminGroup).Scan(&groupId)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
	DELETE FROM user_group WHERE user_id = ? AND group_id = ?;
	`, userId, groupId)
	return err
}
