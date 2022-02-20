package userEntities

import "time"

type User struct {
	ID        string
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
