package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/interactors"
	"github.com/reyesml/RMT/app/core/models"
	"net/http"
	"time"
)

type PersonController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	CreatePersonQuality(w http.ResponseWriter, r *http.Request)
	ListPersonQualities(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	ListNotes(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

func NewPersonController(
	cp interactors.CreatePerson,
	gp interactors.GetPerson,
	cpq interactors.CreatePersonQuality,
	lpq interactors.ListPersonQualities,
	cpn interactors.CreatePersonNote,
	lpn interactors.ListPersonNotes,
	lp interactors.ListPeople,
) personController {
	return personController{
		cp:  cp,
		gp:  gp,
		cpq: cpq,
		lpq: lpq,
		cpn: cpn,
		lpn: lpn,
		lp:  lp,
	}
}

type personController struct {
	cp  interactors.CreatePerson
	gp  interactors.GetPerson
	cpq interactors.CreatePersonQuality
	lpq interactors.ListPersonQualities
	lp  interactors.ListPeople
	cpn interactors.CreatePersonNote
	lpn interactors.ListPersonNotes
}

type CreatePersonRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreatePersonResponse struct {
	Error string    `json:"error,omitempty"`
	UUID  uuid.UUID `json:"uuid"`
}

func (c personController) Create(w http.ResponseWriter, r *http.Request) {
	var createReq CreatePersonRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreatePersonResponse{Error: "invalid body format"})
		return
	}

	p, err := c.cp.Execute(r.Context(), interactors.CreatePersonRequest{
		FirstName: createReq.FirstName,
		LastName:  createReq.LastName,
	})
	if errors.Is(err, interactors.MissingPersonFields) {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreatePersonResponse{Error: "firstName and lastName cannot both be blank"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, CreatePersonResponse{Error: "something went wrong"})
		return
	}

	render.JSON(w, r, CreatePersonResponse{UUID: p.UUID})
}

type Person struct {
	UUID      uuid.UUID `json:"uuid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetPersonResponse struct {
	Error  string `json:"error,omitempty"`
	Person Person `json:"person,omitempty"`
}

func (c personController) Get(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, GetPersonResponse{Error: "not found"})
		return
	}

	p, err := c.gp.Execute(r.Context(), interactors.GetPersonRequest{UUID: reqUUID})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, GetPersonResponse{Error: "not found"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, GetPersonResponse{Error: "something went wrong"})
		return
	}
	render.JSON(w, r, GetPersonResponse{
		Person: mapPerson(p),
	})
}

type ListPersonQualityResponse struct {
	Error           string          `json:"error,omitempty"`
	PersonQualities []PersonQuality `json:"personQualities"`
}

type CreatePersonQualityRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreatePersonQualityResponse struct {
	Error string    `json:"error,omitempty"`
	UUID  uuid.UUID `json:"uuid"`
}

func (c personController) CreatePersonQuality(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, ListPersonQualityResponse{Error: "not found"})
		return
	}
	var createReq CreatePersonQualityRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreatePersonQualityResponse{Error: "invalid body format"})
		return
	}

	pq, err := c.cpq.Execute(r.Context(), interactors.CreatePersonQualityRequest{
		PersonUUID:  reqUUID,
		QualityName: createReq.Name,
		QualityType: createReq.Type,
	})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, CreatePersonQualityResponse{Error: "not found"})
		return
	}
	if err != nil {
		fmt.Printf("BAD THING: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, CreatePersonQualityResponse{Error: "something went wrong"})
		return
	}
	render.JSON(w, r, CreatePersonQualityResponse{UUID: pq.UUID})
}

func (c personController) ListPersonQualities(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, ListPersonQualityResponse{Error: "not found"})
		return
	}

	pqs, err := c.lpq.Execute(r.Context(), interactors.ListPersonQualitiesRequest{PersonUUID: reqUUID})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, ListPersonQualityResponse{Error: "not found"})
		return
	}
	pql := make([]PersonQuality, 0)
	for _, pq := range pqs {
		pql = append(pql, mapPersonQuality(pq))
	}
	render.JSON(w, r, ListPersonQualityResponse{PersonQualities: pql})
}

type CreatePersonNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type CreatePersonNoteResonse struct {
	Error string    `json:"error,omitempty"`
	UUID  uuid.UUID `json:"uuid"`
}

func (c personController) CreateNote(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, CreatePersonNoteResonse{Error: "not found"})
		return
	}
	var req CreatePersonNoteRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, CreatePersonNoteResonse{Error: "invalid body format"})
		return
	}

	n, err := c.cpn.Execute(r.Context(), interactors.CreatePersonNoteRequest{
		PersonUUID: reqUUID,
		NoteTitle:  req.Title,
		NoteBody:   req.Body,
	})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, CreatePersonNoteResonse{Error: "not found"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, CreatePersonNoteResonse{Error: "something went wrong"})
		return
	}
	render.JSON(w, r, CreatePersonNoteResonse{UUID: n.UUID})
}

type ListPersonNotesResponse struct {
	Error string `json:"error,omitempty"`
	Notes []Note `json:"notes"`
}

func (c personController) ListNotes(w http.ResponseWriter, r *http.Request) {
	reqUUIDParam := chi.URLParam(r, "UUID")
	reqUUID, err := uuid.Parse(reqUUIDParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, ListPersonNotesResponse{Error: "not found"})
		return
	}

	ns, err := c.lpn.Execute(r.Context(), interactors.ListPersonNotesRequest{
		PersonUUID: reqUUID,
		Filter:     interactors.OnlyPersonNotesFilter,
	})
	if errors.Is(err, interactors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, ListPersonNotesResponse{Error: "not found"})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, ListPersonNotesResponse{Error: "something went wrong"})
		return
	}
	result := make([]Note, 0)
	for _, n := range ns {
		result = append(result, mapNote(n, reqUUID, n.PersonQuality.UUID))
	}
	render.JSON(w, r, ListPersonNotesResponse{Notes: result})
}

type ListPersonResponse struct {
	Error  string   `json:"error,omitempty"`
	People []Person `json:"people,omitempty"`
}

func (c personController) List(w http.ResponseWriter, r *http.Request) {
	ppl, err := c.lp.Execute(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, GetPersonResponse{Error: "something went wrong"})
		return
	}
	result := make([]Person, 0)
	for _, p := range ppl {
		result = append(result, mapPerson(p))
	}
	render.JSON(w, r, ListPersonResponse{People: result})
}

func mapPerson(p models.Person) Person {
	return Person{
		UUID:      p.UUID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
