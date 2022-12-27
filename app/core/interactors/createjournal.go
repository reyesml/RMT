package interactors

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
)

type CreateJournalRequest struct {
	Mood  string
	Title string
	Body  string
}

var MissingJournalTitleErr = errors.New("missing journal title")
var MissingJournalBodyErr = errors.New("missing journal body")

type CreateJournal interface {
	Execute(ctx context.Context, req CreateJournalRequest) (models.Journal, error)
}

func NewCreateJournal(journalRepo repos.JournalRepo) createJournal {
	return createJournal{journalRepo: journalRepo}
}

type createJournal struct {
	journalRepo repos.JournalRepo
}

func (ia createJournal) Execute(ctx context.Context, req CreateJournalRequest) (models.Journal, error) {
	user, ok := ctx.Value(models.UserCtxKey).(models.CurrentUser)
	if !ok {
		return models.Journal{}, models.UserMissingErr
	}
	if user.SegmentUUID == uuid.Nil {
		return models.Journal{}, database.SegmentMissingErr
	}
	if len(req.Title) == 0 {
		return models.Journal{}, MissingJournalTitleErr
	}
	if len(req.Body) == 0 {
		return models.Journal{}, MissingJournalBodyErr
	}

	je := models.Journal{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		Mood:      req.Mood,
		Title:     req.Title,
		Body:      req.Body,
		UserId:    user.ID,
	}

	err := ia.journalRepo.Create(&je)
	if err != nil {
		return models.Journal{}, err
	}
	return je, nil
}
