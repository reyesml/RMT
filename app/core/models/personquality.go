package models

import (
	"github.com/reyesml/RMT/app/core/database"
)

type PersonQuality struct {
	database.BaseModel
	database.Segmented
	PersonId  uint
	QualityId uint
	Quality   Quality
	UserId    uint
	User      User
	Notes     []Note
}
