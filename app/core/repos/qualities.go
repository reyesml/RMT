package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

type QualityRepo interface {
	Create(quality *models.Quality) error
	CreateMany(qualities *[]models.Quality) error
	GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.Quality, error)
	ListBySegment(segment uuid.UUID) ([]models.Quality, error)
}

type qualityRepo struct {
	db *gorm.DB
}

func NewQualityRepo(db *gorm.DB) qualityRepo {
	return qualityRepo{db: db}
}

func (r qualityRepo) Create(quality *models.Quality) error {
	result := r.db.Create(quality)
	return result.Error
}

func (r qualityRepo) CreateMany(qualities *[]models.Quality) error {
	result := r.db.Create(qualities)
	return result.Error
}

func (r qualityRepo) GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.Quality, error) {
	var q models.Quality
	result := r.db.Where("UUID = ? and SEGMENT_UUID = ?", uuid, segment).First(&q)
	return q, result.Error
}

func (r qualityRepo) ListBySegment(segment uuid.UUID) ([]models.Quality, error) {
	var qs []models.Quality
	result := r.db.Where("segment_uuid = ? ", segment).Find(&qs)
	return qs, result.Error
}
