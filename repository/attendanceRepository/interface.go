package attendanceRepository

import "sirclo/project-capstone/entities/attendanceEntities"

type AttendanceRepoInterface interface {
	GetAttendancesById(attID string) (string, string, error)
	GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error)
	GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error)
	CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
	UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
}
