package attendanceService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/utils/request/attendanceRequest"
)

type AttServiceInterface interface {
<<<<<<< HEAD
	GetAttendancesById(attID string) (string, string, error)
=======
>>>>>>> 7b50bc7e692324b69db179a1087fc11f0dee3064
	GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error)
	GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error)
	CreateAttendance(loginId string, input attendanceRequest.CreateAttRequest) (attendanceEntities.Attendance, error)
	UpdateAttendance(loginId string, input attendanceRequest.UpdateAttRequest) (attendanceEntities.Attendance, error)
	CheckUserRole(loginId string) string
}
