package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

type GroupController struct {
	GroupService models.GroupService
	Logger       zerolog.Logger
}

// GetGroups godoc
// @Summary Get all groups
// @Description Get all groups
// @Tags group
// @Accept json
// @Produce json
// @Success 200 {array} models.Group
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/group [get]
func (gc *GroupController) GetGroups(ctx *fiber.Ctx) error {
	logger := gc.Logger.With().Str("event", "api.groups.get").Logger()

	groups, err := gc.GroupService.GetAllGroups()
	if err != nil {
		logger.Error().Err(err).Msg("Error getting groups")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Msg("Groups retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(groups)
}

// GetGroupById godoc
// @Summary Get group by ID
// @Description Get group by ID
// @Tags group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} models.Group
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/group/{id} [get]
func (gc *GroupController) GetGroupById(ctx *fiber.Ctx) error {
	logger := gc.Logger.With().Str("event", "api.groups.getById").Logger()

	id := ctx.Params("id")
	group, err := gc.GroupService.GetGroupById(id)
	if err != nil {
		if err.Error() == "record not found" {
			logger.Warn().Str("id", id).Msg("Group not found")
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group not found"})
		}
		logger.Error().Err(err).Msg("Error getting group by ID")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("id", id).Msg("Group retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(group)
}
