package attendanceRepository

import "sirclo/project-capstone/entities/attendanceEntities"

//go:generate mockgen --destination=./../../mocks/attendance/repository/mock_repository_attendance.go -source=interface.go
type AttendanceRepoInterface interface {
	GetAttendancesById(attID string) (string, string, error)
	GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error)
	GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error)
	IsCheckins() ([]attendanceEntities.Attendance, error)
	CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
	UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
}
