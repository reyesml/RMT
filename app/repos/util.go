package repos

import (
	"github.com/reyesml/RMT/app/core/identity"
	"gorm.io/gorm"
)

var models = []interface{}{
	&identity.User{},
}

func MigrateAll(db *gorm.DB) error {
	err := db.AutoMigrate(models...)
	return err
}
