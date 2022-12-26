package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

type JournalRepo interface {
	Create(je *models.Journal) error
	GetByUUID(uuid uuid.UUID) (models.Journal, error)
	ListByUserId(uid uint) ([]models.Journal, error)
}

func NewJournalRepo(db *gorm.DB) journalRepo {
	return journalRepo{db: db}
}

type journalRepo struct {
	db *gorm.DB
}

func (r journalRepo) Create(je *models.Journal) error {
	result := r.db.Create(je)
	return result.Error
}

func (r journalRepo) GetByUUID(uuid uuid.UUID) (models.Journal, error) {
	var je models.Journal
	result := r.db.Where("UUID = ?", uuid).First(&je)
	return je, result.Error
}

func (r journalRepo) ListByUserId(uid uint) ([]models.Journal, error) {
	var jes []models.Journal
	result := r.db.Where("user_id = ?", uid).Find(&jes)
	return jes, result.Error
}
