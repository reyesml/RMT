package repos

import (
	"github.com/reyesml/RMT/app/core/identity"
	"gorm.io/gorm"
	"time"
)

type SessionRepo interface {
	Create(session *identity.Session) error
	GetByToken(token string) (identity.Session, error)
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) SessionRepo {
	return sessionRepo{db: db}
}

func (r sessionRepo) Create(session *identity.Session) error {
	result := r.db.Create(session)
	return result.Error
}

func (r sessionRepo) GetByToken(token string) (identity.Session, error) {
	var s identity.Session
	result := r.db.Where("Token = ? AND Expiration > ?", token, time.Now().UTC()).First(&s)
	return s, result.Error
}
