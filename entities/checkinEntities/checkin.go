package checkinEntities

import (
	"time"
)

type Checkin struct {
	ID           string
	AttendanceID string
	Temprature   float32
	IsCheckIns   bool
	IsCheckOuts  bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
