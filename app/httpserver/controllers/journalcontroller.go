package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/interactors"
	"github.com/reyesml/RMT/app/core/models"
	"net/http"
	"time"
)

type JournalController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

func NewJournalController(cje interactors.CreateJournal, gje interactors.GetJournal, lje interactors.ListJournals) journalController {
	return journalController{
		cje: cje,
		gje: gje,
		lje: lje,
	}
}

type journalController struct {
	cje interactors.CreateJournal
	gje interactors.GetJournal
	lje interactors.ListJournals
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

	if errors.Is(err, interactors.MissingJournalTitleErr) {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreateJournalResponse{Error: "title is required"})
		return
	}
	if errors.Is(err, interactors.MissingJournalBodyErr) {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreateJournalResponse{Error: "body is required"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, CreateJournalResponse{Error: "something went wrong"})
		return
	}

	render.JSON(w, r, CreateJournalResponse{
		UUID: je.UUID,
	})
}

type JournalResponse struct {
	UUID          uuid.UUID `json:"uuid"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	Mood          string    `json:"mood"`
	CreatedByUUID uuid.UUID `json:"createdByUUID"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type GetJournalResponse struct {
	Error   string          `json:"error,omitempty"`
	Journal JournalResponse `json:"journal,omitempty"`
}

func (c journalController) Get(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, GetJournalResponse{Error: "not found"})
		return
	}

	je, err := c.gje.Execute(r.Context(), interactors.GetJournalRequest{UUID: reqUUID})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, GetJournalResponse{Error: "not found"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, GetJournalResponse{Error: "something went wrong"})
		return
	}
	render.JSON(w, r, buildJournalResponse(je))
}

type ListJournalResponse struct {
	Error    string            `json:"error,omitempty"`
	Journals []JournalResponse `json:"journals,omitempty"`
}

func (c journalController) List(w http.ResponseWriter, r *http.Request) {
	jes, err := c.lje.Execute(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, GetJournalResponse{Error: "something went wrong"})
		return
	}
	var result []JournalResponse
	for _, je := range jes {
		result = append(result, buildJournalResponse(je))
	}
	render.JSON(w, r, ListJournalResponse{
		Journals: result,
	})
}

func buildJournalResponse(je models.Journal) JournalResponse {
	return JournalResponse{
		UUID:          je.UUID,
		Title:         je.Title,
		Body:          je.Body,
		Mood:          je.Mood,
		CreatedByUUID: je.User.UUID,
		CreatedAt:     je.CreatedAt,
		UpdatedAt:     je.UpdatedAt,
	}
}
