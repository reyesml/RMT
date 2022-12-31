package interactors

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type GetPersonRequest struct {
	UUID uuid.UUID
}

type GetPerson interface {
	Execute(ctx context.Context, req GetPersonRequest) (models.Person, error)
}

func NewGetPerson(personRepo repos.PersonRepo) getPerson {
	return getPerson{personRepo: personRepo}
}

type getPerson struct {
	personRepo repos.PersonRepo
}

func (ia getPerson) Execute(ctx context.Context, req GetPersonRequest) (models.Person, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil {
		return models.Person{}, err
	}
	if user.SegmentUUID == uuid.Nil {
		return models.Person{}, database.SegmentMissingErr
	}

	p, err := ia.personRepo.GetByUUID(req.UUID, user.SegmentUUID)
	if errors.Is(err, repos.ErrNotFound) {
		return models.Person{}, ErrNotFound
	}
	if err != nil {
		return models.Person{}, fmt.Errorf("fetching person: %w", err)
	}
	return p, nil
}
