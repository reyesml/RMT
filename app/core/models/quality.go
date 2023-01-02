package models

import "github.com/reyesml/RMT/app/core/database"

type Quality struct {
	database.BaseModel
	database.Segmented
	Name      string
	NameLower string
	UserId    uint
	User      User
}
