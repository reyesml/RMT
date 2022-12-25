package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
)

type CreateJournalEntryRequest struct {
	Mood  string
	Title string
	Body  string
}

type CreateJournalEntry interface {
	Execute(ctx context.Context, req CreateJournalEntryRequest)
}

func NewCreateJournalEntry(journalRepo repos.JournalEntryRepo) createJournalEntry {
	return createJournalEntry{journalRepo: journalRepo}
}

type createJournalEntry struct {
	journalRepo repos.JournalEntryRepo
}

func (ia createJournalEntry) Execute(ctx context.Context, req CreateJournalEntryRequest) (models.JournalEntry, error) {
	user, ok := ctx.Value(models.UserCtxKey).(models.CurrentUser)
	if !ok {
		return models.JournalEntry{}, models.UserMissingErr
	}
	if user.SegmentUUID == uuid.Nil {
		return models.JournalEntry{}, database.SegmentMissingErr
	}

	je := models.JournalEntry{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		Mood:      req.Mood,
		Title:     req.Title,
		Body:      req.Body,
		UserId:    user.ID,
	}

	err := ia.journalRepo.Create(&je)
	if err != nil {
		return models.JournalEntry{}, err
	}
	return je, nil
}
