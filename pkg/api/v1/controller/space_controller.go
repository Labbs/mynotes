package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

type SpaceController struct {
	SpaceService models.SpaceService
	Logger       zerolog.Logger
}

// GetSpaceById godoc
// @Summary Get space by id
// @Description Get space by id
// @Tags space
// @Accept json
// @Produce json
// @Param spaceId path string true "Space Id"
// @Success 200 {object} models.Space
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/space/{spaceId} [get]
func (sc *SpaceController) GetSpaceById(ctx *fiber.Ctx) error {
	logger := sc.Logger.With().Str("event", "api.spaces.get").Logger()

	spaceId := ctx.Params("spaceId")
	space, err := sc.SpaceService.GetSpaceById(spaceId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting space by id")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("space", spaceId).Msg("Space retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(space)
}

// CreateSpace godoc
// @Summary Create space
// @Description Create space
// @Tags space
// @Accept json
// @Produce json
// @Param space body models.CreateSpaceRequest true "Create space"
// @Success 201 {object} models.Space
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/space [post]
func (sc *SpaceController) CreateSpace(ctx *fiber.Ctx) error {
	logger := sc.Logger.With().Str("event", "api.spaces.create").Logger()

	var spaceRequest models.CreateSpaceRequest
	if err := ctx.BodyParser(&spaceRequest); err != nil {
		logger.Error().Err(err).Msg("Error parsing space")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userId := ctx.Locals("user_id").(string)

	newSpace := models.Space{
		Name: spaceRequest.Name,
		Members: models.Members{
			{
				Id:     userId,
				Type:   models.MemberTypeUser,
				Access: models.AccessTypeFull,
			},
		},
		Type: models.SpaceTypePublic,
	}

	space, err := sc.SpaceService.CreateSpace(newSpace)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating space")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("space", space.Id).Msg("Space created successfully")
	return ctx.Status(fiber.StatusCreated).JSON(space)
}
