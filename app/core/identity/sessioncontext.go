package identity

import "github.com/google/uuid"

const SessionContextKey = "rmt-session"

type SessionContext struct {
	UserUUID    uuid.UUID
	SessionUUID uuid.UUID
}
