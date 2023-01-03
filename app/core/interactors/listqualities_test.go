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

func TestListQualities_Execute(t *testing.T) {
	testDBId := "TestQualities.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	user, err := models.NewUser("test-user", "plain_password")
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	require.NoError(t, userRepo.Create(user))
	qr := repos.NewQualityRepo(db)
	qs := []models.Quality{
		{
			Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
			Name:      "q1",
			NameLower: "q1",
			UserId:    user.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
			Name:      "q2",
			NameLower: "q2",
			UserId:    user.ID,
		},
		{
			Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
			Name:      "q3",
			NameLower: "q3",
			UserId:    user.ID,
		},
	}
	require.NoError(t, qr.CreateMany(&qs))

	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	lq := NewListQualities(qr)
	foundQualities, err := lq.Execute(ctx)
	require.NoError(t, err)
	require.Equal(t, len(qs), len(foundQualities))

	//swap out a segment UUID and try the request again
	user.SegmentUUID, err = uuid.NewRandom()
	require.NoError(t, err)
	ctx = utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	foundQualities, err = lq.Execute(ctx)
	require.NoError(t, err)
	require.Equal(t, 0, len(foundQualities))
}
