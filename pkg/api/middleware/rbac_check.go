package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/labbs/mynotes/pkg/models"
	"github.com/rs/zerolog"
)

func RBACCheckMiddleware(logger zerolog.Logger, userService models.UserService, groupService models.GroupService, spaceService models.SpaceService, documentService models.DocumentService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(string)
		groups, err := userService.GetGroupsByUserId(userId)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to get user groups")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get user groups",
			})
		}

		// check if one of the groups is an admin group
		for _, group := range groups {
			if group.Role == models.RoleAdmin {
				// User is an admin, allow access
				return c.Next()
			}
		}

		// check if the path

		fmt.Println("Checking RBAC for path:", c.Path())

		return c.Next()
	}
}
