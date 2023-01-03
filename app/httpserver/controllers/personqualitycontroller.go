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

type PersonQualityController interface {
	Get(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	ListNotes(w http.ResponseWriter, r *http.Request)
}

func NewPersonQualityController(gpq interactors.GetPersonQuality, cpqn interactors.CreatePersonQualityNote, lpqn interactors.ListPersonQualityNotes) personQualityController {
	return personQualityController{
		gpq:  gpq,
		cpqn: cpqn,
		lpqn: lpqn,
	}
}

type personQualityController struct {
	gpq  interactors.GetPersonQuality
	cpqn interactors.CreatePersonQualityNote
	lpqn interactors.ListPersonQualityNotes
}

type GetPersonQualityResponse struct {
	Error         string        `json:"error,omitempty"`
	PersonQuality PersonQuality `json:"personQuality,omitempty"`
}

type PersonQuality struct {
	UUID       uuid.UUID `json:"uuid"`
	PersonUUID uuid.UUID `json:"personUUID"`
	Name       string    `json:"name"`
	Notes      []Note    `json:"notes"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Note struct {
	UUID              uuid.UUID `json:"uuid"`
	PersonUUID        uuid.UUID `json:"personUUID"`
	PersonQualityUUID uuid.UUID `json:"personQualityUUID"`
	Title             string    `json:"title"`
	Body              string    `json:"body"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (c personQualityController) Get(w http.ResponseWriter, r *http.Request) {
	pqUUIDParam := chi.URLParam(r, "UUID")
	pqUUID, err := uuid.Parse(pqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, GetPersonQualityResponse{Error: "not found"})
		return
	}

	pq, err := c.gpq.Execute(r.Context(), interactors.GetPersonQualityRequest{PersonQualityUUID: pqUUID})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, GetPersonQualityResponse{Error: "not found"})
		return
	}

	render.JSON(w, r, GetPersonQualityResponse{PersonQuality: mapPersonQuality(pq)})
}

type CreatePersonQualityNote struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type CreatePersonQualityNoteResponse struct {
	Error string    `json:"error,omitempty"`
	UUID  uuid.UUID `json:"uuid"`
}

func (c personQualityController) CreateNote(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, CreatePersonQualityNoteResponse{Error: "not found"})
		return
	}
	var req CreatePersonNoteRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreatePersonQualityNoteResponse{Error: "invalid body format"})
		return
	}

	n, err := c.cpqn.Execute(r.Context(), interactors.CreatePersonQualityNoteRequest{
		PersonQualityUUID: reqUUID,
		NoteTitle:         req.Title,
		NoteBody:          req.Body,
	})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, CreatePersonQualityNoteResponse{Error: "not found"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, CreatePersonQualityNoteResponse{Error: "something went wrong"})
		return
	}
	render.JSON(w, r, CreatePersonQualityNoteResponse{UUID: n.UUID})
}

type ListPersonQualityNotesResponse struct {
	Error string `json:"error,omitempty"`
	Notes []Note `json:"notes"`
}

func (c personQualityController) ListNotes(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, ListPersonQualityNotesResponse{Error: "not found"})
		return
	}

	ns, err := c.lpqn.Execute(r.Context(), interactors.ListPersonQualityNotesRequest{
		PersonQualityUUID: reqUUID,
	})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, ListPersonQualityNotesResponse{Error: "not found"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, ListPersonQualityNotesResponse{Error: "something went wrong"})
		return
	}
	result := make([]Note, 0)
	for _, n := range ns {
		result = append(result, mapNote(n, n.Person.UUID, n.PersonQuality.UUID))
	}
	render.JSON(w, r, ListPersonNotesResponse{Notes: result})
}

func mapPersonQuality(pq models.PersonQuality) PersonQuality {
	notes := make([]Note, 0)
	for _, n := range pq.Notes {
		notes = append(notes, mapNote(n, pq.Person.UUID, pq.UUID))
	}
	return PersonQuality{
		UUID:       pq.UUID,
		PersonUUID: pq.Person.UUID,
		Name:       pq.Quality.Name,
		Notes:      notes,
		CreatedAt:  pq.CreatedAt,
		UpdatedAt:  pq.UpdatedAt,
	}
}

func mapNote(note models.Note, personUUID uuid.UUID, personQualityUUID uuid.UUID) Note {
	return Note{
		UUID:              note.UUID,
		PersonUUID:        personUUID,
		PersonQualityUUID: personQualityUUID,
		Title:             note.Title,
		Body:              note.Body,
		CreatedAt:         note.CreatedAt,
		UpdatedAt:         note.UpdatedAt,
	}
}
