package utils

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/models"
)

const UserCtxKey = "current-user"

var UserMissingErr = errors.New("user not provided")

func SetCurrentUser(ctx context.Context, usr models.User, sessionUUID uuid.UUID) context.Context {
	return context.WithValue(ctx, UserCtxKey, models.CurrentUser{
		User:        usr,
		SessionUUID: sessionUUID,
	})
}

func GetCurrentUser(ctx context.Context) (models.CurrentUser, error) {
	user, ok := ctx.Value(UserCtxKey).(models.CurrentUser)
	if !ok || user.ID == 0 {
		return models.CurrentUser{}, UserMissingErr
	}
	return user, nil
}
