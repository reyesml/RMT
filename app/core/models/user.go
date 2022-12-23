package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	database.BaseModel
	database.Segmented
	Username      string `gorm:"unique"`
	UsernameLower string `gorm:"unique;not null;check:username_lower <> ''"`
	PasswordHash  string
	Admin         bool
}

func NewUser(uname string, pass string) (*User, error) {
	if len(uname) == 0 {
		return nil, fmt.Errorf("username is required")
	}

	segmentUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("segment uuid: %w", err)
	}
	u := &User{
		Username:  uname,
		Segmented: database.Segmented{SegmentUUID: segmentUUID},
	}
	err = u.SetPassword(pass)
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
		return fmt.Errorf("failed to hash password: %w", err)
	}
	u.PasswordHash = pwHash
	return nil
}

func (u *User) IsPasswordCorrect(input string) bool {
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
