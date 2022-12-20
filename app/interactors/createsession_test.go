package interactors

import (
	"context"
	"github.com/google/uuid"
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
		Username: "TeSt_UsEr",
		Password: "password123",
	}
	var cs = CreateSession{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}
	resp, err := cs.Execute(context.Background(), req)
	require.NoError(t, err)
	require.NotEmpty(t, resp.Token)
	require.GreaterOrEqual(t, resp.Expiration, time.Now().UTC())

	session, err := sessionRepo.GetByToken(resp.Token)
	require.NoError(t, err)
	require.NotEqual(t, uuid.UUID{}, session.UUID)
	require.Equal(t, session.Expiration.UTC(), resp.Expiration.UTC())
}
