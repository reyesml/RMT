package interactors

import (
	"context"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type ListPeople interface {
	Execute(ctx context.Context) ([]models.Person, error)
}

func NewListPeople(personRepo repos.PersonRepo) listPeople {
	return listPeople{
		personRepo: personRepo,
	}
}

type listPeople struct {
	personRepo repos.PersonRepo
}

func (ia listPeople) Execute(ctx context.Context) ([]models.Person, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil {
		return []models.Person{}, err
	}
	return ia.personRepo.ListBySegment(user.SegmentUUID)
}
