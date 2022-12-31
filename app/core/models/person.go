package models

import (
	"github.com/reyesml/RMT/app/core/database"
)

type Person struct {
	database.BaseModel
	database.Segmented
	FirstName string
	LastName  string
}
