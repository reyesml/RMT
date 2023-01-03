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
	// ListAllByPersonId returns notes associated directly to the person
	// and notes associated to qualities of this person.
	ListAllByPersonId(pid uint) ([]models.Note, error)
	// ListByPersonId returns only those notes directly associated with
	// the person. It excludes notes associated to person qualities.
	ListByPersonId(pid uint) ([]models.Note, error)
	// ListByPersonQualityId returns only notes associated with the
	// provided PersonQuality id.
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

func (r noteRepo) ListAllByPersonId(pid uint) ([]models.Note, error) {
	var ns []models.Note
	result := r.db.Where("person_id = ?", pid).Find(&ns)
	return ns, result.Error
}

func (r noteRepo) ListByPersonId(pid uint) ([]models.Note, error) {
	var ns []models.Note
	result := r.db.Preload("PersonQuality").Where("person_id = ? and person_quality_id = 0", pid).Find(&ns)
	return ns, result.Error
}

func (r noteRepo) ListByPersonQualityId(pqid uint) ([]models.Note, error) {
	var ns []models.Note
	result := r.db.Preload("PersonQuality").Where("person_quality_id = ?", pqid).Find(&ns)
	return ns, result.Error
}
