package utils

import (
	"github.com/reyesml/RMT/app/core/identity"
	"gorm.io/gorm"
)

var models = []interface{}{
	&identity.User{},
	&identity.Session{},
}

// MigrateAllModels runs all pending database
// migrations for all models (determined by the
// slice above). This is not intended to be used
// in production databases. Instead, refer to {tbd}.
func MigrateAllModels(db *gorm.DB) error {
	err := db.AutoMigrate(models...)
	return err
}
