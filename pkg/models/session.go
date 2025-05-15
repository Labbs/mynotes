package models

import "time"

type Session struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`

	UserAgent string `json:"user_agent"`
	IpAddress string `json:"ip_address"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName returns the name of the table in the database.
// This is used by GORM to map the struct to the correct table.
func (s *Session) TableName() string {
	return "session"
}

// SessionRepository defines the methods that a session repository should implement.
type SessionRepository interface {
	GetById(id string) (Session, error)
	GetAllByUserId(userId string) ([]Session, error)
	Create(session *Session) error
	Update(session *Session) error
	Delete(id string) error
}

// SessionService defines the methods that a session service should implement.
type SessionService interface {
	GetById(id string) (Session, error)
	GetAllByUserId(userId string) ([]Session, error)
	Create(session *Session) error
	Update(session *Session) error
	Delete(id string) error
}
