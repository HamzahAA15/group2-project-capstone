package certificateEntities

import (
	"sirclo/project-capstone/entities/userEntities"
	"time"
)

type Certificate struct {
	ID        string
	User      userEntities.User
	Image     string
	Dosage    int
	Status    string
	Admin     userEntities.User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Certificates struct {
	User         userEntities.User
	Certificates []Certificate
}
