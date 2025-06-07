package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/mynotes/pkg/models"
	"github.com/rs/zerolog"
)

type AdminController struct {
	UserService     models.UserService
	GroupService    models.GroupService
	SpaceService    models.SpaceService
	DocumentService models.DocumentService
	Logger          zerolog.Logger
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/admin/users [get]
func (ac *AdminController) GetUsers(ctx *fiber.Ctx) error {
	logger := ac.Logger.With().Str("event", "api.admin.get_users").Logger()

	users, err := ac.UserService.GetAllUsers()
	if err != nil {
		logger.Error().Err(err).Msg("Error getting users")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Int("count", len(users)).Msg("Users retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(users)
}

// GetGroups godoc
// @Summary Get all groups
// @Description Get all groups
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {array} models.Group
// @Failure 500 {object} models.ErrorResponse
func (ac *AdminController) GetGroups(ctx *fiber.Ctx) error {
	logger := ac.Logger.With().Str("event", "api.admin.get_groups").Logger()

	groups, err := ac.GroupService.GetAllGroups()
	if err != nil {
		logger.Error().Err(err).Msg("Error getting groups")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Int("count", len(groups)).Msg("Groups retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(groups)
}

// GetSpaces godoc
// @Summary Get all spaces
// @Description Get all spaces
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {array} models.Space
// @Failure 500 {object} models.ErrorResponse
func (ac *AdminController) GetSpaces(ctx *fiber.Ctx) error {
	logger := ac.Logger.With().Str("event", "api.admin.get_spaces").Logger()

	spaces, err := ac.SpaceService.GetAllSpaces()
	if err != nil {
		logger.Error().Err(err).Msg("Error getting spaces")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Int("count", len(spaces)).Msg("Spaces retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(spaces)
}

// GetDocuments godoc
