package migration

import (
	"unjukketerampilan/models"

	"gorm.io/gorm"
)

func RunMigrationsForUsers(db *gorm.DB) {
	// Auto Migrate untuk membuat tabel
	db.AutoMigrate(&models.User{})
}
