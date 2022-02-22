package attendanceEntities

import (
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/entities/userEntities"
	"time"
)

type Attendance struct {
	ID        string
	Day       dayEntities.Day
	Employee  userEntities.User
	Status    string
	Notes     string
	Admin     userEntities.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
