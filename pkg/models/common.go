package models

import (
	"database/sql/driver"

	"github.com/goccy/go-json"
)

// JSONB is a map of strings to interfaces
type JSONB map[string]any

// Value implements the driver.Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

// Scan implements the sql.Scanner interface
func (j *JSONB) Scan(value any) error {
	return json.Unmarshal([]byte(value.(string)), j)
}

// MemberType is the type of member
type MemberType string

// AccessType is the access type of a member
type AccessType string

type Members []Member
type MembersWithUsersOrGroups []MemberWithUsersOrGroups

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

// MemberWithUser is a model for a member with user information
type MemberWithUsersOrGroups struct {
	Member
	User  User  `json:"user,omitempty"`
	Group Group `json:"group,omitempty"`
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
