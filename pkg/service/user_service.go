package service

import "github.com/labbs/mynotes/pkg/models"

type userService struct {
	userRepository models.UserRepository
}

func NewUserService(ur models.UserRepository) models.UserService {
	return &userService{
		userRepository: ur,
	}
}

func (s *userService) GetByEmailOrUsername(emailOrUsername string) (models.User, error) {
	return s.userRepository.GetByEmailOrUsername(emailOrUsername)
}

func (s *userService) GetByEmail(email string) (models.User, error) {
	return s.userRepository.GetByEmail(email)
}

func (s *userService) GetById(id string) (models.User, error) {
	return s.userRepository.GetById(id)
}

func (s *userService) Create(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) Update(user *models.User) (models.User, error) {
	return s.userRepository.Update(user)
}

func (s *userService) Delete(id string) error {
	return s.userRepository.Delete(id)
}

func (s *userService) GetOrderedFavorites(userId string) ([]models.Favorite, error) {
	return s.userRepository.GetOrderedFavorites(userId)
}

func (s *userService) GetGroups(userId string) ([]models.Group, error) {
	return s.userRepository.GetGroups(userId)
}
