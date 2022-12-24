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

func TestCreateJournalEntry_Execute(t *testing.T) {
	testDBId := "TestCreatePerson.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	user, err := models.NewUser("test-user", "plain_password")
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	err = userRepo.Create(user)
	require.NoError(t, err)
	ctx := context.WithValue(context.Background(), models.UserCtxKey, *user)

	jr := repos.NewJournalEntryRepo(db)
	cje := NewCreateJournalEntry(jr)
	req := CreateJournalEntryRequest{
		Mood:  "productive",
		Title: "Writing Tests...",
		Body:  "Dear Diary,\nToday I wrote a test about journals.",
	}

	journal, err := cje.Execute(ctx, req)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, journal.UUID)
	require.Equal(t, user.SegmentUUID, journal.SegmentUUID)
	require.Equal(t, req.Mood, journal.Mood)
	require.Equal(t, req.Title, journal.Title)
	require.Equal(t, req.Body, journal.Body)
}
