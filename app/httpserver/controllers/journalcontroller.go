package controllers

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/interactors"
	"net/http"
)

type JournalController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewJournalController(cje interactors.CreateJournal) journalController {
	return journalController{cje: cje}
}

type journalController struct {
	cje interactors.CreateJournal
}

type CreateJournalRequest struct {
	Mood  string `json:"mood"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type CreateJournalResponse struct {
	Error string    `json:"error,omitempty"`
	UUID  uuid.UUID `json:"uuid"`
}

func (c journalController) Create(w http.ResponseWriter, r *http.Request) {
	var createReq CreateJournalRequest
	err := json.NewDecoder(r.Body).Decode(&createReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreateJournalResponse{Error: "invalid body format"})
		return
	}

	je, err := c.cje.Execute(r.Context(), interactors.CreateJournalRequest{
		Mood:  createReq.Mood,
		Title: createReq.Title,
		Body:  createReq.Body,
	})
	// TODO: different responses based off of user error vs server error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, CreateJournalResponse{Error: "something went wrong"})
		return
	}

	render.JSON(w, r, CreateJournalResponse{
		UUID: je.UUID,
	})

}
