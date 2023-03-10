package interactors

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
)

type CreatePersonQualityRequest struct {
	PersonUUID  uuid.UUID
	QualityName string
	QualityType string
}

var allowedQualityTypes = map[string]bool{"default": true, "expertise": true}

var ErrInvalidQualityType = errors.New("invalid quality type")

var MissingQualityNameErr = errors.New("QualityName is required")

type CreatePersonQuality interface {
	Execute(ctx context.Context, req CreatePersonQualityRequest) (models.PersonQuality, error)
}

func NewCreatePersonQuality(pr repos.PersonRepo, qr repos.QualityRepo, pqr repos.PersonQualityRepo) createPersonQuality {
	return createPersonQuality{
		pr:  pr,
		qr:  qr,
		pqr: pqr,
	}
}

type createPersonQuality struct {
	pr  repos.PersonRepo
	qr  repos.QualityRepo
	pqr repos.PersonQualityRepo
}

// Execute validates that the person exists and is visible by
// the current user.  It then checks to see if the quality exists,
// and creates the quality if it doesn't exist. It then assigns the
// quality to the person.
func (ia createPersonQuality) Execute(ctx context.Context, req CreatePersonQualityRequest) (models.PersonQuality, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil || user.SegmentUUID == uuid.Nil {
		return models.PersonQuality{}, database.SegmentMissingErr
	}

	if len(req.QualityName) == 0 {
		return models.PersonQuality{}, MissingQualityNameErr
	}
	var qt string
	if len(req.QualityType) == 0 {
		qt = "default"
	} else {
		qt = req.QualityType
	}
	if _, ok := allowedQualityTypes[qt]; !ok {
		return models.PersonQuality{}, ErrInvalidQualityType
	}

	// Find our person
	p, err := ia.pr.GetByUUID(req.PersonUUID, user.SegmentUUID)
	if err != nil {
		return models.PersonQuality{}, ErrNotFound
	}

	// Find/Create our quality
	var q models.Quality
	fqs, err := ia.qr.FindByNameAndType(req.QualityName, qt, user.ID, user.SegmentUUID)
	if len(fqs) == 0 {
		// Insert a new quality record
		q.Name = req.QualityName
		q.Type = qt
		q.SegmentUUID = user.SegmentUUID
		q.UserId = user.ID
		ia.qr.Create(&q)
	} else {
		// User our existing quality record
		q = fqs[0]
	}

	// Create our personQuality
	pq := models.PersonQuality{
		Segmented: database.Segmented{SegmentUUID: user.SegmentUUID},
		PersonId:  p.ID,
		QualityId: q.ID,
		UserId:    user.ID,
	}
	if err := ia.pqr.Create(&pq); err != nil {
		return models.PersonQuality{}, err
	}
	return pq, nil
}
