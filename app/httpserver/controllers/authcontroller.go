package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/reyesml/RMT/app/core/interactors"
	"net/http"
	"time"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	TokenInBody bool   `json:"tokenInBody"`
}

type LoginResponse struct {
	Error      string    `json:"error,omitempty"`
	Token      string    `json:"token,omitempty"`
	Expiration time.Time `json:"expiration"`
}

func NewAuthController(createSession interactors.CreateSession) authController {
	return authController{
		createSession: createSession,
	}
}

type authController struct {
	createSession interactors.CreateSession
}

func (c authController) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, LoginResponse{Error: "invalid body format"})
		return
	}

	result, err := c.createSession.Execute(r.Context(), interactors.CreateSessionRequest{
		Username: loginReq.Username,
		Password: loginReq.Password,
	})
	if errors.Is(err, interactors.BadCredErr) {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, LoginResponse{Error: "invalid login"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, LoginResponse{Error: "unknown"})
		return
	}

	resp := LoginResponse{Expiration: result.Expiration}

	if loginReq.TokenInBody {
		resp.Token = result.Token
		render.JSON(w, r, resp)
		return
	} else {
		authCookie := http.Cookie{
			Name:     "session",
			Value:    result.Token,
			Path:     "/",
			Expires:  result.Expiration,
			HttpOnly: true,
		}
		http.SetCookie(w, &authCookie)
		return
	}
}
