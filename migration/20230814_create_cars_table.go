package migration

import (
	"unjukketerampilan/models"

	"gorm.io/gorm"
)

func RunMigrationsForCars(db *gorm.DB) {
	// Auto Migrate untuk membuat tabel
	db.AutoMigrate(&models.Car{})
}
