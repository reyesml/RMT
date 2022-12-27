package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
)

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

	je, err := ia.journalRepo.GetByUUID(req.UUID, user.SegmentUUID)
	if err != nil {
		return models.Journal{}, fmt.Errorf("fetching journal: %w", err)
	}
	return je, nil
}
