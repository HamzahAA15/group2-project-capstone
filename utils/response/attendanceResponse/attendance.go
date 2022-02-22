package attendanceResponse

import "sirclo/project-capstone/entities/attendanceEntities"

type AttCreateResponse struct {
	ID       string `json:"id"`
	Day      string `json:"day"`
	Employee string `json:"employee"`
}

func FormatAtt(att attendanceEntities.Attendance) AttCreateResponse {
	fomatter := AttCreateResponse{
		ID:       att.ID,
		Day:      att.Day.ID,
		Employee: att.Employee.ID,
	}
	return fomatter
}
