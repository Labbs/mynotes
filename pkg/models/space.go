package models

import (
	"database/sql/driver"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gosimple/slug"
	"github.com/labbs/mynotes/internal/shortuuid"
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

// MemberType is the type of member
type MemberType string

// AccessType is the access type of a member
type AccessType string

type Members []Member

func (m Members) Value() (driver.Value, error) {
	valueString, err := json.Marshal(m)
	return string(valueString), err
}

func (m *Members) Scan(value any) error {
	return json.Unmarshal([]byte(value.(string)), m)
}

// Member is a model for a member
type Member struct {
	Id     string     `json:"id"`
	Type   MemberType `json:"type"`
	Access AccessType `json:"access"`
}

// AccessTypeViewer is the access type viewer
const (
	AccessTypeViewer  AccessType = "viewer"
	AccessTypeEditor  AccessType = "editor"
	AccessTypeComment AccessType = "comment"
	AccessTypeFull    AccessType = "full"
)

// MemberType is the type of member
const (
	MemberTypeUser  MemberType = "user"
	MemberTypeGroup MemberType = "group"
)

// SpaceRepository is the repository for spaces
type SpaceRepository interface {
	GetSpacesForUser(userId string, groups []Group) ([]Space, error)
	GetSpaceById(spaceId string) (Space, error)
	CreateSpace(space Space) (Space, error)
	IsMember(spaceId, userId string) (bool, error)
}

// SpaceService is the service for spaces
type SpaceService interface {
	GetSpacesForUser(userId string, groups []Group) ([]Space, error)
	GetSpaceById(spaceId string) (Space, error)
	CreateSpace(space Space) (Space, error)
	IsMember(spaceId, userId string) (bool, error)
}
