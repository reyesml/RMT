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

type CreateSession struct{}

func (_ CreateSession) Execute(ctx context.Context, userRepo repos.UserRepo, sessionRepo repos.SessionRepo, req CreateSessionRequest) (CreateSessionResponse, error) {
	_ = ctx
	user, err := userRepo.GetByUsername(req.Username)
	if err != nil {
		return CreateSessionResponse{}, err
	}
	if !user.IsPasswordCorrect(req.Password) {
		return CreateSessionResponse{}, fmt.Errorf("invalid password")
	}
	session, err := identity.NewSession(user)
	if err != nil {
		return CreateSessionResponse{}, err
	}
	if err = sessionRepo.Create(session); err != nil {
		return CreateSessionResponse{}, err
	}
	return CreateSessionResponse{
		Token:      session.Token,
		Expiration: session.Expiration,
	}, nil
}
