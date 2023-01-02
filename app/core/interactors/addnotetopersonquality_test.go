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

func TestAddNoteToPersonQuality_Execute(t *testing.T) {
	testDBId := "TestAddNoteToPersonQuality.db"
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
	aqtp := NewAddQualityToPerson(personRepo, qr, pqr)
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	cpqr := AddQualityToPersonRequest{
		PersonUUID:  person.UUID,
		QualityName: "Loves Writing Tests",
	}
	pq, err := aqtp.Execute(ctx, cpqr)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, pq.UUID)

	antpqr := AddNoteToPersonQualityRequest{
		PersonQualityUUID: pq.UUID,
		NoteTitle:         "They insist on testing every interactor",
		NoteBody:          "...but mostly only the happy paths (:",
	}
	nr := repos.NewNoteRepo(db)
	antpq := NewAddNoteToPersonQuality(pqr, nr)
	n, err := antpq.Execute(ctx, antpqr)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, n.UUID)
	require.Equal(t, person.ID, n.PersonId)
	require.Equal(t, pq.ID, n.PersonQualityId)
	require.Equal(t, user.SegmentUUID, pq.SegmentUUID)
	require.Equal(t, antpqr.NoteTitle, n.Title)
	require.Equal(t, antpqr.NoteBody, n.Body)
}
