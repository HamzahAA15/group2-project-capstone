package userEntities

import "time"

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
