package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
	"strings"
)

type QualityRepo interface {
	Create(quality *models.Quality) error
	CreateMany(qualities *[]models.Quality) error
	GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.Quality, error)
	GetByID(id uint, segment uuid.UUID) (models.Quality, error)
	ListBySegment(segment uuid.UUID) ([]models.Quality, error)
	FindByNameAndType(name string, t string, userId uint, segment uuid.UUID) ([]models.Quality, error)
}

type qualityRepo struct {
	db *gorm.DB
}

func NewQualityRepo(db *gorm.DB) qualityRepo {
	return qualityRepo{db: db}
}

func (r qualityRepo) Create(quality *models.Quality) error {
	quality.NameLower = strings.ToLower(quality.Name)
	quality.TypeLower = strings.ToLower(quality.Type)
	result := r.db.Create(quality)
	return result.Error
}

func (r qualityRepo) CreateMany(qualities *[]models.Quality) error {
	// TODO: test this to see if it actually works.
	for _, q := range *qualities {
		q.NameLower = strings.ToLower(q.Name)
	}
	result := r.db.Create(qualities)
	return result.Error
}

func (r qualityRepo) GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.Quality, error) {
	var q models.Quality
	result := r.db.Where("UUID = ? and SEGMENT_UUID = ?", uuid, segment).First(&q)
	return q, result.Error
}

func (r qualityRepo) GetByID(id uint, segment uuid.UUID) (models.Quality, error) {
	var q models.Quality
	result := r.db.Where("id = ? and SEGMENT_UUID = ?", id, segment).First(&q)
	return q, result.Error
}

func (r qualityRepo) ListBySegment(segment uuid.UUID) ([]models.Quality, error) {
	var qs []models.Quality
	result := r.db.Where("segment_uuid = ? ", segment).Find(&qs)
	return qs, result.Error
}

func (r qualityRepo) FindByNameAndType(name string, t string, userId uint, segment uuid.UUID) ([]models.Quality, error) {
	var qs []models.Quality
	result := r.db.Where("name_lower = ? and type_lower = ? and user_id = ? and segment_uuid = ?",
		strings.ToLower(name),
		strings.ToLower(t),
		userId,
		segment).Find(&qs)
	return qs, result.Error
}
