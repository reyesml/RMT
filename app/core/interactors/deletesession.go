package interactors

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type DeleteSession interface {
	Execute(ctx context.Context) error
}

func NewDeleteSession(sr repos.SessionRepo) deleteSession {
	return deleteSession{sr: sr}
}

type deleteSession struct {
	sr repos.SessionRepo
}

func (ia deleteSession) Execute(ctx context.Context) error {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.UUID == uuid.Nil {
		return fmt.Errorf("no current user")
	}

	s, err := ia.sr.GetByUUID(user.SessionUUID)
	if err != nil {
		return err
	}
	if err := ia.sr.Delete(&s); err != nil {
		return err
	}
	return nil
}
