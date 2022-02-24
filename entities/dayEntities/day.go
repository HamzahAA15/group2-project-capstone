package dayEntities

import (
	"sirclo/project-capstone/entities/officeEntities"
	"time"
)

type Day struct {
	ID        string
	OfficeId  officeEntities.Office
	Date      time.Time
	Quota     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
