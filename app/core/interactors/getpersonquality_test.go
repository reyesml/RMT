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

func TestGetPersonQuality_Execute(t *testing.T) {
	testDBId := "TestGetPersonQuality.db"
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
	nr := repos.NewNoteRepo(db)
	aqtp := NewAddQualityToPerson(personRepo, qr, pqr)
	antpq := NewAddNoteToPersonQuality(pqr, nr)
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)

	//add a quality with a note to our test person
	cpqr := AddQualityToPersonRequest{
		PersonUUID:  person.UUID,
		QualityName: "some-quality",
	}
	pq, err := aqtp.Execute(ctx, cpqr)
	require.NoError(t, err)

	pn, err := antpq.Execute(ctx, AddNoteToPersonQualityRequest{
		PersonQualityUUID: pq.UUID,
		NoteTitle:         "quality note title",
		NoteBody:          "note body",
	})
	require.NoError(t, err)

	gpq := NewGetPersonQuality(pqr)
	foundPersonQuality, err := gpq.Execute(ctx, GetPersonQualityRequest{PersonQualityUUID: pq.UUID})
	require.NoError(t, err)
	require.Equal(t, pq.ID, foundPersonQuality.ID)
	require.Equal(t, cpqr.QualityName, foundPersonQuality.Quality.Name)
	require.Equal(t, 1, len(foundPersonQuality.Notes))
	require.Equal(t, pn.UUID, foundPersonQuality.Notes[0].UUID)
}
