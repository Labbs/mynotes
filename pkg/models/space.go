package models

import (
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/gosimple/slug"
	"github.com/labbs/mynotes/internal/shortuuid"
	"github.com/labbs/mynotes/pkg/caching"
	"gorm.io/gorm"
)

// Space is a model for a space
type Space struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Icon        string    `json:"icon"`
	IconColor   string    `json:"icon_color"`
	Description string    `json:"description"`
	Type        SpaceType `json:"type"`

	Documents []Document `json:"documents"`

	Members Members `json:"members"`

	// MembersWithUsers is used to return the members with user information
	MembersWithUsersOrGroups MembersWithUsersOrGroups `json:"members_with_users_or_groups" gorm:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// TableName returns the name of the table
func (s Space) TableName() string {
	return "space"
}

// BeforeCreate is a hook that runs before creating a space
func (s *Space) BeforeCreate(tx *gorm.DB) error {
	s.Id = utils.UUIDv4()
	s.Slug = slug.Make(s.Name + "-" + shortuuid.GenerateShortUUID())
	return nil
}

// BeforeUpdate is a hook that runs before updating a space
func (s *Space) BeforeUpdate(tx *gorm.DB) error {
	// if name changed, update slug
	if tx.Statement.Changed("name") {
		s.Slug = slug.Make(s.Name + "-" + shortuuid.GenerateShortUUID())
	}
	return nil
}

func (s *Space) AfterCreate(tx *gorm.DB) error {
	caching.Cache.Set("space:"+s.Id, s.Members)
	return nil
}

func (s *Space) AfterUpdate(tx *gorm.DB) error {
	caching.Cache.Set("space:"+s.Id, s.Members)
	return nil
}

// AfterDelete is a hook that runs after deleting a space
func (s *Space) AfterDelete(tx *gorm.DB) error {
	// Remove the space from the cache
	caching.Cache.Delete("space:" + s.Id)
	return nil
}

// SpaceType is the type of space
type SpaceType string

const (
	// SpaceTypePublic is the public space type
	SpaceTypePublic SpaceType = "public"
	// SpaceTypePrivate is the private space type
	SpaceTypePrivate SpaceType = "private"
	// SpaceTypeRestricted is the restricted space type
	SpaceTypeRestricted SpaceType = "restricted"
)

// SpaceRepository is the repository for spaces
type SpaceRepository interface {
	GetSpacesForUser(userId string, groups []Group) ([]Space, error)
	GetSpaceById(spaceId string) (Space, error)
	CreateSpace(space Space) (Space, error)
	IsMember(spaceId, userId string) (bool, error)
	GetAllSpaces() ([]Space, error)
}

// SpaceService is the service for spaces
type SpaceService interface {
	GetSpacesForUser(userId string, groups []Group) ([]Space, error)
	GetSpaceById(spaceId string) (Space, error)
	CreateSpace(space Space) (Space, error)
	IsMember(spaceId, userId string) (bool, error)
	GetAllSpaces() ([]Space, error)
}
