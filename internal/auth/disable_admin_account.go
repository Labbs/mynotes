package auth

import (
	"github.com/labbs/zotion/pkg/config"
	"github.com/labbs/zotion/pkg/models"
	"gorm.io/gorm"
)

func DisableAdminAccount(db *gorm.DB) error {
	if config.Auth.DisableAdminAccount {
		return db.Model(&models.User{}).Where("email = ?", "admin@zotion.local").Update("active", false).Error
	} else {
		return db.Model(&models.User{}).Where("email = ?", "admin@zotion.local").Update("active", true).Error
	}
}
