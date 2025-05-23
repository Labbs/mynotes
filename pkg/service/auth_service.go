package service

import (
	"github.com/gofiber/fiber/v2/utils"
	"github.com/labbs/mynotes/internal/tokenutil"
	"github.com/labbs/mynotes/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository     models.UserRepository
	spaceRepository    models.SpaceRepository
	documentRepository models.DocumentRepository
}

func NewAuthService(ur models.UserRepository, sr models.SpaceRepository, dr models.DocumentRepository) models.AuthService {
	return &authService{
		userRepository:     ur,
		spaceRepository:    sr,
		documentRepository: dr,
	}
}

func (s *authService) Login(request models.LoginRequest) (models.LoginResponse, error) {
	user, err := s.userRepository.GetByEmail(request.Email)
	if err != nil {
		return models.LoginResponse{}, err
	}

	if !user.Active {
		return models.LoginResponse{}, models.ErrUserDisabled{
			Message: "User is disabled",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return models.LoginResponse{}, err
	}

	sessionId := utils.UUIDv4()

	accessToken, err := tokenutil.CreateAccessToken(user.Id, sessionId)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		Token:     accessToken,
		SessionId: sessionId,
		UserId:    user.Id,
	}, nil
}

func (s *authService) Register(request models.RegisterRequest) (models.RegisterResponse, error) {
	_, err := s.userRepository.GetByEmail(request.Email)
	if err == nil || err.Error() != "record not found" {
		return models.RegisterResponse{}, models.ErrUserDisabled{
			Message: "User already exists",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	newUser := &models.User{
		Id:       utils.UUIDv4(),
		Email:    request.Email,
		Name:     request.Name,
		Password: string(hashedPassword),
		Active:   true,
	}

	err = s.userRepository.Create(newUser)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	privateSpace := &models.Space{
		Id:      utils.UUIDv4(),
		Name:    "My Private Space",
		Type:    models.SpaceTypePrivate,
		Members: models.Members{{Id: newUser.Id, Type: models.MemberTypeUser, Access: models.AccessTypeFull}},
	}

	_, err = s.spaceRepository.CreateSpace(*privateSpace)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return models.RegisterResponse{}, nil
}
