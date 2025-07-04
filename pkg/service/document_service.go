package service

import (
	"github.com/labbs/zotion/pkg/models"
)

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

func (s *documentService) GetDocumentsFirstLevelByDocumentId(documentId string) ([]models.Document, error) {
	return s.documentRepository.GetDocumentsFirstLevelByDocumentId(documentId)
}

func (s *documentService) GetDocumentBySlug(slug string) (models.Document, error) {
	return s.documentRepository.GetDocumentBySlug(slug)
}

func (s *documentService) GetDocumentById(id string) (models.Document, error) {
	return s.documentRepository.GetDocumentById(id)
}

func (s *documentService) UpdateDocument(document models.Document) (models.Document, error) {
	return s.documentRepository.UpdateDocument(document)
}

func (s *documentService) DeleteDocument(id string) error {
	// get document by id
	document, err := s.documentRepository.GetDocumentById(id)
	// List all child documents and delete them recursively
	if err != nil {
		return err
	}
	// infinite loop to delete all child documents
	for {
		childDocuments, err := s.documentRepository.GetDocumentsFirstLevelByDocumentId(document.Id)
		if err != nil {
			return err
		}
		if len(childDocuments) == 0 {
			break
		}
		for _, childDocument := range childDocuments {
			if err := s.DeleteDocument(childDocument.Id); err != nil {
				return err
			}
		}
	}
	return s.documentRepository.DeleteDocument(id)
}

func (s *documentService) GetAllDocuments() ([]models.Document, error) {
	return s.documentRepository.GetAllDocuments()
}

func (s *documentService) GetAllDeletedDocument() ([]models.Document, error) {
	return s.documentRepository.GetAllDeletedDocument()
}

func (s *documentService) RestoreDocument(id string) error {
	return s.documentRepository.RestoreDocument(id)
}

func (s *documentService) GetDocumentsBySpaceId(spaceId string) ([]models.Document, error) {
	return s.documentRepository.GetDocumentsBySpaceId(spaceId)
}
