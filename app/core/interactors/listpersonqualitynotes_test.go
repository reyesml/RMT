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

func TestListPersonQualityNotes_Execute(t *testing.T) {
	testDBId := "TestListPersonQualityNotes.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	user, err := models.NewUser("test-user", "plain_password")
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	require.NoError(t, userRepo.Create(user))

	personRepo := repos.NewPersonRepo(db)
	person := models.Person{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		FirstName: "sam",
		LastName:  "doe",
		UserId:    user.ID,
	}
	require.NoError(t, personRepo.Create(&person))

	qr := repos.NewQualityRepo(db)
	pqr := repos.NewPersonQualityRepo(db)
	aqtp := NewCreatePersonQuality(personRepo, qr, pqr)
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	cpqr := CreatePersonQualityRequest{
		PersonUUID:  person.UUID,
		QualityName: "Loves Writing Tests",
	}
	pq, err := aqtp.Execute(ctx, cpqr)
	require.NoError(t, err)

	antpqr := CreatePersonQualityNoteRequest{
		PersonQualityUUID: pq.UUID,
		NoteTitle:         "They insist on testing every interactor",
		NoteBody:          "...but mostly only the happy paths (:",
	}
	nr := repos.NewNoteRepo(db)
	antpq := NewCreatePersonQualityNote(pqr, nr)
	//create some notes
	for i := 0; i < 3; i++ {
		_, err = antpq.Execute(ctx, antpqr)
		require.NoError(t, err)
	}

	lpqn := NewListPersonQualityNotes(pqr, nr)
	ns, err := lpqn.Execute(ctx, ListPersonQualityNotesRequest{PersonQualityUUID: pq.UUID})
	require.NoError(t, err)
	require.Equal(t, 3, len(ns))
	//Make sure our relations are loaded
	n := ns[0]
	require.Equal(t, person.UUID, n.Person.UUID)
	require.Equal(t, pq.UUID, n.PersonQuality.UUID)
}
