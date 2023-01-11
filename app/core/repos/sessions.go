package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
	"time"
)

type SessionRepo interface {
	Create(session *models.Session) error
	GetByUUID(uuid uuid.UUID) (models.Session, error)
	Delete(session *models.Session) error
	GetByUUIDWithUser(token uuid.UUID) (models.Session, error)
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) sessionRepo {
	return sessionRepo{db: db}
}

func (r sessionRepo) Create(session *models.Session) error {
	result := r.db.Create(session)
	return result.Error
}

func (r sessionRepo) Delete(session *models.Session) error {
	result := r.db.Delete(session)
	return result.Error
}

func (r sessionRepo) GetByUUID(uuid uuid.UUID) (models.Session, error) {
	var s models.Session
	result := r.db.Where("UUID = ? AND Expiration > ?", uuid, time.Now().UTC()).First(&s)
	return s, result.Error
}

func (r sessionRepo) GetByUUIDWithUser(uuid uuid.UUID) (models.Session, error) {
	var s models.Session
	result := r.db.Preload("User").Where("UUID = ? AND Expiration > ?", uuid, time.Now().UTC()).First(&s)
	return s, result.Error
}
