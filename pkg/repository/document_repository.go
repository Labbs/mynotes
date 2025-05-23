package repository

import (
	"github.com/labbs/mynotes/pkg/models"
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

func (r *documentRepository) GetDocumentById(id string, preloadBlock bool) (models.Document, error) {
	var document models.Document
	query := r.db.Debug().Table("document")
	if preloadBlock {
		query = query.Preload("Blocks", func(db *gorm.DB) *gorm.DB {
			return db.Order("position ASC")
		})
	}
	err := query.First(&document, "id = ?", id).Error
	return document, err
}

func (r *documentRepository) GetDocumentBySlug(slug string, preloadBlock bool) (models.Document, error) {
	var document models.Document
	query := r.db.Debug().Table("document")
	if preloadBlock {
		query = query.Preload("Blocks", func(db *gorm.DB) *gorm.DB {
			return db.Order("position ASC")
		})
	}
	err := query.First(&document, "slug = ?", slug).Error
	return document, err
}
