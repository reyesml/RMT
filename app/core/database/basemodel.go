package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	UUID uuid.UUID `gorm:"uniqueIndex;type:uuid"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	return base.AddUUID()
}

func (base *BaseModel) AddUUID() error {
	//Don't overwrite a UUID if one is already provided
	if base.UUID != uuid.Nil {
		return nil
	}
	//Add the UUID if one doesn't exist
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	base.UUID = uuid
	return nil
}
