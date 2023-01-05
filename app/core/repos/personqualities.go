package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

type PersonQualityRepo interface {
	Create(pq *models.PersonQuality) error
	CreateMany(pqs *[]models.PersonQuality) error
	GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.PersonQuality, error)
	ListByPerson(pid uint) ([]models.PersonQuality, error)
}

type personQualityRepo struct {
	db *gorm.DB
}

func NewPersonQualityRepo(db *gorm.DB) personQualityRepo {
	return personQualityRepo{db: db}
}

func (r personQualityRepo) Create(pq *models.PersonQuality) error {
	result := r.db.Create(pq)
	return result.Error
}

func (r personQualityRepo) CreateMany(pqs *[]models.PersonQuality) error {
	result := r.db.Create(pqs)
	return result.Error
}

func (r personQualityRepo) GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.PersonQuality, error) {
	var pq models.PersonQuality
	result := r.db.Preload("Quality").Preload("Person").Preload("Notes").Where("UUID = ? AND SEGMENT_UUID = ?", uuid, segment).First(&pq)
	return pq, result.Error
}

func (r personQualityRepo) ListByPerson(pid uint) ([]models.PersonQuality, error) {
	var pqs []models.PersonQuality
	result := r.db.Order("created_at desc").Preload("Quality").Preload("Person").Preload("Notes").Where("person_id = ?", pid).Find(&pqs)
	return pqs, result.Error
}
