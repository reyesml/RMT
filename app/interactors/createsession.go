package interactors

import (
	"context"
	"errors"
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

var BadCredErr = errors.New("invalid credentials")

type CreateSession interface {
	Execute(ctx context.Context, req CreateSessionRequest) (CreateSessionResponse, error)
}

func NewCreateSession(userRepo repos.UserRepo, sessionRepo repos.SessionRepo, signingSecret string) createSession {
	return createSession{
		userRepo:      userRepo,
		sessionRepo:   sessionRepo,
		signingSecret: signingSecret,
	}
}

type createSession struct {
	userRepo      repos.UserRepo
	sessionRepo   repos.SessionRepo
	signingSecret string
}

func (ia createSession) Execute(ctx context.Context, req CreateSessionRequest) (CreateSessionResponse, error) {
	_ = ctx
	user, err := ia.userRepo.GetByUsername(req.Username)
	if err != nil {
		return CreateSessionResponse{}, err
	}
	if !user.IsPasswordCorrect(req.Password) {
		return CreateSessionResponse{}, BadCredErr
	}
	session := identity.NewSession(user)
	if err = ia.sessionRepo.Create(session); err != nil {
		return CreateSessionResponse{}, err
	}

	token, err := session.GenerateJWT(ia.signingSecret)
	if err != nil {
		return CreateSessionResponse{}, fmt.Errorf("failed to generate jwt: %w", err)
	}

	return CreateSessionResponse{
		Token:      token,
		Expiration: session.Expiration,
	}, nil
}
