package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
)

type ListPeopleRequest struct {
	SegmentUUID uuid.UUID
}

type ListPeople interface {
	Execute(ctx context.Context, req ListPeopleRequest) ([]models.Person, error)
}

func NewListPeople(personRepo repos.PersonRepo) listPeople {
	return listPeople{
		personRepo: personRepo,
	}
}

type listPeople struct {
	personRepo repos.PersonRepo
}

func (ia listPeople) Execute(ctx context.Context, req ListPeopleRequest) ([]models.Person, error) {
	_ = ctx
	return ia.personRepo.ListBySegment(req.SegmentUUID)
}
