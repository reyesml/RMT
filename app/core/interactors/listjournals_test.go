package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestListJournals_Execute(t *testing.T) {
	testDBId := "TestListJournal.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	user, err := models.NewUser("test-user", "plain_password")
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	require.NoError(t, userRepo.Create(user))
	currUser := models.CurrentUser{
		User:        *user,
		SessionUUID: uuid.Nil,
	}
	user1Journals := []models.Journal{
		{
			Segmented: database.Segmented{SegmentUUID: currUser.SegmentUUID},
			Mood:      "productive",
			Title:     "more tests!!",
			Body:      "OMG I love writing tests.",
			UserId:    currUser.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: currUser.SegmentUUID},
			Mood:      "productive",
			Title:     "wow, tests!",
			Body:      "So many tests (:",
			UserId:    currUser.ID,
		},
	}
	journalRepo := repos.NewJournalRepo(db)
	require.NoError(t, journalRepo.CreateMany(&user1Journals))

	otherUser, err := models.NewUser("someone-else", "random-password")
	require.NoError(t, err)
	require.NoError(t, userRepo.Create(otherUser))
	user2Journals := []models.Journal{
		{
			Segmented: database.Segmented{SegmentUUID: otherUser.SegmentUUID},
			Mood:      "confused",
			Title:     "why so many tests??",
			Body:      "Is this really necessary?",
			UserId:    otherUser.ID,
		},
	}
	require.NoError(t, journalRepo.CreateMany(&user2Journals))

	lje := NewListJournals(journalRepo)
	ctx := context.WithValue(context.Background(), models.UserCtxKey, currUser)
	result, err := lje.Execute(ctx)
	require.Equal(t, len(user1Journals), len(result))
	for i, je := range result {
		require.Equal(t, user1Journals[i].UUID, je.UUID)
		require.Equal(t, currUser.SegmentUUID, je.SegmentUUID)
		require.NotEqual(t, uuid.Nil, je.User.UUID)
	}
}
