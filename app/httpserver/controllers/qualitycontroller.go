package controllers

import (
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/interactors"
	"github.com/reyesml/RMT/app/core/models"
	"net/http"
	"time"
)

type QualityController interface {
	List(w http.ResponseWriter, r *http.Request)
}

func NewQualityController(lq interactors.ListQualities) qualityController {
	return qualityController{
		lq: lq,
	}
}

type qualityController struct {
	lq interactors.ListQualities
}

type ListQualityResponse struct {
	Error     string    `json:"error,omitempty"`
	Qualities []Quality `json:"qualities"`
}

type Quality struct {
	UUID      uuid.UUID `json:"uuid,omitempty"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c qualityController) List(w http.ResponseWriter, r *http.Request) {
	qs, err := c.lq.Execute(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, ListQualityResponse{Error: "something went wrong"})
		return
	}

	result := make([]Quality, 0)
	for _, q := range qs {
		result = append(result, mapQuality(q))
	}
	render.JSON(w, r, ListQualityResponse{Qualities: result})
}

func mapQuality(q models.Quality) Quality {
	return Quality{
		UUID:      q.UUID,
		Name:      q.Name,
		CreatedAt: q.CreatedAt,
	}
}
