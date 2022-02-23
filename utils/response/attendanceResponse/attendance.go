package attendanceResponse

import "sirclo/project-capstone/entities/attendanceEntities"

type AttCreateResponse struct {
	ID       string `json:"id"`
	Day      string `json:"day"`
	Employee string `json:"employee"`
}

type AttUpdateResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Notes  string `json:"notes"`
	Admin  string `json:"admin"`
}

type AttGetResponse struct {
	ID       string `json:"id"`
	Day      string `json:"day"`
	Employee string `json:"employee"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
	Admin    string `json:"admin"`
}

func FormatGetAtt(att attendanceEntities.Attendance) AttGetResponse {
	fomatter := AttGetResponse{
		ID:       att.ID,
		Day:      att.Day.ID,
		Employee: att.Employee.ID,
		Status:   att.Status,
		Notes:    att.Notes,
		Admin:    att.Admin.ID,
	}
	return fomatter
}

func FormatAtt(att attendanceEntities.Attendance) AttCreateResponse {
	fomatter := AttCreateResponse{
		ID:       att.ID,
		Day:      att.Day.ID,
		Employee: att.Employee.ID,
	}
	return fomatter
}

func FormatUpdateAtt(att attendanceEntities.Attendance) AttUpdateResponse {
	formatter := AttUpdateResponse{
		ID:     att.ID,
		Status: att.Status,
		Notes:  att.Notes,
		Admin:  att.Admin.ID,
	}
	return formatter
}
