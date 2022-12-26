package models

import "github.com/reyesml/RMT/app/core/database"

type Journal struct {
	database.BaseModel
	database.Segmented
	Mood   string
	Title  string
	Body   string
	UserId uint
	User   User
}
