package checkInsOutsRequest

type CheckInsRequest struct {
	ID           string  `json:"id" form:"id"`
	AttendanceID string  `json:"attendance_id" form:"attendance_id"`
	Temprature   float32 `json:"temprature" form:"temprature"`
	IsCheckIns   bool    `json:"is_checkins" form:"is_checkins"`
	IsCheckOuts  bool    `json:"is_checkouts" form:"is_checkouts"`
}

type CheckOutsRequest struct {
	ID           string `json:"id" form:"id"`
	AttendanceID string `json:"attendance_id" form:"attendance_id"`
	IsCheckOuts  bool   `json:"is_checkouts" form:"is_checkouts"`
}
