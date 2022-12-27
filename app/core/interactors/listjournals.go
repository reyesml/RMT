package interactors

import (
	"context"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
)

type ListJournals interface {
	Execute(ctx context.Context) ([]models.Journal, error)
}

func NewListJournals(journalRepo repos.JournalRepo) listJournals {
	return listJournals{journalRepo: journalRepo}
}

type listJournals struct {
	journalRepo repos.JournalRepo
}

func (ia listJournals) Execute(ctx context.Context) ([]models.Journal, error) {
	user, ok := ctx.Value(models.UserCtxKey).(models.CurrentUser)
	if !ok || user.ID == 0 {
		return []models.Journal{}, models.UserMissingErr
	}

	return ia.journalRepo.ListByUserIdWithUser(user.ID)
}
