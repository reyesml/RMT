package identity

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"time"
)

const sessionDuration = 15 * time.Minute

type Session struct {
	database.BaseModel
	UserId     uint
	User       User
	Expiration time.Time
}

func NewSession(user User) *Session {
	session := &Session{
		UserId:     user.ID,
		User:       user,
		Expiration: time.Now().UTC().Add(sessionDuration),
	}
	return session
}

func (s *Session) GenerateJWT(signingSecret string) (string, error) {
	if s.UUID == uuid.Nil {
		return "", fmt.Errorf("session must have uuid set before token generation")
	}
	if s.Expiration.IsZero() {
		return "", fmt.Errorf("session must have an expiration specified")
	}
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(s.Expiration),
		Issuer:    "rmt",
		ID:        s.UUID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingSecret))
}

func GetSessionUUIDFromJWT(tokenstr string, signingSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenstr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingSecret), nil
	})

	if err != nil {
		return uuid.Nil, err
	} else if !token.Valid {
		return uuid.Nil, fmt.Errorf("very bad token")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return uuid.Nil, fmt.Errorf("could not read claims")
	}
	return uuid.Parse(claims.ID)
}
