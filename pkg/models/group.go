package models

import (
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"
)

type Group struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Users []User `json:"users" gorm:"many2many:user_group;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (g Group) TableName() string {
	return "group"
}

func (g *Group) BeforeCreate(tx *gorm.DB) error {
	g.Id = utils.UUIDv4()
	return nil
}

func (g *Group) BeforeUpdate(tx *gorm.DB) error {
	return nil
}

// GroupRepository is the repository for groups
type GroupRepository interface {
	Create(group Group) (Group, error)
	GetById(id string) (Group, error)
	GetWithUsers(id string) (Group, error)
	GetAll() ([]Group, error)
	Update(group Group) (Group, error)
	Delete(id string) error
}

// GroupService is the service for groups
type GroupService interface {
	CreateGroup(group Group) (Group, error)
	GetGroupById(id string) (Group, error)
	GetGroupWithUsers(id string) (Group, error)
	GetAllGroups() ([]Group, error)
	UpdateGroup(group Group) (Group, error)
	DeleteGroup(id string) error
}
