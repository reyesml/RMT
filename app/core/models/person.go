package models

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
)

type Person struct {
	database.BaseModel
	database.Segmented
	FirstName string
	LastName  string
}

func NewPerson(segment uuid.UUID, first string, last string) *Person {
	return &Person{
		Segmented: database.Segmented{SegmentUUID: segment},
		FirstName: first,
		LastName:  last,
	}
}
