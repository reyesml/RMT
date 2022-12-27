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

func TestGetJournal_Execute(t *testing.T) {
	testDBId := "TestGetJournal.db"
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
	ctx := context.WithValue(context.Background(), models.UserCtxKey, currUser)

	journalRepo := repos.NewJournalRepo(db)
	je := models.Journal{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		Mood:      "determined",
		Title:     "Dear diary...",
		Body:      "Today I worked on some tests for journal entries.",
		UserId:    user.ID,
	}
	require.NoError(t, journalRepo.Create(&je))

	gje := NewGetJournal(journalRepo)
	foundJournal, err := gje.Execute(ctx, GetJournalRequest{UUID: je.UUID})
	require.NoError(t, err)
	require.Equal(t, je.UUID, foundJournal.UUID)
	require.Equal(t, je.Mood, foundJournal.Mood)
	require.Equal(t, je.Title, foundJournal.Title)
	require.Equal(t, je.Body, foundJournal.Body)

	//attempt to fetch journal from different segment
	currUser.SegmentUUID, err = uuid.NewRandom()
	require.NoError(t, err)
	ctx = context.WithValue(context.Background(), models.UserCtxKey, currUser)
	foundJournal, err = gje.Execute(ctx, GetJournalRequest{UUID: je.UUID})
	require.Error(t, err)
	require.Equal(t, uuid.Nil, foundJournal.UUID)
}
