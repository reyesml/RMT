package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type ListPersonQualityNotesRequest struct {
	PersonQualityUUID uuid.UUID
}

type ListPersonQualityNotes interface {
	Execute(ctx context.Context, req ListPersonQualityNotesRequest) ([]models.Note, error)
}

func NewListPersonQualityNotes(pqr repos.PersonQualityRepo, nr repos.NoteRepo) listPersonQualityNotes {
	return listPersonQualityNotes{
		pqr: pqr,
		nr:  nr,
	}
}

type listPersonQualityNotes struct {
	pqr repos.PersonQualityRepo
	nr  repos.NoteRepo
}

func (ia listPersonQualityNotes) Execute(ctx context.Context, req ListPersonQualityNotesRequest) ([]models.Note, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return []models.Note{}, database.SegmentMissingErr
	}

	pq, err := ia.pqr.GetByUUID(req.PersonQualityUUID, user.SegmentUUID)
	if err != nil {
		return []models.Note{}, ErrNotFound
	}

	ns, err := ia.nr.ListByPersonQualityId(pq.ID)
	if err != nil {
		return []models.Note{}, err
	}
	return ns, err
}
