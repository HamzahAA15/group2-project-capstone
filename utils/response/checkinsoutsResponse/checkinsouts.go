package checkinsoutsResponse

import (
	"sirclo/project-capstone/entities/checkinEntities"
	"time"
)

type CheckInsOutsResponse struct {
	ID           string    `json:"id"`
	AttendanceID string    `json:"attendance_id"`
	Temprature   float32   `json:"temprature"`
	IsCheckIns   bool      `json:"is_checkins"`
	IsCheckOuts  bool      `json:"is_checkouts"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FormatCheckInsOuts(checkinsouts checkinEntities.Checkin) CheckInsOutsResponse {
	fomatter := CheckInsOutsResponse{
		ID:           checkinsouts.ID,
		AttendanceID: checkinsouts.AttendanceID,
		Temprature:   checkinsouts.Temprature,
		IsCheckIns:   checkinsouts.IsCheckIns,
		IsCheckOuts:  checkinsouts.IsCheckOuts,
		CreatedAt:    checkinsouts.CreatedAt,
		UpdatedAt:    checkinsouts.UpdatedAt,
	}

	return fomatter
}
