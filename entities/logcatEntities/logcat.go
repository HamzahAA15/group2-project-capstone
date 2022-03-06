package logcatEntities

import (
	"sirclo/project-capstone/entities/userEntities"
	"time"
)

type Logcat struct {
	ID        string
	User      userEntities.User
	Message   string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
