package repository

import (
	"github.com/labbs/zotion/pkg/models"
	"gorm.io/gorm"
)

type documentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *documentRepository {
	return &documentRepository{db: db}
}

func (r *documentRepository) CreateDocument(document models.Document) (models.Document, error) {
	err := r.db.Debug().Table("document").Create(&document).Error
	return document, err
}

func (r *documentRepository) UpdateDocument(document models.Document) (models.Document, error) {
	err := r.db.Debug().Table("document").Save(&document).Error
	return document, err
}

func (r *documentRepository) DeleteDocument(id string) error {
	return r.db.Table("document").Where("id = ?", id).Delete(&models.Document{}).Error
}

func (r *documentRepository) GetDocumentsFirstLevelForSpace(spaceId string) ([]models.Document, error) {
	var documents []models.Document
	err := r.db.Debug().Table("document").Where("space_id = ? AND (parent_id IS NULL OR parent_id = '')", spaceId).Find(&documents).Error
	return documents, err
}

func (r *documentRepository) GetDocumentsFirstLevelByDocumentId(documentId string) ([]models.Document, error) {
	var documents []models.Document
	err := r.db.Debug().Table("document").Where("parent_id = ?", documentId).Find(&documents).Error
	return documents, err
}

func (r *documentRepository) GetDocumentById(id string) (models.Document, error) {
	var document models.Document
	err := r.db.Debug().Table("document").First(&document, "id = ?", id).Error
	return document, err
}

func (r *documentRepository) GetDocumentBySlug(slug string) (models.Document, error) {
	var document models.Document
	err := r.db.Debug().Table("document").First(&document, "slug = ?", slug).Error
	return document, err
}

func (r *documentRepository) GetAllDocuments() ([]models.Document, error) {
	var documents []models.Document
	err := r.db.Debug().Table("document").Select("id", "name", "type", "updated_at").Find(&documents).Error
	return documents, err
}

func (r *documentRepository) GetAllDeletedDocument() ([]models.Document, error) {
	var documents []models.Document
	err := r.db.Debug().Table("document").Where("deleted_at IS NOT NULL").Find(&documents).Error
	return documents, err
}

func (r *documentRepository) RestoreDocument(id string) error {
	return r.db.Debug().Table("document").Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *documentRepository) GetDocumentsBySpaceId(spaceId string) ([]models.Document, error) {
	var documents []models.Document
	err := r.db.Debug().Table("document").Where("space_id = ?", spaceId).Find(&documents).Error
	return documents, err
}
