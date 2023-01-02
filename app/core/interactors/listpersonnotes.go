package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type ListPersonNotesRequest struct {
	PersonUUID uuid.UUID
	Filter     PersonNoteFilter
}

type PersonNoteFilter uint16

const (
	AllNotesFilter PersonNoteFilter = iota
	OnlyPersonNotesFilter
)

type ListPersonNotes interface {
	Execute(ctx context.Context, req ListPersonNotesRequest) ([]models.Note, error)
}

func NewListPersonNotes(pr repos.PersonRepo, nr repos.NoteRepo) listPersonNotes {
	return listPersonNotes{
		pr: pr,
		nr: nr,
	}
}

type listPersonNotes struct {
	pr repos.PersonRepo
	nr repos.NoteRepo
}

func (ia listPersonNotes) Execute(ctx context.Context, req ListPersonNotesRequest) ([]models.Note, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return []models.Note{}, database.SegmentMissingErr
	}

	p, err := ia.pr.GetByUUID(req.PersonUUID, user.SegmentUUID)
	if err != nil {
		return []models.Note{}, ErrNotFound
	}

	var ns []models.Note
	switch req.Filter {
	case AllNotesFilter:
		ns, err = ia.nr.ListAllByPersonId(p.ID)
	case OnlyPersonNotesFilter:
		ns, err = ia.nr.ListByPersonId(p.ID)
	}
	if err != nil {
		return []models.Note{}, err
	}
	return ns, nil
}
