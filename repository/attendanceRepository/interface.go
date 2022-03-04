package attendanceRepository

import "sirclo/project-capstone/entities/attendanceEntities"

type AttendanceRepoInterface interface {
	GetAttendances(employee, time, status, office, order string) ([]attendanceEntities.Attendance, error)
	GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, office, order string) ([]attendanceEntities.Attendance, error)
	GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error)
	CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
	UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
}
