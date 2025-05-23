package service

import "github.com/labbs/mynotes/pkg/models"

type documentService struct {
	documentRepository models.DocumentRepository
}

func NewDocumentService(documentRepository models.DocumentRepository) *documentService {
	return &documentService{documentRepository: documentRepository}
}

func (s *documentService) CreateDocument(document models.Document) (models.Document, error) {
	return s.documentRepository.CreateDocument(document)
}

func (s *documentService) GetDocumentsFirstLevelForSpace(spaceId string) ([]models.Document, error) {
	return s.documentRepository.GetDocumentsFirstLevelForSpace(spaceId)
}

func (s *documentService) GetDocumentBySlug(slug string, preloadBlock bool) (models.Document, error) {
	return s.documentRepository.GetDocumentBySlug(slug, preloadBlock)
}

func (s *documentService) GetDocumentById(id string, preloadBlock bool) (models.Document, error) {
	return s.documentRepository.GetDocumentById(id, preloadBlock)
}

func (s *documentService) UpdateDocument(document models.Document) (models.Document, error) {
	return s.documentRepository.UpdateDocument(document)
}
