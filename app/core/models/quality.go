package models

import "github.com/reyesml/RMT/app/core/database"

type Quality struct {
	database.BaseModel
	database.Segmented
	Name      string
	Type      string
	TypeLower string
	NameLower string
	UserId    uint
	User      User
}
