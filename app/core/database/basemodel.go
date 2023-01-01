package database

import (
	"gorm.io/gorm"
)

// BaseModel includes a default ID (int), an auto-populated UUID,
// and basic timestamps such as CreatedAt/UpdatedAt. It also enables
// soft-deletes by default.
type BaseModel struct {
	gorm.Model
	WithUUID
}
