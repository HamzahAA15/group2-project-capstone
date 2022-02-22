package attendanceRepository

import "sirclo/project-capstone/entities/attendanceEntities"

type AttendanceRepoInterface interface {
	CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
	UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error)
}
