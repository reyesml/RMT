package utils

import (
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

var AllModels = []interface{}{
	&models.User{},
	&models.Session{},
	&models.Person{},
	&models.JournalEntry{},
}

// MigrateAllModels runs all pending database
// migrations for all models (determined by the
// slice above). This is not intended to be used
// in production databases. Instead, refer to {tbd}.
func MigrateAllModels(db *gorm.DB) error {
	err := db.AutoMigrate(AllModels...)
	return err
}
