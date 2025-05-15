package models

import "time"

type UserGroup struct {
	UserId  string `gorm:"user_id" json:"user_id"`
	GroupId string `gorm:"group_id" json:"group_id"`

	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
}

// TableName returns the name of the table in the database.
// This is used by GORM to map the struct to the correct table.
func (ug *UserGroup) TableName() string {
	return "user_group"
}

// UserGroupRepository defines the methods that a user group repository should implement.
type UserGroupRepository interface{}

// UserGroupService defines the methods that a user group service should implement.
type UserGroupService interface{}
