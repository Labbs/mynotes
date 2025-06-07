package service

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/labbs/mynotes/pkg/config"
	"github.com/labbs/mynotes/pkg/models"
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

func (s *documentService) GetDocumentsFirstLevelByDocumentId(spaceId, documentId string) ([]models.Document, error) {
	return s.documentRepository.GetDocumentsFirstLevelByDocumentId(spaceId, documentId)
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
		childDocuments, err := s.documentRepository.GetDocumentsFirstLevelByDocumentId(document.SpaceId, document.Id)
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

func (s *documentService) GetExcalidrawLibsList() ([]string, error) {
	// Ensure the directory exists
	if _, err := os.Stat(config.Document.ExcalidrawLibsPath); os.IsNotExist(err) {
		if err := os.MkdirAll(config.Document.ExcalidrawLibsPath, 0755); err != nil {
			return nil, fmt.Errorf("failed to create excalidraw libraries directory: %w", err)
		}
	}

	// List files in the Excalidraw libraries path
	root := os.DirFS(config.Document.ExcalidrawLibsPath)

	f, err := fs.Glob(root, "*.excalidrawlib")
	if err != nil {
		return []string{}, nil // Return empty list instead of error
	}

	var files []string
	for _, v := range f {
		files = append(files, strings.Split(v, ".")[0])
	}
	return files, nil
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
