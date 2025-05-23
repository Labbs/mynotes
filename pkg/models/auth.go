package models

import "github.com/golang-jwt/jwt/v4"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	UserId    string `json:"-"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponse struct{}

type JwtCustomClaims struct {
	SessionId string `json:"session_id"`
	UserId    string `json:"user_id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	SessionId string `json:"session_id"`
	UserId    string `json:"user_id"`
	jwt.RegisteredClaims
}

type ErrUserDisabled struct {
	Message string `json:"message"`
}

func (e ErrUserDisabled) Error() string {
	return e.Message
}

type AuthService interface {
	Login(request LoginRequest) (LoginResponse, error)
	Register(request RegisterRequest) (RegisterResponse, error)
}
