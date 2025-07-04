package models

import (
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"
)

type Favorite struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	DocumentId string `json:"document_id"`
	DatabaseId string `json:"database_id"`

	Position string `json:"position"`

	Document Document `json:"document" gorm:"foreignKey:DocumentId;references:Id"`

	CreatedAt time.Time `json:"created_at"`
}

func (f Favorite) TableName() string {
	return "favorite"
}

// BeforeCreate is a hook that runs before creating a user
func (u *Favorite) BeforeCreate(tx *gorm.DB) error {
	u.Id = utils.UUIDv4()
	return nil
}

type FavoriteRepository interface {
	GetFavoritesByUserId(userId string) ([]Favorite, error)
	CreateFavorite(favorite Favorite) (Favorite, error)
	DeleteFavorite(id string) error
	IsFavorite(userId string, documentId string) (bool, error)
	UnFavorite(userId string, documentId string) error
	DeleteFavoritesByDocumentId(documentId string) error
}

type FavoriteService interface {
	GetFavoritesByUserId(userId string) ([]Favorite, error)
	CreateFavorite(favorite Favorite) (Favorite, error)
	DeleteFavorite(id string) error
	IsFavorite(userId string, documentId string) (bool, error)
	UnFavorite(userId string, documentId string) error
	DeleteFavoritesByDocumentId(documentId string) error
}
