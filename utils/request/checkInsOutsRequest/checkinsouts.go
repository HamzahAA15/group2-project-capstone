package checkInsOutsRequest

import "time"

type CheckInsRequest struct {
	ID           string    `json:"id" form:"id"`
	AttendanceID string    `json:"attendance_id" form:"attendance_id"`
	Temprature   float32   `json:"temprature" form:"temprature"`
	IsCheckIns   bool      `json:"is_checkins" form:"is_checkins"`
	IsCheckOuts  bool      `json:"is_checkouts" form:"is_checkouts"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

type CheckOutsRequest struct {
	ID          string    `json:"id" form:"id"`
	IsCheckOuts bool      `json:"is_checkouts" form:"is_checkouts"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}
