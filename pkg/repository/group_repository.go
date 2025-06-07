package repository

import (
	"github.com/labbs/mynotes/pkg/models"
	"gorm.io/gorm"
)

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *groupRepository {
	return &groupRepository{db: db}
}

// Create creates a group
func (r *groupRepository) Create(group models.Group) (models.Group, error) {
	return group, r.db.Create(&group).Error
}

// GetById returns a group by id
func (r *groupRepository) GetById(id string) (models.Group, error) {
	var group models.Group
	if err := r.db.Where("id = ?", id).First(&group).Error; err != nil {
		return models.Group{}, err
	}
	return group, nil
}

// GetGroupWithUsers returns a group with users
func (r *groupRepository) GetGroupWithUsers(id string) (models.Group, error) {
	var group models.Group
	if err := r.db.Preload("Users").Where("id = ?", id).First(&group).Error; err != nil {
		return models.Group{}, err
	}
	return group, nil
}

// GetAll returns all groups
func (r *groupRepository) GetAll() ([]models.Group, error) {
	var groups []models.Group
	if err := r.db.Find(&groups).Error; err != nil {
		return []models.Group{}, err
	}
	return groups, nil
}

// Update updates a group
func (r *groupRepository) Update(group models.Group) (models.Group, error) {
	return group, r.db.Save(&group).Error
}

// Delete deletes a group
func (r *groupRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Group{}).Error
}
