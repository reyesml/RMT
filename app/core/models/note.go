package models

import "github.com/reyesml/RMT/app/core/database"

type Note struct {
	database.BaseModel
	database.Segmented
	Title           string
	Body            string
	UserId          uint
	User            User
	PersonId        uint
	Person          Person
	PersonQualityId uint
	PersonQuality   PersonQuality
}
