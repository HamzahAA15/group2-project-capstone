package attendanceResponse

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"time"
)

type AttGetResponse struct {
	ID       string    `json:"id"`
	Day      time.Time `json:"day"`
	OfficeId string    `json:"office_id"`
	Office   string    `json:"office"`
	Avatar   string    `json:"user_avatar"`
	Email    string    `json:"user_email"`
	Nik      string    `json:"nik"`
	Employee string    `json:"employee"`
	Status   string    `json:"status"`
	Notes    string    `json:"notes"`
	Admin    string    `json:"admin"`
}

type AttGetUserResp struct {
	AttGetResponse []AttGetResponse `json:"current_attendances"`
	Total          int              `json:"total"`
}

func FormatGetAtt(att attendanceEntities.Attendance) AttGetResponse {
	fomatter := AttGetResponse{
		ID:       att.ID,
		Day:      att.Day.Date,
		OfficeId: att.OfficeId,
		Office:   att.Office,
		Avatar:   att.Employee.Avatar,
		Email:    att.Employee.Email,
		Nik:      att.Employee.Nik,
		Employee: att.Employee.Name,
		Status:   att.Status,
		Notes:    att.Notes,
		Admin:    att.Admin.Name,
	}
	return fomatter
}
