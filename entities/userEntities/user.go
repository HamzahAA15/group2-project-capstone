package userEntities

import (
	"sirclo/project-capstone/entities/officeEntities"
	"time"
)

type User struct {
	ID        string
	Office    officeEntities.Office
	Avatar    string
	Name      string
	Nik       string
	Email     string
	Password  string
	Phone     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
