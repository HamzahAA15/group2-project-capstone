package dayEntities

import "time"

type Day struct {
	ID        string
	OfficeId  string
	Date      time.Time
	Quota     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
