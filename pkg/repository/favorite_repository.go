package repository

import (
	"github.com/labbs/mynotion/pkg/models"
	"gorm.io/gorm"
)

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *favoriteRepository {
	return &favoriteRepository{db: db}
}

func (r *favoriteRepository) GetFavoritesByUserId(userId string) ([]models.Favorite, error) {
	var favorites []models.Favorite
	err := r.db.Debug().Table("favorite").Where("user_id = ?", userId).Order("position ASC").Find(&favorites).Error
	return favorites, err
}

func (r *favoriteRepository) IsFavorite(userId string, documentId string) (bool, error) {
	var count int64
	err := r.db.Table("favorite").Where("user_id = ? AND document_id = ?", userId, documentId).Count(&count).Error
	return count > 0, err
}

func (r *favoriteRepository) CreateFavorite(favorite models.Favorite) (models.Favorite, error) {
	return favorite, r.db.Create(&favorite).Error
}

func (r *favoriteRepository) UnFavorite(userId string, documentId string) error {
	return r.db.Where("user_id = ? AND document_id = ?", userId, documentId).Delete(&models.Favorite{}).Error
}

func (r *favoriteRepository) DeleteFavorite(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Favorite{}).Error
}
