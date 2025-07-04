package service

import "github.com/labbs/zotion/pkg/models"

type favoriteService struct {
	favoriteRepository models.FavoriteRepository
}

func NewFavoriteService(favoriteRepository models.FavoriteRepository) *favoriteService {
	return &favoriteService{favoriteRepository: favoriteRepository}
}

func (s *favoriteService) GetFavoritesByUserId(userId string) ([]models.Favorite, error) {
	return s.favoriteRepository.GetFavoritesByUserId(userId)
}

func (s *favoriteService) IsFavorite(userId string, documentId string) (bool, error) {
	return s.favoriteRepository.IsFavorite(userId, documentId)
}

func (s *favoriteService) CreateFavorite(favorite models.Favorite) (models.Favorite, error) {
	return s.favoriteRepository.CreateFavorite(favorite)
}

func (s *favoriteService) UnFavorite(userId string, documentId string) error {
	return s.favoriteRepository.UnFavorite(userId, documentId)
}

func (s *favoriteService) DeleteFavorite(id string) error {
	return s.favoriteRepository.DeleteFavorite(id)
}

func (s *favoriteService) DeleteFavoritesByDocumentId(documentId string) error {
	return s.favoriteRepository.DeleteFavoritesByDocumentId(documentId)
}
