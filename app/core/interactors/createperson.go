package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type CreatePersonRequest struct {
	FirstName string
	LastName  string
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
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return models.Person{}, database.SegmentMissingErr
	}
	person := models.Person{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserId:    user.ID,
	}
	//person := models.NewPerson(user.SegmentUUID, req.FirstName, req.LastName)
	if err := ia.personRepo.Create(&person); err != nil {
		return models.Person{}, fmt.Errorf("creating person: %w", err)
	}
	return person, nil
}
