package interactors

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
)

var ErrNotFound = errors.New("not found")

type GetJournalRequest struct {
	UUID uuid.UUID
}

type GetJournal interface {
	Execute(ctx context.Context, req GetJournalRequest) (models.Journal, error)
}

func NewGetJournal(journalRepo repos.JournalRepo) getJournal {
	return getJournal{
		journalRepo: journalRepo,
	}
}

type getJournal struct {
	journalRepo repos.JournalRepo
}

func (ia getJournal) Execute(ctx context.Context, req GetJournalRequest) (models.Journal, error) {
	user, ok := ctx.Value(models.UserCtxKey).(models.CurrentUser)
	if !ok {
		return models.Journal{}, models.UserMissingErr
	}
	if user.SegmentUUID == uuid.Nil {
		return models.Journal{}, database.SegmentMissingErr
	}

	je, err := ia.journalRepo.GetByUUIDWithUser(req.UUID, user.SegmentUUID)
	if errors.Is(err, repos.ErrNotFound) {
		return models.Journal{}, ErrNotFound
	}
	if err != nil {
		return models.Journal{}, fmt.Errorf("fetching journal: %w", err)
	}
	return je, nil
}
