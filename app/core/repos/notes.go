package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

type NoteRepo interface {
	Create(n *models.Note) error
	CreateMany(ns *[]models.Note) error
	GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.Note, error)
	ListByPersonId(pid uint) ([]models.Note, error)
	ListByPersonQualityId(pqid uint) ([]models.Note, error)
}

func NewNoteRepo(db *gorm.DB) noteRepo {
	return noteRepo{db: db}
}

type noteRepo struct {
	db *gorm.DB
}

func (r noteRepo) Create(n *models.Note) error {
	result := r.db.Create(n)
	return result.Error
}

func (r noteRepo) CreateMany(ns *[]models.Note) error {
	result := r.db.Create(ns)
	return result.Error
}

func (r noteRepo) GetByUUID(uuid uuid.UUID, segment uuid.UUID) (models.Note, error) {
	var n models.Note
	result := r.db.Where("uuid = ? and segment_uuid = ?", uuid, segment).First(&n)
	return n, result.Error
}

func (r noteRepo) ListByPersonId(pid uint) ([]models.Note, error) {
	var ns []models.Note
	result := r.db.Where("person_id = ?", pid).Find(&ns)
	return ns, result.Error
}

func (r noteRepo) ListByPersonQualityId(pqid uint) ([]models.Note, error) {
	var ns []models.Note
	result := r.db.Where("person_quality_id = ?", pqid).Find(&ns)
	return ns, result.Error
}
