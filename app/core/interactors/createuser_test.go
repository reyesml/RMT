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
	"strings"
	"testing"
)

func TestCreateUser_Execute(t *testing.T) {
	testDBId := "TestCreateUser.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)

	require.NoError(t, err)

	err = utils.MigrateAllModels(db)
	require.NoError(t, err)

	userRepo := repos.NewUserRepo(db)
	adminUser, err := models.NewUser("admin", "admin_password")
	require.NoError(t, err)
	adminUser.Admin = true
	require.NoError(t, userRepo.Create(adminUser))
	ctx := utils.SetCurrentUser(context.Background(), *adminUser, uuid.Nil)

	req := CreateUserRequest{
		Username: "foobar",
		Password: "plaintext_password",
	}
	cu := NewCreateUser(userRepo)

	resp, err := cu.Execute(ctx, req)
	require.NoError(t, err)
	require.Equal(t, req.Username, resp.Username)
	require.NotEmpty(t, resp.UUID.String())
	require.NotEqual(t, uuid.Nil, resp.UUID)

	user, err := userRepo.GetByUUID(resp.UUID)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)
	require.NotEqual(t, uuid.Nil, user.UUID)
	require.NotEqual(t, uuid.Nil, user.SegmentUUID)

	require.NotEqual(t, req.Password, user.PasswordHash)
	require.True(t, user.IsPasswordCorrect("plaintext_password"))
	require.False(t, user.IsPasswordCorrect("wrong_password"))

	//Try username collision with different case text
	req2 := CreateUserRequest{
		Username: strings.ToUpper(req.Username),
		Password: "plaintext_password2",
	}
	resp, err = cu.Execute(context.Background(), req2)
	require.Error(t, err)

	//Try non-collision
	req3 := CreateUserRequest{
		Username: fmt.Sprintf("%v2", req.Username),
		Password: "plaintext_password2",
	}
	resp, err = cu.Execute(ctx, req3)
	require.NoError(t, err)
}
