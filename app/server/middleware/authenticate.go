package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/repos"
	"net/http"
	"strings"
)

const RequestSessionKey = "rmt-session"

type RequestSession struct {
	UserUUID    uuid.UUID
	SessionUUID uuid.UUID
}

// Authenticate returns a middleware handler. This handler extracts the auth token
// from the request, giving precedence to the Authorization header, then the cookie.
// If the token is valid, the user id and session id's are added to the request
// context. If the token is invalid, a 401 unauthorized is returned to the client.
func Authenticate(sessionRepo repos.SessionRepo) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := getAuthToken(r)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// query db for matching token
			session, err := sessionRepo.GetByTokenWithUser(token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// add our session to the request context
			ctx := context.WithValue(r.Context(), RequestSessionKey, RequestSession{
				UserUUID:    session.User.UUID,
				SessionUUID: session.UUID,
			})
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// getAuthToken attempts to retrieve the auth token from the Authorization header,
// "Authorization: Bearer {token}". If the auth header is missing, the "access_token"
// cookie will be used instead.
func getAuthToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) > 0 {
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			return "", fmt.Errorf("invalid auth header format")
		}
		return splitToken[1], nil
	}

	authCookie, err := r.Cookie("access_token")
	if err != nil || len(authCookie.Value) == 0 {
		return "", fmt.Errorf("auth not found")
	}
	return authCookie.Value, nil
}
