package interactors

import (
	"context"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
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
	user, err := utils.GetCurrentUser(ctx)
	if err != nil {
		return []models.Journal{}, err
	}

	return ia.journalRepo.ListByUserIdWithUser(user.ID)
}
