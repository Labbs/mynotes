package models

import (
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"
)

// User is a model for a user
type User struct {
	Id string `json:"id"`

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	AvatarUrl   string `json:"avatar_url"`
	Preferences JSONB  `json:"preferences"`
	Active      bool   `json:"active"`

	Groups []Group `json:"groups" gorm:"many2many:user_group;"`

	Favorites []Favorite `json:"favorites"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName returns the name of the table in the database.
// This is used by GORM to map the struct to the correct table.
func (u *User) TableName() string {
	return "user"
}

// BeforeCreate is a GORM hook that is called before a new record is created.
// It sets the Id of the user to a new UUID.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Id = utils.UUIDv4()
	return nil
}

// UserRepository defines the methods that a user repository should implement.
type UserRepository interface {
	GetByEmailOrUsername(emailOrUsername string) (User, error)
	GetByEmail(email string) (User, error)
	GetById(id string) (User, error)
	Create(user *User) error
	Update(user *User) (User, error)
	Delete(id string) error
	GetOrderedFavorites(userId string) ([]Favorite, error)
	GetGroups(userId string) ([]Group, error)
}

// UserService defines the methods that a user service should implement.
type UserService interface {
	GetByEmailOrUsername(emailOrUsername string) (User, error)
	GetByEmail(email string) (User, error)
	GetById(id string) (User, error)
	Create(user *User) error
	Update(user *User) (User, error)
	Delete(id string) error
	GetOrderedFavorites(userId string) ([]Favorite, error)
	GetGroups(userId string) ([]Group, error)
}
