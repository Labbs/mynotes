package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

type PublicController struct {
	DocumentController models.DocumentService
	Logger             zerolog.Logger
}

// GetPublicDocumentBySlug godoc
// @Summary Get public document by slug
// @Description Get public document by slug
// @Tags document
// @Accept json
// @Produce json
// @Param slug path string true "Document Slug"
// @Success 200 {object} models.Document
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/public/document/slug/{slug} [get]
func (pc *PublicController) GetPublicDocumentBySlug(ctx *fiber.Ctx) error {
	logger := pc.Logger.With().Str("event", "api.public_documents.get").Logger()

	slug := ctx.Params("slug")
	document, err := pc.DocumentController.GetDocumentBySlug(slug)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting document by slug")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	if !document.Public {
		logger.Error().Msg("Document is not public")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Document not exist"})
	}

	logger.Debug().Str("document", slug).Msg("Document retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(document)
}
