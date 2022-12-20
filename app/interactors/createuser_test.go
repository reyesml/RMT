package interactors

import (
	"context"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/repos"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCreateUser_Execute(t *testing.T) {
	testDBId := "TestCreateUser.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	err = repos.MigrateAll(db)
	require.NoError(t, err)

	require.NoError(t, err)
	userRepo := repos.NewUserRepo(db)
	req := CreateUserRequest{
		Username: "foobar",
		Password: "plaintext_password",
	}
	var cu CreateUser
	resp, err := cu.Execute(context.Background(), userRepo, req)
	require.NoError(t, err)
	require.Equal(t, req.Username, resp.Username)
	require.NotEmpty(t, resp.UUID)

	user, err := userRepo.GetByUUID(resp.UUID)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	require.True(t, user.IsPasswordCorrect("plaintext_password"))
	require.False(t, user.IsPasswordCorrect("wrong_password"))
}
