package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type ListPersonQualitiesRequest struct {
	PersonUUID uuid.UUID
}

type ListPersonQualities interface {
	Execute(ctx context.Context, req ListPersonQualitiesRequest) ([]models.PersonQuality, error)
}

func NewListPersonQualities(pr repos.PersonRepo, pqr repos.PersonQualityRepo) listPersonQualities {
	return listPersonQualities{
		pr:  pr,
		pqr: pqr,
	}
}

type listPersonQualities struct {
	pr  repos.PersonRepo
	pqr repos.PersonQualityRepo
}

func (ia listPersonQualities) Execute(ctx context.Context, req ListPersonQualitiesRequest) ([]models.PersonQuality, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return []models.PersonQuality{}, database.SegmentMissingErr
	}

	p, err := ia.pr.GetByUUID(req.PersonUUID, user.SegmentUUID)
	if err != nil {
		return []models.PersonQuality{}, ErrNotFound
	}

	// Find our person
	pqs, err := ia.pqr.ListByPerson(p.ID)
	if err != nil {
		return []models.PersonQuality{}, ErrNotFound
	}
	return pqs, nil
}
