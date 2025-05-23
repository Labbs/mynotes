package repository

import (
	"time"

	"github.com/labbs/mynotes/pkg/models"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

// NewSessionRepository creates a new instance of sessionRepository and expose the functions available to the session
// repository. It takes a gorm.DB instance as a parameter to interact with the database.
// The sessionRepository struct implements the SessionRepository interface defined in the models package.
func NewSessionRepository(db *gorm.DB) *sessionRepository {
	return &sessionRepository{db: db}
}

// GetById retrieves a session from the database by their ID.
// It takes an id string as a parameter and returns a session and an error.
// If the session is found, it returns the session and a nil error. If not, it returns an empty session and an error.
// The error is nil if the session is found, otherwise it contains the error message.
// The ID is the unique identifier for the session in the database.
// The session is returned as a models.Session struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the session is found, otherwise it contains the error message.
func (r *sessionRepository) GetById(id string) (models.Session, error) {
	var session models.Session
	err := r.db.Where("id = ?", id).First(&session).Error
	return session, err
}

// GetAllByUserId retrieves all sessions from the database by their user ID.
// It takes a userId string as a parameter and returns a slice of sessions and an error.
// If the sessions are found, it returns the sessions and a nil error. If not, it returns an empty slice and an error.
// The error is nil if the sessions are found, otherwise it contains the error message.
// The userId is the unique identifier for the user in the database.
// The sessions are returned as a slice of models.Session structs.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the sessions are found, otherwise it contains the error message.
// The sessions are returned as a slice of models.Session structs.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the sessions are found, otherwise it contains the error message.
func (r *sessionRepository) GetAllByUserId(userId string) ([]models.Session, error) {
	var sessions []models.Session
	err := r.db.Where("user_id = ?", userId).Find(&sessions).Error
	return sessions, err
}

// Create creates a new session in the database.
// It takes a session pointer as a parameter and returns an error.
// If the session is created successfully, it returns a nil error. If not, it returns an error.
// The session is passed as a pointer to avoid copying the entire struct.
// The session is returned as a models.Session struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the session is created successfully, otherwise it contains the error message.
func (r *sessionRepository) Create(session *models.Session) error {
	return r.db.Debug().Create(session).Error
}

// Update updates an existing session in the database.
// It takes a session pointer as a parameter and returns an error.
// If the session is updated successfully, it returns a nil error. If not, it returns an error.
// The session is passed as a pointer to avoid copying the entire struct.
// The session is returned as a models.Session struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the session is updated successfully, otherwise it contains the error message.
func (r *sessionRepository) Update(session *models.Session) error {
	return r.db.Debug().Save(session).Error
}

// Delete deletes a session from the database by their ID.
// It takes an id string as a parameter and returns an error.
// If the session is deleted successfully, it returns a nil error. If not, it returns an error.
// The id is the unique identifier for the session in the database.
// The session is returned as a models.Session struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the session is deleted successfully, otherwise it contains the error message.
func (r *sessionRepository) Delete(id string) error {
	return r.db.Debug().Where("id = ?", id).Delete(&models.Session{}).Error
}

// DeleteExpiredSessions deletes expired sessions from the database.
// It takes an expirationTime time.Time as a parameter and returns an error.
// If the expired sessions are deleted successfully, it returns a nil error. If not, it returns an error.
// The expirationTime is the time before which the sessions are considered expired.
// The session is returned as a models.Session struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the expired sessions are deleted successfully, otherwise it contains the error message.
func (r *sessionRepository) DeleteExpiredSessions(expirationTime time.Time) error {
	return r.db.Debug().Where("created_at < ?", expirationTime).Delete(&models.Session{}).Error
}
