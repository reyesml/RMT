package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/repos"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestCreateUser_Execute(t *testing.T) {
	testDBId := "TestCreateUser.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = repos.MigrateAll(db)
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
	require.NotEmpty(t, resp.UUID.String())
	require.NotEqual(t, uuid.UUID{}, resp.UUID)

	user, err := userRepo.GetByUUID(resp.UUID)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	require.NotEqual(t, req.Password, user.PasswordHash)
	require.True(t, user.IsPasswordCorrect("plaintext_password"))
	require.False(t, user.IsPasswordCorrect("wrong_password"))

	//Try username collision with different case text
	req2 := CreateUserRequest{
		Username: strings.ToUpper(req.Username),
		Password: "plaintext_password2",
	}
	resp, err = cu.Execute(context.Background(), userRepo, req2)
	require.Error(t, err)

	//Try non-collision
	req3 := CreateUserRequest{
		Username: fmt.Sprintf("%v2", req.Username),
		Password: "plaintext_password2",
	}
	resp, err = cu.Execute(context.Background(), userRepo, req3)
	require.NoError(t, err)
}
