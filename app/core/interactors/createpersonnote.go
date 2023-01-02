package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type CreatePersonNoteRequest struct {
	PersonUUID uuid.UUID
	NoteTitle  string
	NoteBody   string
}

type CreatePersonNote interface {
	Execute(ctx context.Context, req CreatePersonNoteRequest) (models.Note, error)
}

func NewCreatePersonNote(pr repos.PersonRepo, nr repos.NoteRepo) createPersonNote {
	return createPersonNote{
		pr: pr,
		nr: nr,
	}
}

type createPersonNote struct {
	pr repos.PersonRepo
	nr repos.NoteRepo
}

func (ia createPersonNote) Execute(ctx context.Context, req CreatePersonNoteRequest) (models.Note, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return models.Note{}, database.SegmentMissingErr
	}

	if len(req.NoteTitle) == 0 {
		return models.Note{}, MissingNoteTitleErr
	}

	p, err := ia.pr.GetByUUID(req.PersonUUID, user.SegmentUUID)
	if err != nil {
		return models.Note{}, ErrNotFound
	}

	n := models.Note{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		Title:     req.NoteTitle,
		Body:      req.NoteBody,
		UserId:    user.ID,
		PersonId:  p.ID,
	}
	if err := ia.nr.Create(&n); err != nil {
		return models.Note{}, err
	}
	return n, nil
}
