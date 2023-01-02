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

func TestAddNoteToPerson_Execute(t *testing.T) {
	testDBId := "TestAddNoteToPerson.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	user, err := models.NewUser("test-user", "plain_password")
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	require.NoError(t, userRepo.Create(user))

	pr := repos.NewPersonRepo(db)
	person := models.Person{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		FirstName: "sam",
		LastName:  "doe",
		UserId:    user.ID,
	}
	require.NoError(t, pr.Create(&person))

	nr := repos.NewNoteRepo(db)
	antp := NewCreatePersonNote(pr, nr)
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	antpr := CreatePersonNoteRequest{
		PersonUUID: person.UUID,
		NoteTitle:  "They spent most of their holiday working on a CRM",
		NoteBody:   "IDK what to put as the body.",
	}
	n, err := antp.Execute(ctx, antpr)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, n.UUID)
	require.Equal(t, person.ID, n.PersonId)
	require.Equal(t, uint(0), n.PersonQualityId)
	require.Equal(t, user.SegmentUUID, n.SegmentUUID)
	require.Equal(t, antpr.NoteTitle, n.Title)
	require.Equal(t, antpr.NoteBody, n.Body)

	n2, err := antp.Execute(ctx, antpr)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, n2.UUID)
	require.NotEqual(t, n.ID, n2.ID)
}
