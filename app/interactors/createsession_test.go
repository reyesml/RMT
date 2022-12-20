package interactors

import (
	"context"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/identity"
	"github.com/reyesml/RMT/app/repos"
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

	err = repos.MigrateAll(db)
	require.NoError(t, err)

	userRepo := repos.NewUserRepo(db)
	user, err := identity.NewUser("test_user", "password123")
	require.NoError(t, err)
	require.NoError(t, userRepo.Create(user))

	sessionRepo := repos.NewSessionRepo(db)
	req := CreateSessionRequest{
		Username: user.Username,
		Password: "password123",
	}
	var cs CreateSession
	resp, err := cs.Execute(context.Background(), userRepo, sessionRepo, req)
	require.NoError(t, err)
	require.NotEmpty(t, resp.Token)
	require.GreaterOrEqual(t, resp.Expiration, time.Now().UTC())

	session, err := sessionRepo.GetByToken(resp.Token)
	require.NoError(t, err)
	require.Equal(t, session.Expiration.UTC(), resp.Expiration.UTC())
}
