package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/identity"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetByUsername(uname string) (identity.User, error)
	FindByUsername(uname string) ([]identity.User, error)
	GetByUUID(uuid uuid.UUID) (identity.User, error)
	Create(user *identity.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{
		db: db,
	}
}

// GetByUsername retrieves the user with the supplied username.
// Returns an error if the user cannot be found.
func (r userRepo) GetByUsername(uname string) (identity.User, error) {
	var u identity.User
	result := r.db.Where("Username = ?", uname).First(&u)
	if result.Error != nil {
		return u, result.Error
	}
	return u, nil
}

// FindByUsername retrieves the list of users whose username matches
// the supplied username. An empty list is returned if no users are found.
func (r userRepo) FindByUsername(uname string) ([]identity.User, error) {
	var u []identity.User
	result := r.db.Where("Username = ?", uname).Find(&u)
	if result.Error != nil {
		return u, result.Error
	}
	return u, nil
}

// GetByUUID retrieves the user with the supplied UUID. Returns an
// error if the user cannot be found.
func (r userRepo) GetByUUID(uuid uuid.UUID) (identity.User, error) {
	var u identity.User
	result := r.db.Where("UUID = ?", uuid).First(&u)
	if result.Error != nil {
		return u, result.Error
	}
	return u, nil
}

func (r userRepo) Create(user *identity.User) error {
	result := r.db.Create(user)
	return result.Error
}
