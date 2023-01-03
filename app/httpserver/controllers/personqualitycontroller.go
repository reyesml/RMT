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
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

func NewPersonQualityController(cpq interactors.CreatePersonQuality, gpq interactors.GetPersonQuality) personQualityController {
	return personQualityController{
		cpq: cpq,
		gpq: gpq,
	}
}

type personQualityController struct {
	cpq interactors.CreatePersonQuality
	gpq interactors.GetPersonQuality
}

type CreatePersonQualityRequest struct {
	PersonUUID  uuid.UUID `json:"personUUID"`
	QualityName string    `json:"qualityName"`
}

type CreatePersonQualityResponse struct {
	Error string    `json:"error,omitempty"`
	UUID  uuid.UUID `json:"uuid"`
}

func (c personQualityController) Create(w http.ResponseWriter, r *http.Request) {
	var createReq CreatePersonQualityRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreatePersonQualityResponse{Error: "invalid body format"})
		return
	}

	pq, err := c.cpq.Execute(r.Context(), interactors.CreatePersonQualityRequest{
		PersonUUID:  createReq.PersonUUID,
		QualityName: createReq.QualityName,
	})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, CreatePersonQualityResponse{Error: "not found"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, CreatePersonQualityResponse{Error: "something went wrong"})
		return
	}
	render.JSON(w, r, CreatePersonQualityResponse{UUID: pq.UUID})
}

type GetPersonQualityResponse struct {
	Error         string        `json:"error,omitempty"`
	PersonQuality PersonQuality `json:"personQuality"`
}

type PersonQuality struct {
	PersonUUID uuid.UUID `json:"personUUID"`
	Name       string    `json:"name"`
	Notes      []Note    `json:"notes"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Note struct {
	PersonUUID uuid.UUID `json:"personUUID"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
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

func mapPersonQuality(pq models.PersonQuality) PersonQuality {
	notes := make([]Note, 0)
	for _, n := range pq.Notes {
		notes = append(notes, mapNote(n, pq.Person.UUID))
	}
	return PersonQuality{
		PersonUUID: pq.Person.UUID,
		Name:       pq.Quality.Name,
		Notes:      notes,
		CreatedAt:  pq.CreatedAt,
		UpdatedAt:  pq.UpdatedAt,
	}
}

func mapNote(note models.Note, personUUID uuid.UUID) Note {
	return Note{
		PersonUUID: personUUID,
		Title:      note.Title,
		Body:       note.Body,
		CreatedAt:  note.CreatedAt,
		UpdatedAt:  note.UpdatedAt,
	}
}
