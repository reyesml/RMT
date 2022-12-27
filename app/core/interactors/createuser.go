package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
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

func NewCreateUser(userRepo repos.UserRepo) createUser {
	return createUser{userRepo: userRepo}
}

type createUser struct {
	userRepo repos.UserRepo
}

func (ia createUser) Execute(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	currUser, err := utils.GetCurrentUser(ctx)
	if err != nil || !currUser.Admin {
		return CreateUserResponse{}, fmt.Errorf("create user: requires admin")
	}
	users, err := ia.userRepo.FindByUsername(req.Username)
	if err != nil {
		return CreateUserResponse{}, err
	}
	if len(users) > 0 {
		return CreateUserResponse{}, fmt.Errorf("username already exists: %v", req.Username)
	}

	user, err := models.NewUser(req.Username, req.Password)
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
