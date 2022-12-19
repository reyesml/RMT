package identity

import (
	"fmt"
	"github.com/reyesml/RMT/app/core/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	database.BaseModel
	Username     string
	PasswordHash string
}

func NewUser(uname string, pass string) (*User, error) {
	if len(uname) == 0 {
		return nil, fmt.Errorf("username is required")
	}

	u := &User{Username: uname}
	err := u.SetPassword(pass)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) SetPassword(pass string) error {
	if len(pass) < 10 {
		return fmt.Errorf("passwords must be at least 10 characters long")
	}
	pwHash, err := hashAndSalt(pass)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	u.PasswordHash = pwHash
	return nil
}

func (u *User) isPasswordCorrect(input string) bool {
	byteHash := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(input))
	return err == nil
}

func hashAndSalt(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
