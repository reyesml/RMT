package repos

import (
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"gorm.io/gorm"
)

type PersonRepo interface {
	Create(person *models.Person) error
	CreateMany(people *[]models.Person) error
	GetByUUID(uuid uuid.UUID) (models.Person, error)
	ListBySegment(segment uuid.UUID) ([]models.Person, error)
}

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) personRepo {
	return personRepo{db: db}
}

func (r personRepo) Create(person *models.Person) error {
	result := r.db.Create(person)
	return result.Error
}

func (r personRepo) CreateMany(people *[]models.Person) error {
	result := r.db.Create(people)
	return result.Error
}

func (r personRepo) GetByUUID(uuid uuid.UUID) (models.Person, error) {
	var p models.Person
	result := r.db.Where("UUID = ?", uuid).First(&p)
	return p, result.Error
}

func (r personRepo) ListBySegment(segment uuid.UUID) ([]models.Person, error) {
	var ppl []models.Person
	result := r.db.Where("segment_uuid = ?", segment).Find(&ppl)
	return ppl, result.Error
}