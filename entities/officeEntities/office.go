package officeEntities

import "time"

type Office struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
