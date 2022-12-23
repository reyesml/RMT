package utils

import (
	models2 "github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

var models = []interface{}{
	&models2.User{},
	&models2.Session{},
}

// MigrateAllModels runs all pending database
// migrations for all models (determined by the
// slice above). This is not intended to be used
// in production databases. Instead, refer to {tbd}.
func MigrateAllModels(db *gorm.DB) error {
	err := db.AutoMigrate(models...)
	return err
}
