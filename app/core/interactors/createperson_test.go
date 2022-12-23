package interactors

import (
	"context"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
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

	personRepo := repos.NewPersonRepo(db)
	cp := NewCreatePerson(personRepo)
	segment, err := uuid.NewRandom()
	require.NoError(t, err)
	req := CreatePersonRequest{
		FirstName:   "Sam",
		LastName:    "Doe",
		SegmentUUID: segment,
	}

	person, err := cp.Execute(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, req.LastName, person.LastName)
	require.Equal(t, req.FirstName, person.FirstName)
	require.NotEqual(t, uuid.Nil, person.UUID)
	require.NotEqual(t, uuid.Nil, person.SegmentUUID)
	require.Equal(t, req.SegmentUUID, person.SegmentUUID)

	//attempt insert with nil segment
	req2 := CreatePersonRequest{
		FirstName: "John",
		LastName:  "Doe",
	}
	_, err = cp.Execute(context.Background(), req2)
	require.Error(t, err)
}
