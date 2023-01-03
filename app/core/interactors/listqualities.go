package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type ListQualities interface {
	Execute(ctx context.Context) ([]models.Quality, error)
}

func NewListQualities(qr repos.QualityRepo) listQualities {
	return listQualities{
		qr: qr,
	}
}

type listQualities struct {
	qr repos.QualityRepo
}

func (ia listQualities) Execute(ctx context.Context) ([]models.Quality, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return []models.Quality{}, database.SegmentMissingErr
	}

	return ia.qr.ListBySegment(user.SegmentUUID)
}
