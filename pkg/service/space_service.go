package service

import "github.com/labbs/mynotes/pkg/models"

type spaceService struct {
	spaceRepository models.SpaceRepository
}

func NewSpaceService(spaceRepository models.SpaceRepository) *spaceService {
	return &spaceService{spaceRepository: spaceRepository}
}

func (s *spaceService) GetSpacesForUser(userId string, groups []models.Group) ([]models.Space, error) {
	spaces, err := s.spaceRepository.GetSpacesForUser(userId, groups)
	if err != nil {
		return []models.Space{}, err
	}

	return spaces, nil
}

func (s *spaceService) GetSpaceById(spaceId string) (models.Space, error) {
	space, err := s.spaceRepository.GetSpaceById(spaceId)
	if err != nil {
		return models.Space{}, err
	}

	return space, nil
}

func (s *spaceService) CreateSpace(space models.Space) (models.Space, error) {
	space, err := s.spaceRepository.CreateSpace(space)
	if err != nil {
		return models.Space{}, err
	}

	return space, nil
}

func (s *spaceService) IsMember(spaceId, userId string) (bool, error) {
	return s.spaceRepository.IsMember(spaceId, userId)
}

func (s *spaceService) GetAllSpaces() ([]models.Space, error) {
	spaces, err := s.spaceRepository.GetAllSpaces()
	if err != nil {
		return []models.Space{}, err
	}

	return spaces, nil
}
