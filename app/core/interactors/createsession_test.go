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
	"time"
)

func TestCreateSession_Execute(t *testing.T) {
	testDBId := "TestCreateSession.db"
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
	cs := NewCreateSession(userRepo, sessionRepo, "secret")
	resp, err := cs.Execute(context.Background(), req)
	require.NoError(t, err)
	require.NotEmpty(t, resp.Token)
	require.GreaterOrEqual(t, resp.Expiration, time.Now().UTC())
	require.Equal(t, user.UUID, resp.User.UUID)

	//simulate a mismatch on signatures
	_, err = models.GetSessionUUIDFromJWT(resp.Token, "wrong-secret")
	require.Error(t, err)

	//retrieve using matching secret/valid signature
	sessionUUID, err := models.GetSessionUUIDFromJWT(resp.Token, "secret")
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, sessionUUID)

	session, err := sessionRepo.GetByUUIDWithUser(sessionUUID)
	require.NoError(t, err)
	require.Equal(t, user.UUID, session.User.UUID)
	require.Equal(t, session.Expiration.UTC(), resp.Expiration.UTC())
}
