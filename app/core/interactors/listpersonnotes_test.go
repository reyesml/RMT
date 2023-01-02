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

func TestListPersonNotes_Execute(t *testing.T) {
	testDBId := "TestListPersonNotes.db"
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

	qr := repos.NewQualityRepo(db)
	pqr := repos.NewPersonQualityRepo(db)
	aqtp := NewAddQualityToPerson(pr, qr, pqr)
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	cpqr := AddQualityToPersonRequest{
		PersonUUID:  person.UUID,
		QualityName: "Loves Writing Tests",
	}
	pq, err := aqtp.Execute(ctx, cpqr)
	require.NoError(t, err)

	antpqr := AddNoteToPersonQualityRequest{
		PersonQualityUUID: pq.UUID,
		NoteTitle:         "They insist on testing every interactor",
		NoteBody:          "...but mostly only the happy paths (:",
	}
	nr := repos.NewNoteRepo(db)
	antpq := NewAddNoteToPersonQuality(pqr, nr)
	_, err = antpq.Execute(ctx, antpqr)
	require.NoError(t, err)

	antpr := AddNoteToPersonRequest{
		PersonUUID: person.UUID,
		NoteTitle:  "They spent most of their holiday working on a CRM",
		NoteBody:   "IDK what to put as the body.",
	}
	antp := NewAddNoteToPerson(pr, nr)
	pn, err := antp.Execute(ctx, antpr)
	require.NoError(t, err)

	lpn := NewListPersonNotes(pr, nr)
	ns, err := lpn.Execute(ctx, ListPersonNotesRequest{
		PersonUUID: person.UUID,
		Filter:     OnlyPersonNotesFilter,
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(ns))
	require.Equal(t, pn.UUID, ns[0].UUID)

	ns, err = lpn.Execute(ctx, ListPersonNotesRequest{
		PersonUUID: person.UUID,
		Filter:     AllNotesFilter,
	})
	require.NoError(t, err)
	require.Equal(t, 2, len(ns))
}
