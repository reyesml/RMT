package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestListPersonQualities_Execute(t *testing.T) {
	testDBId := "TestListPersonQualities.db"
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
	people := []models.Person{
		{
			Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
			FirstName: "sam",
			LastName:  "doe",
			UserId:    user.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
			FirstName: "chris",
			LastName:  "cross",
			UserId:    user.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
			FirstName: "taylor",
			LastName:  "smith",
			UserId:    user.ID,
		},
	}
	require.NoError(t, personRepo.CreateMany(&people))

	qr := repos.NewQualityRepo(db)
	pqr := repos.NewPersonQualityRepo(db)
	aqtp := NewCreatePersonQuality(personRepo, qr, pqr)
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)

	//add some qualities to the first person in our list.
	for i := 0; i < 5; i++ {
		cpqr := CreatePersonQualityRequest{
			PersonUUID:  people[0].UUID,
			QualityName: fmt.Sprintf("quality-%d", i),
		}
		_, err = aqtp.Execute(ctx, cpqr)
		require.NoError(t, err)
	}

	lpqs := NewListPersonQualities(personRepo, pqr)
	pqs, err := lpqs.Execute(ctx, ListPersonQualitiesRequest{
		PersonUUID: people[0].UUID,
	})
	require.NoError(t, err)
	require.Equal(t, 5, len(pqs))
	require.NotEqual(t, 0, len(pqs[0].Quality.Name))

	pqs, err = lpqs.Execute(ctx, ListPersonQualitiesRequest{
		PersonUUID: people[1].UUID,
	})
	require.NoError(t, err)
	require.Equal(t, 0, len(pqs))
}
