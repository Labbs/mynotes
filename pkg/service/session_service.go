package service

import "github.com/labbs/zotion/pkg/models"

type sessionService struct {
	sessionRepository models.SessionRepository
}

func NewSessionService(sr models.SessionRepository) models.SessionService {
	return &sessionService{
		sessionRepository: sr,
	}
}

func (s *sessionService) GetById(id string) (models.Session, error) {
	return s.sessionRepository.GetById(id)
}

func (s *sessionService) GetAllByUserId(userId string) ([]models.Session, error) {
	return s.sessionRepository.GetAllByUserId(userId)
}

func (s *sessionService) Create(session *models.Session) error {
	return s.sessionRepository.Create(session)
}

func (s *sessionService) Update(session *models.Session) error {
	return s.sessionRepository.Update(session)
}

func (s *sessionService) Delete(id string) error {
	return s.sessionRepository.Delete(id)
}
