package rbac

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/labbs/mynotes/internal/tokenutil"
	"github.com/labbs/mynotes/pkg/models"
)

func (c *Config) Check() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		_logger := c.Logger.With().Str("request_id", fmt.Sprintf("%v", ctx.Locals("requestid"))).Logger()
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			_logger.Error().Str("event", "middleware.rbac_check_middleware.missing_authorization_header").Msg("Missing authorization header")
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing authorization header",
			})
		}

		t := strings.Split(authHeader, " ")
		userId, _, _ := tokenutil.GetSessionInformationFromToken(t[1])
		groups, err := c.UserService.GetGroupsByUserId(userId)
		if err != nil {
			_logger.Error().Err(err).Msg("Failed to get user groups")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get user groups",
			})
		}

		var isAdmin bool = false

		// check if one of the groups is an admin group
		for _, group := range groups {
			if group.Role == models.RoleAdmin {
				ctx.Context().SetUserValue("is_admin", true)
				return ctx.Next()
			}
		}

		// check if the path
		if strings.HasPrefix(ctx.Path(), "/api/v1/admin") && !isAdmin {
			_logger.Error().Str("event", "middleware.rbac_check_middleware.not_authorized").Msg("User is not authorized to access admin routes")
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "You are not authorized to access this resource",
			})
		}

		return ctx.Next()
	}
}
