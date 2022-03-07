package checkinsoutsResponse

import (
	"sirclo/project-capstone/entities/checkinEntities"
	"time"
)

type CheckInsOutsResponse struct {
	ID          string    `json:"id"`
	UserName    string    `json:"user_name"`
	UserAvatar  string    `json:"user_avatar"`
	UserEmail   string    `json:"user_email"`
	OfficeName  string    `json:"office_name"`
	Temprature  float32   `json:"temprature"`
	IsCheckIns  bool      `json:"is_checkins"`
	IsCheckOuts bool      `json:"is_checkouts"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatCheckInsOuts(checkinsouts checkinEntities.Checkin) CheckInsOutsResponse {
	fomatter := CheckInsOutsResponse{
		ID:          checkinsouts.ID,
		UserName:    checkinsouts.Attendance.Employee.Name,
		UserAvatar:  checkinsouts.Attendance.Employee.Avatar,
		UserEmail:   checkinsouts.Attendance.Employee.Email,
		OfficeName:  checkinsouts.Attendance.Day.OfficeId.Name,
		Temprature:  checkinsouts.Temprature,
		IsCheckIns:  checkinsouts.IsCheckIns,
		IsCheckOuts: checkinsouts.IsCheckOuts,
		CreatedAt:   checkinsouts.CreatedAt,
		UpdatedAt:   checkinsouts.UpdatedAt,
	}

	return fomatter
}
