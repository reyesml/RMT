package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

type JournalEntryRepo interface {
	Create(je *models.JournalEntry) error
	GetByUUID(uuid uuid.UUID) (models.JournalEntry, error)
	ListByUserId(uid uint) ([]models.JournalEntry, error)
}

func NewJournalEntryRepo(db *gorm.DB) journalEntryRepo {
	return journalEntryRepo{db: db}
}

type journalEntryRepo struct {
	db *gorm.DB
}

func (r journalEntryRepo) Create(je *models.JournalEntry) error {
	result := r.db.Create(je)
	return result.Error
}

func (r journalEntryRepo) GetByUUID(uuid uuid.UUID) (models.JournalEntry, error) {
	var je models.JournalEntry
	result := r.db.Where("UUID = ?", uuid).First(&je)
	return je, result.Error
}

func (r journalEntryRepo) ListByUserId(uid uint) ([]models.JournalEntry, error) {
	var jes []models.JournalEntry
	result := r.db.Where("user_id = ?", uid).Find(&jes)
	return jes, result.Error
}
