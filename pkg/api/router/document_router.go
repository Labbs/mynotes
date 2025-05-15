package router

import (
	"github.com/labbs/mynotion/pkg/api/middleware"
	"github.com/labbs/mynotion/pkg/api/v1/controller"
	"github.com/labbs/mynotion/pkg/repository"
	"github.com/labbs/mynotion/pkg/service"
)

func NewDocumentRouter(config *Config) {
	// Set up the document routes
	config.Logger.Info().Msg("Setting up document routes")

	// initialize the document repository
	dr := repository.NewDocumentRepository(config.Db)

	// initialize the user repository with the database connection
	c := controller.DocumentController{
		DocumentService: service.NewDocumentService(dr),
		Logger:          config.Logger,
	}

	v1Document := config.Fiber.Group("/api/v1/document", middleware.JwtAuthMiddleware(config.Logger, service.NewSessionService(repository.NewSessionRepository(config.Db))), middleware.RBACCheckMiddleware(config.Logger))
	v1Document.Get("/space/:spaceId", c.GetDocumentsFromSpace)
	v1Document.Get("/slug/:slug", c.GetDocumentBySlug)
	v1Document.Get("/:documentId", c.GetDocumentById)
	v1Document.Post("/", c.CreateDocument)
	v1Document.Put("/:documentId", c.UpdateDocument)
}
