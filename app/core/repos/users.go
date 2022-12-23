package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
	"strings"
)

type UserRepo interface {
	// GetByUsername retrieves the user with the supplied username.
	// Returns an error if the user cannot be found.
	GetByUsername(uname string) (models.User, error)
	// FindByUsername retrieves the list of users whose username matches
	// the supplied username. An empty list is returned if no users are found.
	FindByUsername(uname string) ([]models.User, error)
	// GetByUUID retrieves the user with the supplied UUID. Returns an
	// error if the user cannot be found.
	GetByUUID(uuid uuid.UUID) (models.User, error)
	// Create inserts a new user into the database. Returns an error if
	// the insertion fails.
	Create(user *models.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userRepo {
	return userRepo{
		db: db,
	}
}

func (r userRepo) GetByUsername(uname string) (models.User, error) {
	var u models.User
	result := r.db.Where("lower(Username) = ?", strings.ToLower(uname)).First(&u)
	return u, result.Error
}

func (r userRepo) FindByUsername(uname string) ([]models.User, error) {
	var u []models.User
	result := r.db.Where("Username = ?", uname).Find(&u)
	return u, result.Error
}

func (r userRepo) GetByUUID(uuid uuid.UUID) (models.User, error) {
	var u models.User
	result := r.db.Where("UUID = ?", uuid).First(&u)
	return u, result.Error
}

func (r userRepo) Create(user *models.User) error {
	user.UsernameLower = strings.ToLower(user.Username)
	result := r.db.Create(user)
	return result.Error
}
