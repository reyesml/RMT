package repos

import (
	"errors"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type JournalRepo interface {
	Create(je *models.Journal) error
	CreateMany(jes *[]models.Journal) error
	GetByUUIDWithUser(uuid uuid.UUID, segment uuid.UUID) (models.Journal, error)
	ListByUserIdWithUser(uid uint) ([]models.Journal, error)
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

func (r journalRepo) CreateMany(jes *[]models.Journal) error {
	result := r.db.Create(jes)
	return result.Error
}

func (r journalRepo) GetByUUIDWithUser(uuid uuid.UUID, segment uuid.UUID) (models.Journal, error) {
	var je models.Journal
	result := r.db.Preload("User").Where("UUID = ? AND SEGMENT_UUID = ?", uuid, segment).First(&je)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return je, ErrNotFound
	}
	return je, result.Error
}

func (r journalRepo) ListByUserIdWithUser(uid uint) ([]models.Journal, error) {
	var jes []models.Journal
	result := r.db.Preload("User").Where("user_id = ?", uid).Find(&jes)
	return jes, result.Error
}
