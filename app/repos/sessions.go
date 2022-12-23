package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/identity"
	"gorm.io/gorm"
	"time"
)

type SessionRepo interface {
	Create(session *identity.Session) error
	GetByUUIDWithUser(token uuid.UUID) (identity.Session, error)
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) sessionRepo {
	return sessionRepo{db: db}
}

func (r sessionRepo) Create(session *identity.Session) error {
	result := r.db.Create(session)
	return result.Error
}

func (r sessionRepo) GetByUUIDWithUser(uuid uuid.UUID) (identity.Session, error) {
	var s identity.Session
	result := r.db.Preload("User").Where("UUID = ? AND Expiration > ?", uuid, time.Now().UTC()).First(&s)
	return s, result.Error
}
