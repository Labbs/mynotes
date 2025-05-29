package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/labbs/mynotes/internal/shortuuid"
	"github.com/labbs/mynotes/pkg/models"
	"github.com/rs/zerolog"
)

type DocumentController struct {
	DocumentService models.DocumentService
	Logger          zerolog.Logger
}

// GetDocumentsFromSpace godoc
// @Summary Get document from space
// @Description Get document from space
// @Tags document
// @Accept json
// @Produce json
// @Param spaceId path string true "Space Id"
// @Success 200 {array} models.Document
// @Failure 500 {object} models.ErrorResponse
// @Router /api/document/space/{spaceId} [get]
func (dc *DocumentController) GetDocumentsFromSpace(ctx *fiber.Ctx) error {
	logger := dc.Logger.With().Str("event", "api.documents.get").Logger()

	spaceId := ctx.Params("spaceId")
	documents, err := dc.DocumentService.GetDocumentsFirstLevelForSpace(spaceId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting documents from space")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("space", spaceId).Msg("Documents retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(documents)
}

// GetDocumentsFromParentDocument godoc
// @Summary Get documents from parent document
// @Description Get documents from parent document
// @Tags document
// @Accept json
// @Produce json
// @Param spaceId path string true "Space Id"
// @Param documentId path string true "Document Id"
// @Success 200 {array} models.Document
// @Failure 500 {object} models.ErrorResponse
func (dc *DocumentController) GetDocumentsFromParentDocument(ctx *fiber.Ctx) error {
	logger := dc.Logger.With().Str("event", "api.documents.get").Logger()

	spaceId := ctx.Params("spaceId")
	documentId := ctx.Params("documentId")
	documents, err := dc.DocumentService.GetDocumentsFirstLevelByDocumentId(spaceId, documentId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting documents from parent document")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("space", spaceId).Str("document", documentId).Msg("Documents retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(documents)
}

// GetDocumentById godoc
// @Summary Get document by id
// @Description Get document by id
// @Tags document
// @Accept json
// @Produce json
// @Param documentId path string true "Document Id"
// @Success 200 {object} models.Document
// @Failure 500 {object} models.ErrorResponse
// @Router /api/document/{documentId} [get]
func (dc *DocumentController) GetDocumentById(ctx *fiber.Ctx) error {
	logger := dc.Logger.With().Str("event", "api.documents.get").Logger()

	documentId := ctx.Params("documentId")
	document, err := dc.DocumentService.GetDocumentById(documentId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting document by id")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("document", documentId).Msg("Document retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(document)
}

// GetDocumentBySlug godoc
// @Summary Get document by slug
// @Description Get document by slug
// @Tags document
// @Accept json
// @Produce json
// @Param slug path string true "Document Slug"
// @Success 200 {object} models.Document
// @Failure 500 {object} models.ErrorResponse
// @Router /api/document/slug/{slug} [get]
func (dc *DocumentController) GetDocumentBySlug(ctx *fiber.Ctx) error {
	logger := dc.Logger.With().Str("event", "api.documents.get").Logger()

	slug := ctx.Params("slug")
	document, err := dc.DocumentService.GetDocumentBySlug(slug)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting document by slug")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("document", slug).Msg("Document retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(document)
}

// CreateDocument godoc
// @Summary Create document
// @Description Create document
// @Tags document
// @Accept json
// @Produce json
// @Param document body models.Document true "Document"
// @Success 201 {object} models.Document
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/document [post]
func (dc *DocumentController) CreateDocument(ctx *fiber.Ctx) error {
	logger := dc.Logger.With().Str("event", "api.documents.create").Logger()

	var document models.Document
	if err := ctx.BodyParser(&document); err != nil {
		logger.Error().Err(err).Msg("Error parsing request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	document, err := dc.DocumentService.CreateDocument(document)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating document")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("document", document.Id).Msg("Document created successfully")
	return ctx.Status(fiber.StatusCreated).JSON(document)
}

// UpdateDocument godoc
// @Summary Update document
// @Description Update document
// @Tags document
// @Accept json
// @Produce json
// @Param documentId path string true "Document Id"
// @Param document body models.Document true "Document"
// @Success 200 {object} models.Document
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/document/{documentId} [put]
func (dc *DocumentController) UpdateDocument(ctx *fiber.Ctx) error {
	logger := dc.Logger.With().Str("event", "api.documents.update").Logger()

	documentId := ctx.Params("documentId")
	var documentRequest models.Document
	if err := ctx.BodyParser(&documentRequest); err != nil {
		logger.Error().Err(err).Msg("Error parsing request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	document, err := dc.DocumentService.GetDocumentById(documentId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting document by id")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	if documentRequest.Name == "" {
		logger.Error().Msg("Document name is required")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Document name is required"})
	}

	if documentRequest.Name != document.Name {
		document.Slug = slug.Make(documentRequest.Name + "-" + shortuuid.GenerateShortUUID())
		document.Name = documentRequest.Name
	}

	document.Content = documentRequest.Content
	document.Config = documentRequest.Config

	document, err = dc.DocumentService.UpdateDocument(document)
	if err != nil {
		logger.Error().Err(err).Msg("Error updating document")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Debug().Str("document", document.Id).Msg("Document updated successfully")
	return ctx.Status(fiber.StatusOK).JSON(document)
}
