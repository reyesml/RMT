package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type GetPersonQualityRequest struct {
	PersonQualityUUID uuid.UUID
}

type GetPersonQuality interface {
	Execute(ctx context.Context, req GetPersonQualityRequest) (models.PersonQuality, error)
}

func NewGetPersonQuality(pqr repos.PersonQualityRepo) getPersonQuality {
	return getPersonQuality{
		pqr: pqr,
	}
}

type getPersonQuality struct {
	pqr repos.PersonQualityRepo
}

func (ia getPersonQuality) Execute(ctx context.Context, req GetPersonQualityRequest) (models.PersonQuality, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return models.PersonQuality{}, database.SegmentMissingErr
	}

	pq, err := ia.pqr.GetByUUID(req.PersonQualityUUID, user.SegmentUUID)
	if err != nil {
		return models.PersonQuality{}, ErrNotFound
	}
	return pq, nil
}
