package models

import (
	"github.com/google/uuid"
)

// CurrentUser adds a SessionUUID to the user object. It is used
// for request contexts
type CurrentUser struct {
	User
	SessionUUID uuid.UUID
}
