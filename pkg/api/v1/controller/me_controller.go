package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/mynotes/pkg/models"
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
	groups, err := mc.UserService.GetGroupsByUserId(userId)
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

	// check if is_admin is set in context
	if ctx.Locals("is_admin") != nil {
		if ctx.Locals("is_admin").(bool) {
			user.IsAdmin = true
		}
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

// GetMyPreferences godoc
// @Summary Get my preferences
// @Description Get my preferences
// @Tags me
// @Accept json
// @Produce json
// @Param userId path string true "User Id"
// @Success 200 {object} models.JSONB
// @Failure 500 {object} models.ErrorResponse
// @Router /api/me/preferences [get]
func (mc *MeController) GetMyPreferences(ctx *fiber.Ctx) error {
	logger := mc.Logger.With().Str("event", "api.me.get_preferences").Logger()

	userId := ctx.Locals("user_id").(string)
	preferences, err := mc.UserService.GetPreferencesById(userId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user preferences")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("user", userId).Msg("User preferences retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(preferences)
}

// UpdateMyPreferences godoc
// @Summary Update my preferences
// @Description Update my preferences
// @Tags me
// @Accept json
// @Produce json
// @Param userId path string true "User Id"
// @Param preferences body models.JSONB true "Preferences"
// @Success 200 {object} models.JSONB
// @Failure 500 {object} models.ErrorResponse
// @Router /api/me/preferences [put]
func (mc *MeController) UpdateMyPreferences(ctx *fiber.Ctx) error {
	logger := mc.Logger.With().Str("event", "api.me.update_preferences").Logger()

	userId := ctx.Locals("user_id").(string)
	var preferences models.JSONB

	if err := ctx.BodyParser(&preferences); err != nil {
		logger.Error().Err(err).Msg("Error parsing preferences")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := mc.UserService.UpdatePreferences(userId, preferences); err != nil {
		logger.Error().Err(err).Msg("Error updating user preferences")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("user", userId).Msg("User preferences updated successfully")
	return ctx.Status(fiber.StatusOK).JSON(preferences)
}
