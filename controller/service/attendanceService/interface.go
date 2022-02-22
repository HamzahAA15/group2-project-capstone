package attendanceService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/utils/request/attendaceRequest"
)

type AttServiceInterface interface {
	CreateAttendace(input attendaceRequest.CreateAttRequest) (attendanceEntities.Attendance, error)
}
