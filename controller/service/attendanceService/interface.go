package attendanceService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/utils/request/attendanceRequest"
)

type AttServiceInterface interface {
	GetAttendances(employee, time string) ([]attendanceEntities.Attendance, error)
	CreateAttendance(loginId string, input attendanceRequest.CreateAttRequest) (attendanceEntities.Attendance, error)
	UpdateAttendance(loginId string, input attendanceRequest.UpdateAttRequest) (attendanceEntities.Attendance, error)
	CheckUserRole(loginId string) string
}
