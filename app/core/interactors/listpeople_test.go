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
		*models.NewPerson(seg1, "sam", "doe"),
		*models.NewPerson(seg1, "chris", "cross"),
		*models.NewPerson(seg1, "taylor", "smith"),
	}
	seg2people := []models.Person{
		*models.NewPerson(seg2, "jack", "frost"),
		*models.NewPerson(seg2, "jill", "frost"),
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
