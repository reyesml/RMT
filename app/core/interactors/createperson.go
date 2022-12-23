package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
)

type CreatePersonRequest struct {
	FirstName   string
	LastName    string
	SegmentUUID uuid.UUID
}

type CreatePerson interface {
	Execute(ctx context.Context, req CreatePersonRequest) (models.Person, error)
}

func NewCreatePerson(personRepo repos.PersonRepo) createPerson {
	return createPerson{
		personRepo: personRepo,
	}
}

type createPerson struct {
	personRepo repos.PersonRepo
}

func (ia createPerson) Execute(ctx context.Context, req CreatePersonRequest) (models.Person, error) {
	_ = ctx
	person := models.NewPerson(req.SegmentUUID, req.FirstName, req.LastName)
	if err := ia.personRepo.Create(person); err != nil {
		return models.Person{}, fmt.Errorf("creating person: %w", err)
	}
	return *person, nil
}
