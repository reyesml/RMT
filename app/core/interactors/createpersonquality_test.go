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
	"strings"
	"testing"
)

func TestCreatePersonQuality_Execute(t *testing.T) {
	testDBId := "TestCreatePersonQuality.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	user, err := models.NewUser("test-user", "plain_password")
	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	require.NoError(t, userRepo.Create(user))

	pr := repos.NewPersonRepo(db)
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
	require.NoError(t, pr.CreateMany(&people))

	qr := repos.NewQualityRepo(db)
	pqr := repos.NewPersonQualityRepo(db)
	aqtp := NewCreatePersonQuality(pr, qr, pqr)
	ctx := utils.SetCurrentUser(context.Background(), *user, uuid.Nil)
	cpqr := CreatePersonQualityRequest{
		PersonUUID:  people[0].UUID,
		QualityName: "Loves Writing Tests",
	}
	pq, err := aqtp.Execute(ctx, cpqr)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, pq.UUID)
	require.Equal(t, user.SegmentUUID, pq.SegmentUUID)
	require.Equal(t, user.ID, pq.UserId)
	require.Equal(t, people[0].ID, pq.PersonId)

	q, err := qr.GetByID(pq.QualityId, user.SegmentUUID)
	require.NoError(t, err)
	require.Equal(t, user.SegmentUUID, q.SegmentUUID)
	require.NotEqual(t, uuid.Nil, q.UUID)
	require.Equal(t, cpqr.QualityName, q.Name)
	require.Equal(t, strings.ToLower(cpqr.QualityName), q.NameLower)

	cpqr2 := CreatePersonQualityRequest{
		PersonUUID:  people[1].UUID,
		QualityName: "LOVES WRITING TESTS",
	}
	pq2, err := aqtp.Execute(ctx, cpqr2)
	require.NoError(t, err)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, pq2.UUID)
	require.Equal(t, user.SegmentUUID, pq2.SegmentUUID)
	require.Equal(t, user.ID, pq2.UserId)
	require.Equal(t, people[1].ID, pq2.PersonId)

	// The quality had the same name as an existing quality,
	// so the existing quality should have been used.
	require.Equal(t, q.ID, pq2.QualityId)

	cpqr3 := CreatePersonQualityRequest{
		PersonUUID:  people[1].UUID,
		QualityName: "Hates writing tests",
	}
	pq3, err := aqtp.Execute(ctx, cpqr3)
	require.NoError(t, err)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, pq3.UUID)
	require.Equal(t, user.SegmentUUID, pq3.SegmentUUID)
	require.Equal(t, user.ID, pq3.UserId)
	require.Equal(t, people[1].ID, pq3.PersonId)

	// The quality doesn't match existing qualities,
	// so a new one should have been created.
	require.NotEqual(t, q.ID, pq3.QualityId)
}
