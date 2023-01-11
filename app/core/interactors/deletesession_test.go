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

func TestDeleteSession_Execute(t *testing.T) {
	testDBId := "TestDeleteSession.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	userRepo := repos.NewUserRepo(db)
	user, err := models.NewUser("test_user", "password123")
	require.NoError(t, err)
	require.NoError(t, userRepo.Create(user))
	require.NotEqual(t, uuid.Nil, user.UUID)

	sessionRepo := repos.NewSessionRepo(db)
	req := CreateSessionRequest{
		Username: "TeSt_UsEr",
		Password: "password123",
	}

	//Create two sessions...
	cs := NewCreateSession(userRepo, sessionRepo, "secret")
	resp, err := cs.Execute(context.Background(), req)
	require.NoError(t, err)

	sessionUUID1, err := models.GetSessionUUIDFromJWT(resp.Token, "secret")
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, sessionUUID1)

	resp, err = cs.Execute(context.Background(), req)
	require.NoError(t, err)

	sessionUUID2, err := models.GetSessionUUIDFromJWT(resp.Token, "secret")
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, sessionUUID2)

	//clear the first session
	ds := NewDeleteSession(sessionRepo)
	ctx := utils.SetCurrentUser(context.Background(), *user, sessionUUID1)
	require.NoError(t, ds.Execute(ctx))

	//first session should be deleted now
	sess1, err := sessionRepo.GetByUUID(sessionUUID1)
	require.Error(t, err)
	require.Equal(t, uuid.Nil, sess1.UUID)

	//other session still exists
	sess2, err := sessionRepo.GetByUUID(sessionUUID2)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, sess2.UUID)
}
