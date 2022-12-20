package identity

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/reyesml/RMT/app/core/database"
	"time"
)

const sessionDuration = 15 * time.Minute

type Session struct {
	database.BaseModel
	UserId     uint
	User       User
	Token      string
	Expiration time.Time
}

func NewSession(user User) (*Session, error) {
	token, err := generateToken()
	if err != nil {
		return nil, err
	}
	return &Session{
		UserId:     user.ID,
		User:       user,
		Token:      token,
		Expiration: time.Now().UTC().Add(sessionDuration),
	}, nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate session token: %v", err)
	}
	return hex.EncodeToString(b), nil
}
