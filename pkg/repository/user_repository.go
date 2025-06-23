package repository

import (
	"github.com/labbs/zotion/pkg/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository and expose the functions available to the user
// repository. It takes a gorm.DB instance as a parameter to interact with the database.
// The userRepository struct implements the UserRepository interface defined in the models package.
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// GetByEmailOrUsername retrieves a user from the database by their email or username.
// It takes an emailOrUsername string as a parameter and returns a user and an error.
// If the user is found, it returns the user and a nil error. If not, it returns an empty user and an error.
// The error is nil if the user is found, otherwise it contains the error message.
func (r *userRepository) GetByEmailOrUsername(emailOrUsername string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR username = ?", emailOrUsername, emailOrUsername).First(&user).Error
	return user, err
}

// GetByEmail returns a user from the database by their email.
// It takes an email string as a parameter and returns a user and an error.
// If the user is found, it returns the user and a nil error. If not, it returns an empty user and an error.
// The error is nil if the user is found, otherwise it contains the error message.
// The email is the unique identifier for the user in the database.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is found, otherwise it contains the error message.
// The email is the unique identifier for the user in the database.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is found, otherwise it contains the error message.
func (r *userRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetById retrieves a user from the database by their ID.
// It takes an id string as a parameter and returns a user and an error.
// If the user is found, it returns the user and a nil error. If not, it returns an empty user and an error.
// The error is nil if the user is found, otherwise it contains the error message.
// The ID is the unique identifier for the user in the database.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is found, otherwise it contains the error message.
func (r *userRepository) GetById(id string) (models.User, error) {
	var user models.User
	err := r.db.Debug().Where("id = ?", id).First(&user).Error
	return user, err
}

// Create creates a new user in the database.
// It takes a user pointer as a parameter and returns an error.
// If the user is created successfully, it returns a nil error. If not, it returns an error.
// The user is passed as a pointer to avoid copying the entire struct.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is created successfully, otherwise it contains the error message.
// The user is created in the database using the GORM Create method.
// The Create method takes a pointer to the user struct and creates a new record in the database.
func (r *userRepository) Create(user *models.User) error {
	err := r.db.Debug().Create(user).Error
	return err
}

// Update updates a user
func (r *userRepository) Update(user *models.User) (models.User, error) {
	return *user, r.db.Debug().Save(user).Error
}

// Delete deletes a user
func (r *userRepository) Delete(id string) error {
	return r.db.Debug().Where("id = ?", id).Delete(&models.User{}).Error
}

// GetGroups returns the groups of a user
func (r *userRepository) GetGroups(id string) ([]models.Group, error) {
	var groups []models.Group
	err := r.db.Debug().Unscoped().Table("user_group").Select("user_id, group_id").Where("user_id = ?", id).Find(&groups).Error
	return groups, err
}
