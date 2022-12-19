package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(dbid string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbid), &gorm.Config{})
	return db, err
}
