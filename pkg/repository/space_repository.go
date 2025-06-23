package repository

import (
	"fmt"

	"github.com/labbs/zotion/pkg/models"
	"gorm.io/gorm"
)

type spaceRepository struct {
	db *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) *spaceRepository {
	return &spaceRepository{db: db}
}

// Get Space for a user
func (r *spaceRepository) GetSpacesForUser(userId string, groups []models.Group) ([]models.Space, error) {
	var spaces []models.Space
	var query *gorm.DB
	if r.db.Dialector.Name() == "sqlite" {
		query = r.db.Table("space").Where("JSON_EXTRACT(members, '$') LIKE ?", fmt.Sprintf("%%\"id\":\"%s\"%%", userId))
		if len(groups) > 0 {
			for _, group := range groups {
				query = query.Or("JSON_EXTRACT(members, '$') LIKE ?", fmt.Sprintf("%%\"id\":\"%s\"%%", group.Id))
			}
		}
	} else {
		query = r.db.Table("space").Where("members @> ?", fmt.Sprintf(`[{"id": "%s", "type": "user"}]`, userId))
		if len(groups) > 0 {
			for _, group := range groups {
				query = query.Or("members @> ?", fmt.Sprintf(`[{"id": "%s", "type": "group"}]`, group.Id))
			}
		}
	}

	err := query.Find(&spaces).Error
	return spaces, err
}

// GetSpaceById returns a space by its id
func (r *spaceRepository) GetSpaceById(spaceId string) (models.Space, error) {
	var space models.Space
	err := r.db.Table("space").Preload("Documents").First(&space, "id = ?", spaceId).Error
	return space, err
}

// CreateSpace creates a new space
func (sr *spaceRepository) CreateSpace(space models.Space) (models.Space, error) {
	err := sr.db.Table("space").Create(&space).Error
	return space, err
}

// IsMember checks if a user is a member of a space
func (sr *spaceRepository) IsMember(spaceId, userId string) (bool, error) {
	var count int64

	// Get database dialect
	dialect := sr.db.Dialector.Name()

	if dialect == "sqlite" {
		// For SQLite: Use JSON_EXTRACT and LIKE for checking members
		err := sr.db.Table("space").
			Where("id = ? AND (JSON_EXTRACT(members, '$') LIKE ?)",
				spaceId, userId, "%\"id\":\""+userId+"\"%").
			Count(&count).Error
		return count > 0, err
	} else {
		// For PostgreSQL: Use the @> containment operator
		err := sr.db.Table("space").
			Where("id = ? AND members @> ?",
				spaceId, userId, fmt.Sprintf(`[{"id": "%s"}]`, userId)).
			Count(&count).Error
		return count > 0, err
	}
}
