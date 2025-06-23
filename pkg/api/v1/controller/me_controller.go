package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

type MeController struct {
	SpaceService    models.SpaceService
	UserService     models.UserService
	FavoriteService models.FavoriteService
	Logger          zerolog.Logger
}

// GetMySpaces godoc
// @Summary Get my spaces
// @Description Get my spaces
// @Tags me
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {array} models.Space
// @Failure 500 {object} models.ErrorResponse
// @Router /api/me/spaces [get]
func (mc *MeController) GetMySpaces(ctx *fiber.Ctx) error {
	logger := mc.Logger.With().Str("event", "api.spaces.get").Logger()

	userId := ctx.Locals("user_id").(string)
	groups, err := mc.UserService.GetGroups(userId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user groups")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	spaces, err := mc.SpaceService.GetSpacesForUser(userId, groups)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user spaces")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("user", userId).Msg("User spaces retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(spaces)
}

// GetMyProfile godoc
// @Summary Get my profile
// @Description Get my profile
// @Tags me
// @Accept json
// @Produce json
// @Param userId path string true "User Id"
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /api/me/profile [get]
func (mc *MeController) GetMyProfile(ctx *fiber.Ctx) error {
	logger := mc.Logger.With().Str("event", "api.me.get").Logger()

	userId := ctx.Locals("user_id").(string)
	user, err := mc.UserService.GetById(userId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user profile")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	user.Password = ""
	logger.Debug().Str("user", userId).Msg("User profile retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// GetMyFavorites godoc
// @Summary Get my favorites
// @Description Get my favorites
// @Tags me
// @Accept json
// @Produce json
// @Param userId path string true "User Id"
// @Success 200 {array} models.Favorite
// @Failure 500 {object} models.ErrorResponse
// @Router /api/me/favorites [get]
func (mc *MeController) GetMyFavorites(ctx *fiber.Ctx) error {
	logger := mc.Logger.With().Str("event", "api.me.get").Logger()

	userId := ctx.Locals("user_id").(string)
	favorites, err := mc.FavoriteService.GetFavoritesByUserId(userId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user favorites")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("user", userId).Msg("User favorites retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(favorites)
}

// AddFavorite godoc
// @Summary Add favorite with document_id
// @Description Add favorite with document_id
// @Tags me
// @Accept json
// @Produce json
// @Param userId path string true "User Id"
// @Param documentId path string true "Document Id"
// @Success 200 {array} models.Favorite
// @Failure 500 {object} models.ErrorResponse
// @Router /api/me/favorites/{documentId} [post]
func (mc *MeController) AddFavorite(ctx *fiber.Ctx) error {
	logger := mc.Logger.With().Str("event", "api.me.add_favorite").Logger()

	userId := ctx.Locals("user_id").(string)
	documentId := ctx.Params("documentId")

	_, err := mc.FavoriteService.CreateFavorite(models.Favorite{
		UserId:     userId,
		DocumentId: documentId,
	})
	if err != nil {
		logger.Error().Err(err).Msg("Error adding favorite")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	favorites, err := mc.FavoriteService.GetFavoritesByUserId(userId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user favorites")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("user", userId).Msg("User favorite added successfully")
	return ctx.Status(fiber.StatusOK).JSON(favorites)
}

// UnFavorite godoc
// @Summary Unfavorite with document_id
// @Description Unfavorite with document_id
// @Tags me
// @Accept json
// @Produce json
// @Param userId path string true "User Id"
// @Param documentId path string true "Document Id"
// @Success 200 {array} models.Favorite
// @Failure 500 {object} models.ErrorResponse
// @Router /api/me/favorites/{documentId} [delete]
func (mc *MeController) UnFavorite(ctx *fiber.Ctx) error {
	logger := mc.Logger.With().Str("event", "api.me.unfavorite").Logger()

	userId := ctx.Locals("user_id").(string)
	documentId := ctx.Params("documentId")

	err := mc.FavoriteService.UnFavorite(userId, documentId)
	if err != nil {
		logger.Error().Err(err).Msg("Error removing favorite")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	favorites, err := mc.FavoriteService.GetFavoritesByUserId(userId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user favorites")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("user", userId).Msg("User favorite removed successfully")
	return ctx.Status(fiber.StatusOK).JSON(favorites)
}
