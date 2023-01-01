package models

import (
	"github.com/reyesml/RMT/app/core/database"
	"gorm.io/gorm"
	"time"
)

type PersonQuality struct {
	database.WithUUID
	database.Segmented
	PersonId  uint `gorm:"primaryKey"`
	QualityId uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
	UserId    uint
	User      User
	Notes     []Note
}
