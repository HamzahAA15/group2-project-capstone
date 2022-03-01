package checkinEntities

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"time"
)

type Checkin struct {
	ID          string
	Attendance  attendanceEntities.Attendance
	Temprature  float32
	IsCheckIns  bool
	IsCheckOuts bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
