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

type AttIsCheckins struct {
	ID            string `json:"id"`
	UserAvatar    string `json:"user_avatar"`
	UserEmail     string `json:"user_email"`
	UserName      string `json:"user_name"`
	OfficeName    string `json:"office_name"`
	StatusCheckin string `json:"status_checkin"`
}

func FormatDataIsCheckins(att attendanceEntities.Attendance) AttIsCheckins {
	fomatter := AttIsCheckins{
		ID:            att.ID,
		UserAvatar:    att.Employee.Avatar,
		UserEmail:     att.Employee.Email,
		UserName:      att.Employee.Name,
		OfficeName:    att.Day.OfficeId.Name,
		StatusCheckin: att.StatusCheckin,
	}

	return fomatter
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
