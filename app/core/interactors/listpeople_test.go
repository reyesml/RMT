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

func TestListPeople_Execute(t *testing.T) {
	testDBId := "TestListPeople.db"
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
	seg1, err := uuid.NewRandom()
	require.NoError(t, err)
	seg2, err := uuid.NewRandom()
	require.NoError(t, err)

	seg1people := []models.Person{
		{
			Segmented: database.Segmented{SegmentUUID: seg1},
			FirstName: "sam",
			LastName:  "doe",
			UserId:    user.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: seg1},
			FirstName: "chris",
			LastName:  "cross",
			UserId:    user.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: seg1},
			FirstName: "taylor",
			LastName:  "smith",
			UserId:    user.ID,
		},
	}
	seg2people := []models.Person{
		{
			Segmented: database.Segmented{SegmentUUID: seg2},
			FirstName: "jack",
			LastName:  "frost",
			UserId:    user.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: seg2},
			FirstName: "jill",
			LastName:  "frost",
			UserId:    user.ID,
		},
	}

	require.NoError(t, personRepo.CreateMany(&seg1people))
	require.NoError(t, personRepo.CreateMany(&seg2people))

	listPeople := NewListPeople(personRepo)
	user.SegmentUUID = seg1
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	seg1list, err := listPeople.Execute(ctx)
	require.NoError(t, err)
	require.Equal(t, len(seg1people), len(seg1list))
	for i, p := range seg1list {
		require.Equal(t, p.UUID, seg1people[i].UUID)
	}

	user.SegmentUUID = seg2
	ctx = utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	seg2list, err := listPeople.Execute(ctx)
	require.NoError(t, err)
	require.Equal(t, len(seg2people), len(seg2list))
	for i, p := range seg2list {
		require.Equal(t, p.UUID, seg2people[i].UUID)
	}
}
