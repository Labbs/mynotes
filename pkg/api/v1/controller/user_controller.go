package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

type UserController struct {
	UserService models.UserService
	Logger      zerolog.Logger
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/user [get]
func (uc *UserController) GetUsers(ctx *fiber.Ctx) error {
	logger := uc.Logger.With().Str("event", "api.users.get").Logger()

	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		logger.Error().Err(err).Msg("Error getting users")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Msg("Users retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(users)
}
