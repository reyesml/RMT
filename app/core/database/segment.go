package database

import (
	"github.com/google/uuid"
)

// Segmented includes a SegmentUUID. This should be included in
// models that need to be segmented across different user accounts.
// We use SegmentUUID instead of UserUUID in the event of users
// needing to access the same underlying data, such as a Team.
type Segmented struct {
	SegmentUUID uuid.UUID `gorm:"not null;check:segment_uuid <> '00000000-0000-0000-0000-000000000000'"`
}

// TODO: Add segment scope and "afterFind" methods to enforce data segmentation
// We should add the scope once a user has been identified (auth middleware?)
