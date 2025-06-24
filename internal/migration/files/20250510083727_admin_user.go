package migrations

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/pressly/goose/v3"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	goose.AddMigrationContext(upAdminUser, downAdminUser)
}

const emailAdminUser string = "admin@zotion.local"
const nameAdminUser string = "admin"

func upAdminUser(ctx context.Context, tx *sql.Tx) error {
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
	INSERT INTO user (id, name, email, password, active, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, datetime('now'), datetime('now'));
	`, utils.UUIDv4(), nameAdminUser, emailAdminUser, bcryptHash, true)
	return err
}

func downAdminUser(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
	DELETE FROM user WHERE email = ?;
	`, emailAdminUser)
	return err
}
