package interactors

import (
	"context"
	"fmt"
	"github.com/reyesml/RMT/app/core/identity"
	"github.com/reyesml/RMT/app/repos"
	"time"
)

type CreateSessionRequest struct {
	Username string
	Password string
}

type CreateSessionResponse struct {
	Token      string
	Expiration time.Time
}

type CreateSession struct {
	UserRepo      repos.UserRepo
	SessionRepo   repos.SessionRepo
	SigningSecret string
}

func (ia CreateSession) Execute(ctx context.Context, req CreateSessionRequest) (CreateSessionResponse, error) {
	_ = ctx
	user, err := ia.UserRepo.GetByUsername(req.Username)
	if err != nil {
		return CreateSessionResponse{}, err
	}
	if !user.IsPasswordCorrect(req.Password) {
		return CreateSessionResponse{}, fmt.Errorf("invalid password")
	}
	session := identity.NewSession(user)
	if err = ia.SessionRepo.Create(session); err != nil {
		return CreateSessionResponse{}, err
	}

	token, err := session.GenerateJWT(ia.SigningSecret)
	if err != nil {
		return CreateSessionResponse{}, fmt.Errorf("failed to generate jwt: %v", err)
	}

	return CreateSessionResponse{
		Token:      token,
		Expiration: session.Expiration,
	}, nil
}
