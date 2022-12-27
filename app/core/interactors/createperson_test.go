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

func TestCreatePerson_Execute(t *testing.T) {
	testDBId := "TestCreatePerson.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	user, err := models.NewUser("foobar", "some_password")
	require.NoError(t, err)
	require.NoError(t, userRepo.Create(user))

	personRepo := repos.NewPersonRepo(db)
	cp := NewCreatePerson(personRepo)
	ctx := utils.SetCurrentUser(context.Background(), *user, user.SegmentUUID)
	require.NoError(t, err)
	req := CreatePersonRequest{
		FirstName: "Sam",
		LastName:  "Doe",
	}

	person, err := cp.Execute(ctx, req)
	require.NoError(t, err)
	require.Equal(t, req.LastName, person.LastName)
	require.Equal(t, req.FirstName, person.FirstName)
	require.NotEqual(t, uuid.Nil, person.UUID)
	require.NotEqual(t, uuid.Nil, person.SegmentUUID)
	require.Equal(t, user.SegmentUUID, person.SegmentUUID)

	//attempt insert with nil segment
	ctx2 := context.WithValue(context.Background(), user, uuid.Nil)
	_, err = cp.Execute(ctx2, req)
	require.Error(t, err)
}
