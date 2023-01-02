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

type AddNoteToPersonQualityRequest struct {
	PersonQualityUUID uuid.UUID
	NoteTitle         string
	NoteBody          string
}

type AddNoteToPersonQuality interface {
	Execute(ctx context.Context, req AddNoteToPersonQuality) (models.Note, error)
}

func NewAddNoteToPersonQuality(pqr repos.PersonQualityRepo, nr repos.NoteRepo) addNoteToPersonQuality {
	return addNoteToPersonQuality{
		pqr: pqr,
		nr:  nr,
	}
}

type addNoteToPersonQuality struct {
	pqr repos.PersonQualityRepo
	nr  repos.NoteRepo
}

var MissingNoteTitleErr = errors.New("note title is required")

func (ia addNoteToPersonQuality) Execute(ctx context.Context, req AddNoteToPersonQualityRequest) (models.Note, error) {
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
