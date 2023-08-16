package migration

import (
	"unjukketerampilan/models"

	"gorm.io/gorm"
)

func RunMigrationsForMerks(db *gorm.DB) {
	// Auto Migrate untuk membuat tabel
	db.AutoMigrate(&models.Merk{})
}
