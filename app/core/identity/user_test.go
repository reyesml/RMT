package identity

import (
	"github.com/reyesml/RMT/app/core/database"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	testDBId := "TestUser.db"
	db, err := database.Connect(testDBId)
	defer os.Remove(testDBId)
	require.NoError(t, err)

	err = db.AutoMigrate(&User{})

	ptpw := "plain_text_password"
	u, err := NewUser("test_user", ptpw)
	require.NoError(t, err)
	require.NotEqual(t, u.PasswordHash, ptpw)
	db.Create(&u)
	require.NotEmpty(t, u.UUID.String())

	db.Where("UUID = ?", u.UUID).First(&u)

	require.True(t, u.isPasswordCorrect(ptpw))
	oldPwHash := u.PasswordHash
	err = u.SetPassword(ptpw)
	require.NoError(t, err)
	require.NotEqual(t, oldPwHash, u.PasswordHash)
	require.True(t, u.isPasswordCorrect(ptpw))
	require.False(t, u.isPasswordCorrect("incorrect_password"))
}
