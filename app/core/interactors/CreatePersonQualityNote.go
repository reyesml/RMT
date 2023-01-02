package interactors

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type CreatePersonQualityNoteRequest struct {
	PersonQualityUUID uuid.UUID
	NoteTitle         string
	NoteBody          string
}

type CreatePersonQualityNote interface {
	Execute(ctx context.Context, req CreatePersonQualityNoteRequest) (models.Note, error)
}

func NewCreatePersonQualityNote(pqr repos.PersonQualityRepo, nr repos.NoteRepo) createPersonQualityNote {
	return createPersonQualityNote{
		pqr: pqr,
		nr:  nr,
	}
}

type createPersonQualityNote struct {
	pqr repos.PersonQualityRepo
	nr  repos.NoteRepo
}

var MissingNoteTitleErr = errors.New("note title is required")

func (ia createPersonQualityNote) Execute(ctx context.Context, req CreatePersonQualityNoteRequest) (models.Note, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return models.Note{}, database.SegmentMissingErr
	}

	if len(req.NoteTitle) == 0 {
		return models.Note{}, MissingNoteTitleErr
	}

	pq, err := ia.pqr.GetByUUID(req.PersonQualityUUID, user.SegmentUUID)
	if err != nil {
		return models.Note{}, ErrNotFound
	}

	n := models.Note{
		Segmented:       database.Segmented{SegmentUUID: user.SegmentUUID},
		Title:           req.NoteTitle,
		Body:            req.NoteBody,
		UserId:          user.ID,
		PersonId:        pq.PersonId,
		PersonQualityId: pq.ID,
	}
	if err := ia.nr.Create(&n); err != nil {
		return models.Note{}, err
	}
	return n, nil
}
