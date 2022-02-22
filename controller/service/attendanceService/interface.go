package attendanceService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/utils/request/attendanceRequest"
)

type AttServiceInterface interface {
	CreateAttendance(loginId string, input attendanceRequest.CreateAttRequest) (attendanceEntities.Attendance, error)
}
