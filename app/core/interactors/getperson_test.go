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

func TestGetPerson_Execute(t *testing.T) {
	testDBId := "TestGetPerson.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	user, err := models.NewUser("test-user", "plain_password")
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	require.NoError(t, userRepo.Create(user))
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)

	personRepo := repos.NewPersonRepo(db)
	p := models.Person{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		FirstName: "Sam",
		LastName:  "Piper",
		UserId:    user.ID,
	}
	require.NoError(t, personRepo.Create(&p))

	other := models.Person{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		FirstName: "Some Other",
		LastName:  "Person",
		UserId:    user.ID,
	}
	require.NoError(t, personRepo.Create(&other))

	gp := NewGetPerson(personRepo)
	foundPerson, err := gp.Execute(ctx, GetPersonRequest{UUID: p.UUID})
	require.NoError(t, err)
	require.Equal(t, p.ID, foundPerson.ID)
	require.Equal(t, p.UUID, foundPerson.UUID)
	require.NotEqual(t, uuid.Nil, foundPerson.UUID)
	require.Equal(t, p.FirstName, foundPerson.FirstName)
	require.Equal(t, p.LastName, foundPerson.LastName)
	require.Equal(t, user.ID, foundPerson.UserId)
	require.Equal(t, user.SegmentUUID, foundPerson.SegmentUUID)

	//Attempt to fetch person from a different segment
	user.SegmentUUID, err = uuid.NewRandom()
	require.NoError(t, err)
	ctx = context.WithValue(context.Background(), user, uuid.Nil)
	foundPerson, err = gp.Execute(ctx, GetPersonRequest{UUID: p.UUID})
	require.Error(t, err)
	require.Equal(t, uuid.Nil, foundPerson.UUID)
}
