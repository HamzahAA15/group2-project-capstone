package attendanceRepository

import "sirclo/project-capstone/entities/attendanceEntities"

type AttendanceRepoInterface interface {
	GetAttendances() ([]attendanceEntities.Attendance, error)
	CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
	UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
}
