package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	UUID uuid.UUID `gorm:"uniqueIndex;type:uuid"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	base.UUID = uuid
	return nil
}
