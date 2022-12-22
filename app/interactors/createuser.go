package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/identity"
	"github.com/reyesml/RMT/app/repos"
)

type CreateUserRequest struct {
	Username string
	Password string
}

type CreateUserResponse struct {
	Username string
	UUID     uuid.UUID
}

type CreateUser interface {
	Execute(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
}

func NewCreateUser(userRepo repos.UserRepo) CreateUser {
	return createUser{userRepo: userRepo}
}

type createUser struct {
	userRepo repos.UserRepo
}

func (ia createUser) Execute(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	_ = ctx
	users, err := ia.userRepo.FindByUsername(req.Username)
	if err != nil {
		return CreateUserResponse{}, err
	}
	if len(users) > 0 {
		return CreateUserResponse{}, fmt.Errorf("username already exists: %v", req.Username)
	}

	user, err := identity.NewUser(req.Username, req.Password)
	if err != nil {
		return CreateUserResponse{}, err
	}

	err = ia.userRepo.Create(user)
	if err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		Username: user.Username,
		UUID:     user.UUID,
	}, nil
}
