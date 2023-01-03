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

type PersonController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	ListPersonQualities(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

func NewPersonController(cp interactors.CreatePerson, gp interactors.GetPerson, lpq interactors.ListPersonQualities, lp interactors.ListPeople) personController {
	return personController{
		cp:  cp,
		gp:  gp,
		lpq: lpq,
		lp:  lp,
	}
}

type personController struct {
	cp  interactors.CreatePerson
	gp  interactors.GetPerson
	lpq interactors.ListPersonQualities
	lp  interactors.ListPeople
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
	Error           string `json:"error,omitempty"`
	PersonQualities []PersonQuality
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
	var result []Person
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
