package dayEntities

import (
	"sirclo/project-capstone/entities/officeEntities"
	"time"
)

type Day struct {
	ID             string
	OfficeId       officeEntities.Office
	Date           time.Time
	Quota          int
	TotalApproved  int
	RemainingQuota int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
