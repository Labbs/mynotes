package service

import (
	"github.com/labbs/zotion/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

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

func (s *userService) GetGroupsByUserId(userId string) ([]models.Group, error) {
	return s.userRepository.GetGroupsByUserId(userId)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *userService) GetAllInactiveUsers() ([]models.User, error) {
	return s.userRepository.GetAllInactiveUsers()
}

func (s *userService) GetPreferencesById(id string) (models.JSONB, error) {
	return s.userRepository.GetPreferencesById(id)
}

func (s *userService) UpdatePreferences(id string, preferences models.JSONB) error {
	return s.userRepository.UpdatePreferences(id, preferences)
}

func (s *userService) GetUserWithGroups(id string) (models.User, error) {
	return s.userRepository.GetUserWithGroups(id)
}

func (s *userService) GetUsersWithGroups() ([]models.User, error) {
	return s.userRepository.GetUsersWithGroups()
}

func (s *userService) ChangePassword(userId string, currentPassword string, newPassword string) error {
	user, err := s.userRepository.GetById(userId)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword))
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	_, err = s.userRepository.Update(&user)

	return err
}
